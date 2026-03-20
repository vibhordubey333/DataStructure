# Longest Common Subsequence — LeetCode 1143 (M)

## 1. What is the problem?

**LeetCode 1143 — Longest Common Subsequence**

Given two strings `text1` and `text2`, return the length of their longest common subsequence (LCS). A subsequence is characters that appear in order but not necessarily contiguous.

**Example 1**

- **Input:** text1 = "abcde", text2 = "ace"
- **Output:** 3 (LCS "ace")

**Example 2**

- **Input:** text1 = "abc", text2 = "abc"
- **Output:** 3

---

## 2. Brute Force Solution

Recurse: if text1[i]==text2[j], 1+LCS(i+1,j+1); else max(LCS(i+1,j), LCS(i,j+1)). Exponential without memo.

### Python (brute force)

```python
def longest_common_subsequence_brute(text1: str, text2: str) -> int:
    def f(i: int, j: int) -> int:
        if i >= len(text1) or j >= len(text2):
            return 0
        if text1[i] == text2[j]:
            return 1 + f(i + 1, j + 1)
        return max(f(i + 1, j), f(i, j + 1))
    return f(0, 0)
```

### Golang (brute force)

```go
func longestCommonSubsequenceBrute(text1, text2 string) int {
	var f func(int, int) int
	f = func(i, j int) int {
		if i >= len(text1) || j >= len(text2) {
			return 0
		}
		if text1[i] == text2[j] {
			return 1 + f(i+1, j+1)
		}
		a, b := f(i+1, j), f(i, j+1)
		if a > b {
			return a
		}
		return b
	}
	return f(0, 0)
}
```

- **Time:** O(2^(m+n))  
- **Space:** O(m+n) stack

---

## 3. Optimized Solution (DP)

**Idea:** dp[i][j] = LCS of text1[0:i] and text2[0:j]. If text1[i-1]==text2[j-1]: dp[i][j]=1+dp[i-1][j-1]. Else: dp[i][j]=max(dp[i-1][j], dp[i][j-1]). Can use 1D dp to save space.

### Python (optimized)

```python
def longest_common_subsequence(text1: str, text2: str) -> int:
    m, n = len(text1), len(text2)
    dp = [[0] * (n + 1) for _ in range(m + 1)]
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if text1[i - 1] == text2[j - 1]:
                dp[i][j] = 1 + dp[i - 1][j - 1]
            else:
                dp[i][j] = max(dp[i - 1][j], dp[i][j - 1])
    return dp[m][n]
```

### Golang (optimized)

```go
func longestCommonSubsequence(text1, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				a, b := dp[i-1][j], dp[i][j-1]
				if a > b {
					dp[i][j] = a
				} else {
					dp[i][j] = b
				}
			}
		}
	}
	return dp[m][n]
}
```

- **Time:** O(m * n)  
- **Space:** O(m * n); O(min(m,n)) with 1D optimization

---

## 4. Walkthrough

text1="abcde", text2="ace". dp[1][1]=1 (a==a); dp[2][2]=1 (b!=c, max(0,1)=1); dp[3][3]=2 (c==c); dp[5][3]=3. Return 3.

---

## 5. Edge Cases

1. **Empty string:** return 0.  
2. **One char match:** 1.  
3. **No common:** 0.  
4. **Identical strings:** length of string.

---

## 6. Pitfalls

1. **Index off-by-one:** dp[i][j] uses text1[i-1], text2[j-1].  
2. **Subsequence vs subarray:** Subsequence can skip; subarray must be contiguous.  
3. **Space optimization:** Use prev row only; swap arrays.

---

## 7. Follow-ups

1. **Return the LCS string?** Backtrack from dp[m][n] using parent pointers.  
2. **Longest common subarray?** Different: dp[i][j]=1+dp[i-1][j-1] if match else 0.  
3. **Edit distance (Levenshtein)?** Similar 2D DP with insert/delete/replace.

---

## 8. When to Use

- **LCS / sequence alignment:** Classic 2D DP; dp[i][j] from dp[i-1][j-1], dp[i-1][j], dp[i][j-1].  
- **Pattern:** Same idea in "edit distance", "shortest common supersequence".
