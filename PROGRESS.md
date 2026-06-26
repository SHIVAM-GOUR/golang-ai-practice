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
| 2026-06-26 | 1 | 1 | F01: launch goroutine, print hello |

---

## Notes & Observations

> Recurring mistakes, gotchas, aha moments.

- (empty)
