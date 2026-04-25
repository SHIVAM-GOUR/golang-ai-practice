# Go Core Concepts Practice Tracker

**Overall: 0/49 solved**

> Legend: `[ ]` Unsolved · `[~]` Attempted · `[x]` Solved · `[★]` Mastered · `►` Current

---

## Group 1: Goroutines [0/4]

| ID | Challenge | Difficulty | Status |
|----|-----------|------------|--------|
| G01 ► | Launch N goroutines with WaitGroup, collect results in order | Easy | [ ] |
| G02 | Fix a data race on a shared counter | Easy | [ ] |
| G03 | Detect and fix a goroutine leak | Medium | [ ] |
| G04 | Semaphore: limit max concurrent goroutines using a buffered channel | Medium | [ ] |

---

## Group 2: Channels [0/5]

| ID | Challenge | Difficulty | Status |
|----|-----------|------------|--------|
| C01 | Demonstrate unbuffered deadlock, fix with buffered channel | Easy | [ ] |
| C02 | Generator: function that returns a `<-chan int` | Easy | [ ] |
| C03 | Pipeline: three-stage channel chain (generate → square → print) | Medium | [ ] |
| C04 | Cancel a pipeline mid-stream with a done channel | Medium | [ ] |
| C05 | Nil channel trick: disable a select case dynamically | Hard | [ ] |

---

## Group 3: Select [0/3]

| ID | Challenge | Difficulty | Status |
|----|-----------|------------|--------|
| S01 | Non-blocking receive with select + default | Easy | [ ] |
| S02 | Operation timeout using time.After in select | Easy | [ ] |
| S03 | Fan-in: merge two channels into one using select | Medium | [ ] |

---

## Group 4: sync Package [0/4]

| ID | Challenge | Difficulty | Status |
|----|-----------|------------|--------|
| P01 | Mutex: protect a shared map from concurrent writes | Easy | [ ] |
| P02 | sync.Once: goroutine-safe singleton initialization | Easy | [ ] |
| P03 | RWMutex: read-heavy cache — show throughput gain vs Mutex | Medium | [ ] |
| P04 | sync.Cond: producer/consumer with Wait/Broadcast | Hard | [ ] |

---

## Group 5: sync/atomic [0/2]

| ID | Challenge | Difficulty | Status |
|----|-----------|------------|--------|
| A01 | Thread-safe counter using atomic.AddInt64 | Easy | [ ] |
| A02 | CAS loop: compare-and-swap retry pattern | Medium | [ ] |

---

## Group 6: Context [0/4]

| ID | Challenge | Difficulty | Status |
|----|-----------|------------|--------|
| K01 | WithCancel: cancel a tree of goroutines from a single cancel() | Easy | [ ] |
| K02 | WithTimeout: HTTP call that must finish within 500ms | Easy | [ ] |
| K03 | WithValue: thread a request-ID through a call chain | Medium | [ ] |
| K04 | Fix a goroutine that ignores ctx.Done() (context leak) | Hard | [ ] |

---

## Group 7: Concurrency Patterns [0/5]

| ID | Challenge | Difficulty | Status |
|----|-----------|------------|--------|
| CP01 | Worker pool: N workers processing M jobs, collect results | Medium | [ ] |
| CP02 | Rate limiter: allow K requests/second using time.Ticker | Medium | [ ] |
| CP03 | errgroup: run N goroutines, cancel all on first error | Medium | [ ] |
| CP04 | Producer-consumer with bounded buffer and backpressure | Medium | [ ] |
| CP05 | Circuit breaker: open after N failures, reset after timeout | Hard | [ ] |

---

## Group 8: OS Signals & Graceful Shutdown [0/2]

| ID | Challenge | Difficulty | Status |
|----|-----------|------------|--------|
| OS01 | Trap SIGINT/SIGTERM, clean up resources, exit cleanly | Easy | [ ] |
| OS02 | HTTP server graceful shutdown with server.Shutdown(ctx) | Medium | [ ] |

---

## Group 9: Error Handling [0/4]

| ID | Challenge | Difficulty | Status |
|----|-----------|------------|--------|
| E01 | Custom error type implementing the error interface | Easy | [ ] |
| E02 | Wrap errors with fmt.Errorf %w, inspect with errors.Is / errors.As | Easy | [ ] |
| E03 | panic/recover: catch panics in a goroutine without crashing | Medium | [ ] |
| E04 | Retry with exponential backoff on a flaky operation | Medium | [ ] |

---

## Group 10: Interfaces [0/3]

| ID | Challenge | Difficulty | Status |
|----|-----------|------------|--------|
| I01 | Implement io.Reader on a custom struct | Easy | [ ] |
| I02 | Type assertion and type switch on an interface value | Medium | [ ] |
| I03 | Mock an interface for testing (inject fake HTTP client) | Medium | [ ] |

---

## Group 11: Generics [0/3]

| ID | Challenge | Difficulty | Status |
|----|-----------|------------|--------|
| GEN01 | Generic Stack[T] with Push/Pop | Easy | [ ] |
| GEN02 | Generic Map[T, U] function over a slice | Easy | [ ] |
| GEN03 | Generic Set[T comparable] using map | Medium | [ ] |

---

## Group 12: HTTP & Networking [0/3]

| ID | Challenge | Difficulty | Status |
|----|-----------|------------|--------|
| H01 | HTTP handler that cancels slow requests via context timeout | Medium | [ ] |
| H02 | Middleware chain: logging + recovery in net/http | Medium | [ ] |
| H03 | HTTP client with retry and exponential backoff | Medium | [ ] |

---

## Group 13: io Patterns [0/3]

| ID | Challenge | Difficulty | Status |
|----|-----------|------------|--------|
| IO01 | Custom io.Writer that counts bytes written | Easy | [ ] |
| IO02 | io.Pipe: connect a writer goroutine to a reader goroutine | Medium | [ ] |
| IO03 | bufio.Scanner: stream-process a large file line by line | Easy | [ ] |

---

## Group 14: Testing [0/4]

| ID | Challenge | Difficulty | Status |
|----|-----------|------------|--------|
| T01 | Table-driven test with t.Run subtests | Easy | [ ] |
| T02 | Benchmark: write a Benchmark* and interpret ns/op output | Medium | [ ] |
| T03 | Write a test that the race detector catches, then fix it | Medium | [ ] |
| T04 | Test for goroutine leaks using goleak.VerifyNone | Medium | [ ] |

---

## Group 15: Memory & Performance [0/3]

| ID | Challenge | Difficulty | Status |
|----|-----------|------------|--------|
| M01 | Escape analysis: prove a variable escapes with `go build -gcflags="-m"` | Medium | [ ] |
| M02 | sync.Pool: reuse buffers across goroutines to reduce GC | Medium | [ ] |
| M03 | Pre-allocate slices: show alloc difference with make([]T, 0, n) | Easy | [ ] |

---

> Edit this file to update status: `[ ]` → `[x]` solved · `[~]` attempted · `[★]` mastered
> Move `►` to the next unsolved problem when you advance.
