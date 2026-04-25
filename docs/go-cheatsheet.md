# Go Concepts Cheatsheet

Quick reference for Go core concepts, concurrency patterns, and common pitfalls.

---

## Goroutines

| Concept | Key Fact |
|---------|----------|
| Cost | ~2-8KB stack, grows dynamically (OS thread: ~1MB fixed) |
| Scheduling | M:N on OS threads, managed by Go runtime (GMP model) |
| GOMAXPROCS | Number of OS threads (P's); defaults to runtime.NumCPU() |
| Stack growth | Doubled when needed; shrunk by GC |
| Preemption | Async preemption since Go 1.14 (no more tight loops blocking scheduler) |

```go
go func() { /* goroutine */ }()
runtime.NumGoroutine() // count alive goroutines
runtime.GOMAXPROCS(n)  // set Ps (usually left at default)
```

---

## Channels

| Type | Behaviour |
|------|-----------|
| `make(chan T)` | Unbuffered — synchronous rendezvous |
| `make(chan T, n)` | Buffered — async up to n, then blocks |
| nil channel | Blocks forever on send/receive — useful for disabling select cases |
| Closed channel receive | Returns `(zero, false)` |
| Send on closed channel | **PANIC** |
| Close from receiver | **PANIC** if sender sends after — rule: sender closes |

```go
ch := make(chan int, 5)
ch <- 42          // send
v := <-ch         // receive
v, ok := <-ch     // receive with close check
close(ch)         // only sender closes
for v := range ch { } // iterate until closed
```

---

## Select

```go
select {
case v := <-ch1:     // receive
case ch2 <- v:       // send
case <-time.After(t): // timeout
case <-ctx.Done():   // cancellation
default:             // non-blocking (run if no case ready)
}
```

| Behaviour | Detail |
|-----------|--------|
| Multiple ready cases | Go picks one at random |
| No ready cases | Blocks until one is ready |
| `default` present | Never blocks — returns immediately |
| Nil channel case | Never selected |

---

## sync Package

| Type | Use Case |
|------|----------|
| `sync.Mutex` | Exclusive access to shared state |
| `sync.RWMutex` | Multiple readers OR one writer |
| `sync.WaitGroup` | Wait for N goroutines to finish |
| `sync.Once` | Run exactly once, goroutine-safe |
| `sync.Cond` | Signal/broadcast goroutines waiting on a condition |
| `sync.Map` | Concurrent map for mostly-read workloads |
| `sync.Pool` | Reuse objects to reduce GC pressure |

```go
var mu sync.Mutex
mu.Lock()
defer mu.Unlock()

var wg sync.WaitGroup
wg.Add(1)          // BEFORE launching goroutine
go func() { defer wg.Done(); work() }()
wg.Wait()

var once sync.Once
once.Do(func() { expensiveInit() })
```

**WaitGroup gotcha:** `Add()` must be called before `go func()`, not inside it.

---

## sync/atomic

```go
var counter int64
atomic.AddInt64(&counter, 1)
atomic.LoadInt64(&counter)
atomic.StoreInt64(&counter, 0)
atomic.CompareAndSwapInt64(&counter, old, new) // CAS

// Go 1.19+
var flag atomic.Bool
flag.Store(true)
flag.Load()
```

**Use atomic when:** Single variable, simple operations (increment, flag).
**Use mutex when:** Multiple variables must be updated together atomically.

---

## Context

```go
ctx, cancel := context.WithCancel(parent)
defer cancel() // ALWAYS defer cancel

ctx, cancel := context.WithTimeout(parent, 5*time.Second)
ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))

// Request-scoped values (use unexported key type to avoid collisions)
type ctxKey struct{}
ctx = context.WithValue(ctx, ctxKey{}, "request-id-123")
v := ctx.Value(ctxKey{})

// Check cancellation
select {
case <-ctx.Done():
    return ctx.Err() // context.Canceled or context.DeadlineExceeded
default:
}
```

**Rules:**
- Pass `ctx` as first function parameter, never store in struct
- Always `defer cancel()` to release resources
- Pass ctx to all blocking calls (HTTP, DB, gRPC)

---

## errgroup

```go
import "golang.org/x/sync/errgroup"

g, ctx := errgroup.WithContext(parentCtx)
for _, item := range items {
    item := item // capture for Go < 1.22
    g.Go(func() error {
        return process(ctx, item)
    })
}
if err := g.Wait(); err != nil {
    // first non-nil error; ctx is already cancelled
}
```

**errgroup vs WaitGroup:**
- WaitGroup: know when done, no error propagation
- errgroup: know when done + cancel on first error + collect error

---

## OS Signals

```go
sigs := make(chan os.Signal, 1)
signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
sig := <-sigs // block until signal

// Graceful HTTP server shutdown
srv := &http.Server{Addr: ":8080"}
go srv.ListenAndServe()

<-sigs
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()
srv.Shutdown(ctx)
```

---

## Error Handling

```go
// Wrap with context
return fmt.Errorf("loading config: %w", err)

// Unwrap
if errors.Is(err, fs.ErrNotExist) { }
var pathErr *fs.PathError
if errors.As(err, &pathErr) { }

// Sentinel errors (define at package level)
var ErrNotFound = errors.New("not found")

// Custom error type
type ValidationError struct {
    Field   string
    Message string
}
func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}
```

**Rules:**
- Wrap with `%w` (not `%v`) to enable unwrapping
- Use `errors.Is` for sentinel errors (handles wrapped errors)
- Use `errors.As` for type-specific error inspection

---

## Interfaces

```go
// Accept interfaces, return concrete types
func Process(r io.Reader) *Result { }

// Interface embedding
type ReadWriter interface {
    io.Reader
    io.Writer
}

// Type assertion
v, ok := i.(ConcreteType)

// Type switch
switch v := i.(type) {
case *http.Request: // v is *http.Request here
case string:
default:
}
```

**Key interfaces to know:**
- `io.Reader`, `io.Writer`, `io.Closer`, `io.ReadWriter`
- `fmt.Stringer` (`String() string`)
- `error` (`Error() string`)
- `http.Handler` (`ServeHTTP(ResponseWriter, *Request)`)
- `sort.Interface` (`Len/Less/Swap`)

---

## Common Goroutine Leak Patterns

```go
// ❌ Leak: nobody reads from ch after timeout
func fetchWithTimeout() {
    ch := make(chan Result) // unbuffered!
    go func() { ch <- fetch() }() // goroutine stuck if nobody reads
    select {
    case r := <-ch: return r
    case <-time.After(time.Second): return ErrTimeout // goroutine leaks!
    }
}

// ✅ Fix: buffered channel or pass ctx
func fetchWithTimeout(ctx context.Context) {
    ch := make(chan Result, 1) // buffered — goroutine can always send
    go func() {
        select {
        case ch <- fetch():
        case <-ctx.Done():
        }
    }()
    select {
    case r := <-ch: return r
    case <-ctx.Done(): return ctx.Err()
    }
}
```

---

## Generics (Go 1.18+)

```go
// Type parameter
func Map[T, U any](s []T, f func(T) U) []U {
    result := make([]U, len(s))
    for i, v := range s { result[i] = f(v) }
    return result
}

// Constraint
type Number interface { int | int64 | float64 }
func Sum[T Number](nums []T) T { ... }

// comparable constraint (enables == and map keys)
func Contains[T comparable](s []T, v T) bool { ... }

type Set[T comparable] map[T]struct{}
```

---

## Testing Patterns

```go
// Table-driven
func TestFoo(t *testing.T) {
    cases := []struct{ in, want int }{ {1, 2}, {2, 4} }
    for _, tc := range cases {
        t.Run(fmt.Sprint(tc.in), func(t *testing.T) {
            if got := foo(tc.in); got != tc.want {
                t.Errorf("got %d, want %d", got, tc.want)
            }
        })
    }
}

// Benchmark
func BenchmarkFoo(b *testing.B) {
    for i := 0; i < b.N; i++ { foo(42) }
}

// Race detector: go test -race ./...
// Goroutine leak: go.uber.org/goleak
func TestNoLeak(t *testing.T) {
    defer goleak.VerifyNone(t)
    // ... test code ...
}
```

---

## Quick Rules

| Rule | Detail |
|------|--------|
| Close only from sender | Closing from receiver can panic |
| WaitGroup.Add before go | Add inside goroutine races with Wait |
| Always defer cancel() | WithCancel/Timeout leaks goroutine without it |
| Check ctx.Done in loops | Long-running goroutines must be cancellable |
| Don't store context in struct | Pass as first arg to functions |
| Use errors.Is not == | Handles wrapped errors correctly |
| Avoid time.After in loops | Creates new goroutine/timer each call |
| Buffer channel = semaphore | cap N means max N concurrent operations |
