# Climbing Stairs — LeetCode 70 (E)

## 1. What is the problem?

**LeetCode 70 — Climbing Stairs**

You are climbing a staircase. It takes n steps to reach the top. Each time you can either climb 1 or 2 steps. How many distinct ways can you climb to the top?

**Example 1**

- **Input:** n = 2
- **Output:** 2 (1+1 or 2)

**Example 2**

- **Input:** n = 3
- **Output:** 3 (1+1+1, 1+2, 2+1)

---

## 2. Brute Force Solution

Recurse: ways(n) = ways(n-1) + ways(n-2) with base cases ways(0)=1, ways(1)=1. Correct but repeats work (exponential time).

### Python (brute force)

```python
def climb_stairs_brute(n: int) -> int:
    if n <= 1:
        return 1
    return climb_stairs_brute(n - 1) + climb_stairs_brute(n - 2)
```

### Golang (brute force)

```go
func climbStairsBrute(n int) int {
	if n <= 1 {
		return 1
	}
	return climbStairsBrute(n-1) + climbStairsBrute(n-2)
}
```

- **Time:** O(2^n)  
- **Space:** O(n) stack

---

## 3. Optimized Solution (DP)

**Idea:** Same recurrence: f(i) = f(i-1) + f(i-2). Compute from 1 to n with two variables (like Fibonacci).

**Steps:**

1. If n <= 1 return 1.  
2. prev, cur = 1, 1. For i from 2 to n: prev, cur = cur, prev + cur.  
3. Return cur.

### Python (optimized)

```python
def climb_stairs(n: int) -> int:
    if n <= 1:
        return 1
    prev, cur = 1, 1
    for _ in range(2, n + 1):
        prev, cur = cur, prev + cur
    return cur
```

### Golang (optimized)

```go
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	prev, cur := 1, 1
	for i := 2; i <= n; i++ {
		prev, cur = cur, prev+cur
	}
	return cur
}
```

- **Time:** O(n)  
- **Space:** O(1)

---

## 4. Walkthrough

n = 3: prev=1, cur=1 → prev=1, cur=2 → prev=2, cur=3. Return 3.

---

## 5. Edge Cases

1. **n = 0 or 1:** 1 way.  
2. **n = 2:** 2 ways.  
3. **Large n:** Use iterative DP to avoid overflow (use bigint if needed).

---

## 6. Pitfalls

1. **Base case:** ways(0)=1 (one way: don’t move).  
2. **Off-by-one:** f(n) is the answer for n steps.  
3. **Recursion only:** Mention memoization or iterative to get O(n).

---

## 7. Follow-ups

1. **k steps at a time?** f(i) = f(i-1)+…+f(i-k).  
2. **Different cost per step?** Min cost to reach top (LC 746).  
3. **Print all paths?** Backtrack and record steps.

---

## 8. When to Use

- **Linear recurrence (Fibonacci-like):** Iterate with O(1) state.  
- **Pattern:** Classic 1D DP; same as Fibonacci.
