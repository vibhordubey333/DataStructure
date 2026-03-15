# Counting Bits — LeetCode 338 (E)

## 1. What is the problem?

**LeetCode 338 — Counting Bits**

Given an integer n, return an array ans of length n+1 such that for each i (0 ≤ i ≤ n), ans[i] is the number of 1's in the binary representation of i.

**Example 1**

- **Input:** n = 2
- **Output:** [0, 1, 1]

**Example 2**

- **Input:** n = 5
- **Output:** [0, 1, 1, 2, 1, 2]

---

## 2. Brute Force Solution

For each i from 0 to n, count set bits (e.g. loop and do i & 1, i >>= 1). Push count to result. Correct by definition.

### Python (brute force)

```python
def count_bits_brute(n: int) -> list[int]:
    def popcount(x):
        c = 0
        while x:
            c += x & 1
            x >>= 1
        return c
    return [popcount(i) for i in range(n + 1)]
```

### Golang (brute force)

```go
func countBitsBrute(n int) []int {
	out := make([]int, n+1)
	for i := 0; i <= n; i++ {
		x := i
		for x > 0 {
			out[i] += x & 1
			x >>= 1
		}
	}
	return out
}
```

- **Time:** O(n * bit_length) = O(n log n) for unbounded bits; for 32-bit O(32n).  
- **Space:** O(1) excluding output

---

## 3. Optimized Solution (DP)

**Idea:** ans[i] = ans[i >> 1] + (i & 1). The number of 1's in i equals the count in i/2 plus the last bit. (Or: ans[i] = ans[i & (i-1)] + 1 — clear lowest set bit and add 1.)

**Steps:**

1. ans = [0] * (n+1).  
2. For i from 1 to n: ans[i] = ans[i >> 1] + (i & 1).  
3. Return ans.

### Python (optimized)

```python
def count_bits(n: int) -> list[int]:
    ans = [0] * (n + 1)
    for i in range(1, n + 1):
        ans[i] = ans[i >> 1] + (i & 1)
    return ans
```

### Golang (optimized)

```go
func countBits(n int) []int {
	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ans[i] = ans[i>>1] + (i & 1)
	}
	return ans
}
```

- **Time:** O(n)  
- **Space:** O(1) excluding output

---

## 4. Walkthrough

n = 5. ans[0]=0; ans[1]=ans[0]+1=1; ans[2]=ans[1]+0=1; ans[3]=ans[1]+1=2; ans[4]=ans[2]+0=1; ans[5]=ans[2]+1=2. Result: [0,1,1,2,1,2].

---

## 5. Edge Cases

1. **n = 0:** return [0].  
2. **n = 1:** return [0, 1].  
3. **Powers of two:** ans[2^k] = 1.  
4. **Large n:** DP is still O(n).

---

## 6. Pitfalls

1. **Starting loop at 0:** ans[0]=0 is correct; formula applies from i=1.  
2. **Wrong recurrence:** ans[i] = ans[i/2] + i%2 (integer division).  
3. **Off-by-one:** ans has length n+1, indices 0..n.

---

## 7. Follow-ups

1. **Only count for a single number?** Use popcount (builtin or i & (i-1) trick).  
2. **Range query [a,b]?** Precompute prefix sum of ans.  
3. **Return count of numbers with k set bits in [0,n]?** Digit DP or combinatorial count.

---

## 8. When to Use

- **Popcount for 0..n:** DP recurrence ans[i] = ans[i>>1] + (i&1).  
- **Pattern:** Linear DP with bit recurrence; also “clear lowest set bit” recurrence.
