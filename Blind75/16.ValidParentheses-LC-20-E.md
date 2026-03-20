# Valid Parentheses — LeetCode 20 (E)

## 1. What is the problem?

**LeetCode 20 — Valid Parentheses**

Given a string s containing only '(', ')', '{', '}', '[', ']', determine if the input is valid: open brackets closed by same type in correct order.

**Example 1**

- **Input:** s = "()"
- **Output:** true

**Example 2**

- **Input:** s = "()[]{}"
- **Output:** true

**Example 3**

- **Input:** s = "(]"
- **Output:** false

---

## 2. Brute Force Solution

Repeatedly remove adjacent matching pairs "()", "[]", "{}" until no change or string empty. If empty, valid. Correct but O(n²) for naive string replacement.

### Python (brute — replace)

```python
def is_valid_brute(s: str) -> bool:
    while "()" in s or "[]" in s or "{}" in s:
        s = s.replace("()", "").replace("[]", "").replace("{}", "")
    return len(s) == 0
```

### Golang (stack — standard)

Stack is the natural approach; “brute” could be recursive: match first open with a close and recurse. Stack is simpler and optimal.

---

## 3. Optimized Solution (Stack)

**Idea:** For each character: if open bracket, push; if close bracket, pop and check it matches. At end stack must be empty.

**Steps:**

1. stack = []. Map: closing → opening.  
2. For c in s: If c is open: push c. Else: if stack empty or pop != map[c]: return False.  
3. Return len(stack) == 0.

### Python (optimized)

```python
def is_valid(s: str) -> bool:
    stack = []
    pair = {')': '(', ']': '[', '}': '{'}
    for c in s:
        if c in '([{':
            stack.append(c)
        else:
            if not stack or stack.pop() != pair[c]:
                return False
    return len(stack) == 0
```

### Golang (optimized)

```go
func isValid(s string) bool {
	stack := []byte{}
	pair := map[byte]byte{')': '(', ']': '[', '}': '{'}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != pair[c] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
```

- **Time:** O(n)  
- **Space:** O(n) stack

---

## 4. Walkthrough

s = "()[]{}". ( push; ) pop, match. [ push; ] pop, match. { push; } pop, match. Stack empty → true.

s = "(]". ( push; ] pop → [ != pair[']'] → false.

---

## 5. Edge Cases

1. **Empty string:** true.  
2. **Only open:** stack not empty → false.  
3. **Only close:** stack empty when we try to pop → false.  
4. **Odd length:** false.  
5. **Nested:** stack handles correctly.

---

## 6. Pitfalls

1. **Check stack empty before pop:** Popping when empty is wrong.  
2. **Wrong mapping:** ')' matches '(', not '['.  
3. **Using one counter:** One counter can’t handle mixed types (e.g. "([)]" is invalid).

---

## 7. Follow-ups

1. **Generate all valid parentheses (LC 22)?** Backtrack: add '(' or ')' keeping count valid.  
2. **Longest valid parentheses (LC 32)?** Stack with indices or DP.  
3. **Min add to make valid (LC 921)?** Count unmatched open and close.

---

## 8. When to Use

- **Matching nested brackets:** Stack; push open, pop on close and check match.  
- **Pattern:** Classic stack; same in expression evaluation, XML/HTML tags.
