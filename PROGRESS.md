# Progress Tracker

## Streak
**Current streak:** 2 days
**Last practiced:** 2026-05-13
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
| 2026-04-29 | 1 | 1 | F01: goroutine + WaitGroup basics |
| 2026-05-01 | 1 | 1 | F02: WaitGroup pattern clean |
| 2026-05-01 | 1 | 1 | F03: unbuffered channel send/receive |
| 2026-05-01 | 1 | 1 | F04: buffered channel, no goroutine needed |
| 2026-05-01 | 1 | 1 | F05: close channel, range to drain |
| 2026-05-07 | 1 | 1 | F06: select, receive from first ready channel |
| 2026-05-07 | 1 | 1 | F07: select with default for non-blocking receive |
| 2026-05-12 | 1 | 1 | F08: sync.Mutex pointer sharing, WaitGroup pattern |
| 2026-05-13 | 3 | 3 | F09: sync.Once + WaitGroup; F10: loop variable capture (&i bug); F11: atomic.AddInt64 |
| 2026-05-16 | 1 | 1 | F12: context.WithCancel, goroutine respects ctx.Done() via select |
| 2026-05-17 | 1 | 1 | F13: context.WithTimeout, select with time.After to show deadline exceeded |
| 2026-05-18 | 1 | 1 | F14: context.WithValue, unexported key type, nil check + type assertion |
| 2026-06-06 | 2 | 2 | F15: defer file.Close() after error check — canonical Go cleanup pattern; F16: defer mu.Unlock() after mu.Lock(), mutex must be passed as pointer |

---

## Notes & Observations

> Recurring mistakes, gotchas, aha moments.

- (empty)
