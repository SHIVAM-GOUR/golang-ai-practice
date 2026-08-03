# Progress Tracker

## Streak
**Current streak:** 1 day
**Last practiced:** 2026-06-26
**Longest streak:** 2 days

---

## Milestones
- [x] First problem solved
- [ ] Complete Fundamentals (30/30)
- [ ] Complete Goroutines (4/4)
- [ ] Complete Channels (5/5)
- [ ] 40 problems solved
- [ ] Complete Select + sync + atomic (9/9)
- [ ] Complete Context (4/4)
- [ ] Complete Concurrency Patterns (5/5)
- [ ] 60 problems solved
- [ ] Complete Signals + Errors (6/6)
- [ ] Complete Interfaces + Generics (6/6)
- [ ] Complete HTTP + io + Testing + Memory (13/13)
- [ ] All 79 problems solved 🎉

**Next milestone:** Complete Fundamentals (30/30)

---

## Weak Patterns

> Patterns where you keep getting tripped up. Update via "log today's session".

| Pattern | Attempted | Mastered | Notes |
|---------|-----------|----------|-------|
| Loop variable capture | F10 | — | Passed &i instead of i; all goroutines saw final value (5) |
| Atomic on wrong pointer | F11 | — | Called AddInt64 on local copy instead of the actual shared var |
| Two atomics ≠ atomic | F11 (cross-Q) | — | Load + Add separately is still a race; use single AddInt64 |

---

## Session Log

| Date | Problems Attempted | Problems Solved | Notes |
|------|-------------------|-----------------|-------|
| 2026-06-26 | 3 | 3 | F01: launch goroutine, print hello. F02: WaitGroup to wait for goroutine. F03: unbuffered channel send/receive |
| 2026-06-27 | 3 | 3 | F04: buffered channel of size 3, send 3 values without blocking. F05: close channel, iterate with range. F06: select from two channels |
| 2026-06-28 | 11 | 11 | F07: select with default case for non-blocking receive; caught goroutine leak with unbuffered channels. F08: sync.Mutex to protect shared counter; replaced time.Sleep with WaitGroup. F09: sync.Once to initialize config struct exactly once across goroutines. F10: WaitGroup to launch 5 goroutines and wait for all to finish. F11: atomic.AddInt64 to increment counter from 100 goroutines. F12: context.WithCancel passed to goroutine; goroutine exits on ctx.Done(). F13: context.WithTimeout expires after 1s; ctx.Err() shows deadline exceeded. F14: context.WithValue with unexported key type to avoid collisions. F15: defer close on fake resource; fires even on early return. F16: defer mu.Unlock() immediately after mu.Lock(); unlocks even on early return. F17: sentinel error with errors.New; checked with errors.Is |
| 2026-06-29 | 5 | 5 | F19: custom error type with Code field; implemented error interface, wrapped with %w, unwrapped with errors.As. F20: recover() inside deferred function catches panic; named return lets defer set the error value. F21: implemented fmt.Stringer on a struct; used %s for string field and %d for int field in Sprintf. F22: defined Doer interface; implemented on two structs; passed both to a function accepting Doer. F23: type assertion with ok form; handled both success and failure cases |

---

## Notes & Observations

> Recurring mistakes, gotchas, aha moments.

- (empty)
