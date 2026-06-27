# Go Core Concepts Practice Tracker

**Overall: 5/79 solved**

> Legend: `[ ]` Unsolved · `[~]` Attempted · `[x]` Solved · `[★]` Mastered · `►` Current

---

## Fundamentals [1/30]

> Mixed topics, very easy to easy. Start here if you want to warm up before the deeper groups.

| ID | Challenge | Topic | Status |
|----|-----------|-------|--------|
| F01 | Launch a goroutine, print "hello" from inside it | Goroutines | [x] |
| F02 | Launch a goroutine and wait for it to finish using WaitGroup | Goroutines | [x] |
| F03 | Send a value into an unbuffered channel from a goroutine, receive in main | Channels | [x] |
| F04 | Create a buffered channel of size 3, send 3 values without blocking | Channels | [x] |
| F05 | Close a channel after sending; iterate its values with range | Channels | [x] |
| F06 ► | Use select to receive from whichever of two channels is ready first | Select | [ ] |
| F07 | Use select with a default case to do a non-blocking receive | Select | [ ] |
| F08 | Protect a shared integer counter with sync.Mutex across 5 goroutines | sync | [ ] |
| F09 | Use sync.Once to initialize a config struct exactly once | sync | [ ] |
| F10 | Use sync.WaitGroup to launch 5 goroutines and wait for all to finish | sync | [ ] |
| F11 | Increment a counter from 100 goroutines using atomic.AddInt64 | atomic | [ ] |
| F12 | Create a context with cancel, pass it to a goroutine, call cancel() | Context | [ ] |
| F13 | Create a context with a 1-second timeout; show it expires | Context | [ ] |
| F14 | Use context.WithValue to attach a request-ID string, read it inside a func | Context | [ ] |
| F15 | Use defer to ensure a file (or fake closer) is always closed | defer | [ ] |
| F16 | Use defer mu.Unlock() immediately after mu.Lock() in a function | defer | [ ] |
| F17 | Define a sentinel error with errors.New; check it with errors.Is | Errors | [ ] |
| F18 | Wrap an error with fmt.Errorf("...: %w", err); unwrap with errors.As | Errors | [ ] |
| F19 | Create a custom error type with an extra Code field; implement error interface | Errors | [ ] |
| F20 | Use recover() inside a deferred function to catch a panic | panic/recover | [ ] |
| F21 | Implement fmt.Stringer on a struct (print a formatted representation) | Interfaces | [ ] |
| F22 | Define a small interface (Doer); implement it on two different structs | Interfaces | [ ] |
| F23 | Type-assert an interface value to a concrete type; handle the false case | Interfaces | [ ] |
| F24 | Implement io.Reader on a struct that reads from a fixed string | Interfaces | [ ] |
| F25 | Write a generic Map[T, U] function that transforms a slice | Generics | [ ] |
| F26 | Write a goroutine that stops when a done channel is closed | Channels | [ ] |
| F27 | Fan-out: launch N goroutines each writing a result to a shared slice (with mutex) | Goroutines | [ ] |
| F28 | Use time.After in a select to time out a slow operation | Select | [ ] |
| F29 | Write a table-driven test for a simple pure function | Testing | [ ] |
| F30 | Use strings.Builder to concatenate 1000 strings efficiently | Performance | [ ] |

---

## Group 1: Goroutines [0/4]

| ID | Challenge | Difficulty | Status |
|----|-----------|------------|--------|
| G01 | Launch N goroutines with WaitGroup, collect results in order | Easy | [ ] |
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
