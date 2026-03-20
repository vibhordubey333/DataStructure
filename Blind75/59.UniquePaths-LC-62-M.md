# Unique Paths — LeetCode 62 (M)

## 1. What is the problem?

**LeetCode 62 — Unique Paths**

A robot is at the top-left of an m×n grid. It can move only right or down. How many unique paths to reach bottom-right?

**Example 1**

- **Input:** m = 3, n = 7
- **Output:** 28

**Example 2**

- **Input:** m = 3, n = 2
- **Output:** 3

---

## 2. Brute Force Solution

Recurse: paths(i,j) = paths(i+1,j) + paths(i,j+1) with base at (m-1,n-1)=1. Exponential.

### Python (brute force)

```python
def unique_paths_brute(m: int, n: int) -> int:
    def f(i: int, j: int) -> int:
        if i >= m or j >= n:
            return 0
        if i == m - 1 and j == n - 1:
            return 1
        return f(i + 1, j) + f(i, j + 1)
    return f(0, 0)
```

### Golang (brute force)

```go
func uniquePathsBrute(m, n int) int {
	var f func(int, int) int
	f = func(i, j int) int {
		if i >= m || j >= n {
			return 0
		}
		if i == m-1 && j == n-1 {
			return 1
		}
		return f(i+1, j) + f(i, j+1)
	}
	return f(0, 0)
}
```

- **Time:** O(2^(m+n))  
- **Space:** O(m+n) stack

---

## 3. Optimized Solution (DP)

**Idea:** dp[i][j] = paths from (i,j) to (m-1,n-1). dp[m-1][n-1]=1; dp[i][j]=dp[i+1][j]+dp[i][j+1]. Or: combinatorial — need (m-1) downs and (n-1) rights: C(m+n-2, m-1).

### Python (optimized)

```python
def unique_paths(m: int, n: int) -> int:
    dp = [[1] * n for _ in range(m)]
    for i in range(m - 2, -1, -1):
        for j in range(n - 2, -1, -1):
            dp[i][j] = dp[i + 1][j] + dp[i][j + 1]
    return dp[0][0]
```

### Golang (optimized)

```go
func uniquePaths(m, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = 1
		}
	}
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			dp[i][j] = dp[i+1][j] + dp[i][j+1]
		}
	}
	return dp[0][0]
}
```

- **Time:** O(m * n)  
- **Space:** O(m * n); O(n) with 1D

---

## 4. Walkthrough

m=3, n=2. Last row and col = 1. dp[1][0]=dp[2][0]+dp[1][1]=1+1=2. dp[0][0]=dp[1][0]+dp[0][1]=2+1=3. Return 3.

---

## 5. Edge Cases

1. **1×1:** 1.  
2. **1×n or m×1:** 1.  
3. **Large m,n:** Use combinatorial or DP; watch overflow.

---

## 6. Pitfalls

1. **Initialization:** Last row and col are 1 (one path to finish).  
2. **Direction:** Fill from bottom-right to top-left.  
3. **Unique Paths II (obstacles):** dp[i][j]=0 if obstacle.

---

## 7. Follow-ups

1. **Obstacles (LC 63)?** dp[i][j]=0 if grid[i][j]==1.  
2. **Min path sum?** Add grid values; take min of down/right.  
3. **Combinatorial formula:** C(m+n-2, m-1) = (m+n-2)!/((m-1)!(n-1)!).

---

## 8. When to Use

- **Grid paths (right/down):** 2D DP or combinatorial.  
- **Pattern:** Classic grid DP; same in "min path sum", "unique paths II".
