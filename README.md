<div align="center">

# Go Sensei 🥋

**Your Personal Go Core Concepts Mentor**

Machine-round interview prep for Go developers. Practice goroutines, channels, concurrency patterns, and more — with intelligent hints, code review, and mock interviews.

[![Claude Code](https://img.shields.io/badge/Claude-Code-blueviolet)](https://claude.ai/code)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

</div>

---

## What This Is

Go Sensei is a Claude Code skill that helps experienced Go developers prepare for **Go machine-round interviews**. It is not a DSA/LeetCode tool — it focuses entirely on Go-specific concepts:

- Goroutines, channels, select, WaitGroup
- sync package (Mutex, RWMutex, Once, Cond, Pool)
- Context propagation and cancellation
- Concurrency patterns (worker pool, rate limiter, errgroup, circuit breaker)
- OS signals and graceful shutdown
- Error handling idioms (errors.Is, errors.As, wrapping)
- Interfaces, generics, HTTP, io patterns
- Testing with the race detector, benchmarks, goroutine leak detection
- Memory and GC (escape analysis, sync.Pool)

**49 challenges** across 15 topic groups, sequenced from foundational to production-grade.

---

## Setup

```bash
# Clone the repo
git clone <your-repo-url>
cd golang-ai-practice

# Open Claude Code in this directory
claude .
```

That's it. The skill activates automatically through `SKILL.md` and `CLAUDE.md`.

---

## Daily Workflow

```
1. Open terminal in repo → run: claude .

2. Type:  lets solve
   → Claude finds your current ► problem, presents the challenge
      and a one-line "why this matters in production" hook.
      No hints yet — you figure out the approach first.

3. Think it through, write your Go solution in solution.go

4. Type:  here is my solution
   [paste your code]

5. Claude reviews for: race conditions, goroutine leaks,
   idiomatic Go, correct channel/context usage

6. Answer 3-5 cross-questions about WHY your solution works

7. Problem gets marked [x], ► moves to the next one

8. Optionally type:  log today's session
   → Updates PROGRESS.md with streak and notes
```

---

## Commands Reference

| What you type | What happens |
|---------------|-------------|
| `lets solve` | Presents the current ► challenge |
| `lets solve #G03` | Jump to a specific challenge by ID |
| `here is my solution` | Enter post-solve review + cross-questioning |
| `I solved it` | Same as above |
| `skip` | Marks current problem `[~]` (attempted), moves ► forward |
| `mark solved #G03` | Marks G03 as `[x]` without a full review |
| `update practice` | Shows current topic progress summary |
| `log today's session` | Updates PROGRESS.md with today's notes |
| `check solution` | Get hints on your current solution without giving the answer |

---

## Modes (Auto-detected)

Go Sensei detects what you need and switches modes automatically:

### Tutor Mode
Triggered by: "explain goroutines", "how does the scheduler work", "I don't understand sync.Cond"

Teaches Go at a senior-engineer level — GMP scheduler, runtime internals, escape analysis, GC behaviour. Not syntax tutorials.

```
You: "explain how context cancellation propagates"
→ Gets a peer-level explanation of the context tree, ctx.Done(),
  and what the runtime does when cancel() is called
```

### Hint Mode
Triggered by: "give me a hint", "I'm stuck", "don't tell me the answer"

5-level progressive hints anchored to Go's failure modes: races, leaks, deadlocks, context, channel discipline.

```
You: "stuck on the worker pool, give me a hint"
→ Hint #1: "How do you know when all workers are done?"
  (guides you to WaitGroup without naming it)
```

### Review Mode
Triggered by: "check solution", "review my code", "is this idiomatic?"

Checks for:
- Would `go test -race` pass?
- Any goroutine leaks?
- Is the channel closed by the sender only?
- Is `WaitGroup.Add` called before the goroutine starts?
- Are errors wrapped with `%w`, checked with `errors.Is`?

### Interview Mode
Triggered by: "mock interview", "practice machine round", "be the interviewer"

Simulates a Go machine-round interview at a company that runs Go in production. Evaluates concurrency correctness, idiomatic style, communication, and production thinking. Picks from Easy / Medium / Hard challenges.

```
You: "mock interview, medium difficulty"
→ Interviewer roleplay starts, presents a concurrency problem,
  evaluates approach before code, checks for race conditions,
  gives detailed feedback rubric at the end
```

### Pattern Mapper Mode
Triggered by: "what pattern should I use", "channel vs mutex here?", "what concurrency design fits this?"

Maps your problem to the right Go pattern from the full catalog: worker pool, pipeline, fan-in/out, semaphore, rate limiter, pub/sub, circuit breaker, errgroup, done channel, nil channel trick.

```
You: "I need to call 5 APIs in parallel, cancel all if any fails"
→ "This is errgroup.WithContext — here's why and how"
```

---

## Project Structure

```
golang-ai-practice/
├── SKILL.md                    # Go Sensei skill (mode router)
├── CLAUDE.md                   # Review behaviour, language rules
├── PRACTICE.md                 # 49 Go challenges with progress tracking
├── PROGRESS.md                 # Streak, session log, milestones
├── solution.go                 # Your scratch file for solutions
├── modes/
│   ├── tutor-mode.md           # Runtime internals explanations
│   ├── hint-mode.md            # Progressive hint system
│   ├── review-mode.md          # Go-specific code review checklist
│   ├── interview-mode.md       # Machine-round simulation
│   └── pattern-mapper-mode.md  # Go concurrency pattern catalog
└── docs/
    └── go-cheatsheet.md        # Quick reference: channels, sync, context, patterns
```

---

## The 49 Challenges at a Glance

| Group | Topics | Count |
|-------|--------|-------|
| Goroutines | WaitGroup, races, leaks, semaphore | 4 |
| Channels | Buffered/unbuffered, generator, pipeline, done, nil | 5 |
| Select | Non-blocking, timeout, fan-in | 3 |
| sync Package | Mutex, Once, RWMutex, Cond | 4 |
| sync/atomic | Atomic counter, CAS | 2 |
| Context | Cancel, timeout, WithValue, leak fix | 4 |
| Concurrency Patterns | Worker pool, rate limiter, errgroup, circuit breaker | 5 |
| OS Signals | SIGTERM/SIGINT, HTTP graceful shutdown | 2 |
| Error Handling | Custom errors, wrapping, panic/recover, retry | 4 |
| Interfaces | io.Reader, type switch, mock for testing | 3 |
| Generics | Stack, Map func, Set with comparable | 3 |
| HTTP & Networking | Handler timeout, middleware, retry client | 3 |
| io Patterns | Custom Writer, io.Pipe, bufio.Scanner | 3 |
| Testing | Table-driven, benchmark, race detector, goleak | 4 |
| Memory & Performance | Escape analysis, sync.Pool, slice pre-allocation | 3 |

---

## Manual Progress Edits

`PRACTICE.md` is a plain Markdown file — edit it directly any time:

- Mark solved: `[ ]` → `[x]`
- Mark attempted: `[ ]` → `[~]`
- Mark mastered: `[x]` → `[★]`
- Move `►` to whatever problem you want to work on next

---

## What Cross-Questioning Looks Like

After a correct solution, Claude asks 3-5 questions anchored to your actual code — not generic Go trivia:

- "Would `go test -race` flag the goroutine on line 12? Why?"
- "What happens if GOMAXPROCS=1 — does your solution still work?"
- "You chose a mutex here instead of a channel — when would you flip that decision?"
- "Under what condition does the goroutine you spawned on line 8 never exit?"
- "How does `database/sql`'s connection pool solve the same problem you just wrote?"

You only mark the problem solved after you've answered satisfactorily. That's what builds the instinct.

---

## License

MIT — see [LICENSE](LICENSE).
