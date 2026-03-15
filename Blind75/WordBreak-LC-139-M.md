# Word Break — LeetCode 139 (M)

## 1. What is the problem?

**LeetCode 139 — Word Break**

Given a string s and a list of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words (each word can be reused).

**Example 1**

- **Input:** s = "leetcode", wordDict = ["leet", "code"]
- **Output:** true ("leet code")

**Example 2**

- **Input:** s = "applepenapple", wordDict = ["apple", "pen"]
- **Output:** true

---

## 2. Brute Force Solution

Recurse: try every prefix; if prefix is in wordDict, recurse on remainder. Base: empty string → true. Correct but exponential without memo.

### Python (brute)

```python
def word_break_brute(s: str, word_dict: list[str]) -> bool:
    words = set(word_dict)
    def f(s):
        if not s:
            return True
        for i in range(1, len(s) + 1):
            if s[:i] in words and f(s[i:]):
                return True
        return False
    return f(s)
```

### Golang (brute)

```go
func wordBreakBrute(s string, wordDict []string) bool {
	words := make(map[string]bool)
	for _, w := range wordDict {
		words[w] = true
	}
	var f func(string) bool
	f = func(s string) bool {
		if s == "" {
			return true
		}
		for i := 1; i <= len(s); i++ {
			if words[s[:i]] && f(s[i:]) {
				return true
			}
		}
		return false
	}
	return f(s)
}
```

- **Time:** O(2^n) without memo  
- **Space:** O(n) stack

---

## 3. Optimized Solution (DP)

**Idea:** dp[i] = can we segment s[0:i]? dp[0]=true. For each i from 1 to n, dp[i] = true if for some j < i, dp[j] and s[j:i] in wordDict.

**Steps:**

1. words = set(wordDict). dp = [True] + [False]*len(s).  
2. For i from 1 to len(s): for j from 0 to i-1: if dp[j] and s[j:i] in words: dp[i]=True; break.  
3. Return dp[len(s)].

### Python (optimized)

```python
def word_break(s: str, word_dict: list[str]) -> bool:
    words = set(word_dict)
    n = len(s)
    dp = [False] * (n + 1)
    dp[0] = True
    for i in range(1, n + 1):
        for j in range(i):
            if dp[j] and s[j:i] in words:
                dp[i] = True
                break
    return dp[n]
```

### Golang (optimized)

```go
func wordBreak(s string, wordDict []string) bool {
	words := make(map[string]bool)
	for _, w := range wordDict {
		words[w] = true
	}
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && words[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
```

- **Time:** O(n² * m) where m = max word length for substring; with set lookup O(1), O(n²).  
- **Space:** O(n)

---

## 4. Walkthrough

s = "leetcode", words = {"leet","code"}. dp[0]=T. dp[4]=T (s[0:4]=leet); dp[8]=T (s[4:8]=code). Return True.

---

## 5. Edge Cases

1. **Empty s:** true (dp[0]).  
2. **Word longer than s:** Skip in inner loop.  
3. **Repeated words:** Set handles.  
4. **s = "a", dict = ["a"]:** dp[1]=T.

---

## 6. Pitfalls

1. **dp[0]=true:** Empty prefix is valid.  
2. **Substring slice:** s[j:i] in Python/Go is correct.  
3. **Trie optimization:** For long wordDict, check only prefixes that exist in trie.

---

## 7. Follow-ups

1. **Word Break II (return all sentences)?** DP + backtrack to build paths.  
2. **Concatenated words?** Build graph of suffixes; similar DP.  
3. **Space optimization?** Only need dp for positions that are true and extend; BFS of reachable indices.

---

## 8. When to Use

- **Segment string from dictionary:** DP over prefixes; dp[i] from dp[j] + s[j:i] in dict.  
- **Pattern:** String DP; same idea in “decode ways”, “palindrome partition”.
