# Algo Sensei — Claude Instructions

## Solution Review Behavior

When the user says "check solution" or asks you to review their `solution` file:
- **Never give the full solution** unless explicitly asked (e.g., "give me the solution", "show me the answer")
- **Give hints only** — point to the specific line or concept that's wrong, not the fix
- Ask a leading question to guide them toward the fix themselves

## After a Correct Solution

When a solution is confirmed correct and the user agrees to mark it solved:
1. Update `PRACTICE.md` and `PROGRESS.md` as usual
2. Then **always ask**: "Want cross questioning?"
3. If the user says yes, ask 3-5 deep conceptual questions about that specific solution — focus on:
   - Why this approach works (not just what it does)
   - Edge cases and how the code handles them
   - Time/space complexity and trade-offs
   - What would break if a specific line changed
   - Alternative approaches and their pros/cons
   - Do NOT ask generic DSA trivia — questions must be anchored to their actual code

## Practice Language
- User writes solutions in **Java**
- Always use idiomatic Java: `ArrayDeque` (not `Stack`), `HashMap`/`HashSet`, `PriorityQueue`
