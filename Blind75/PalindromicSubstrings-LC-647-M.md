# Palindromic Substrings — LeetCode 647 (M)

## 1. What is the problem?

**LeetCode 647 — Palindromic Substrings**

Given a string s, return the number of palindromic substrings in it.

**Example 1**

- **Input:** s = "abc"
- **Output:** 3
- **Explanation:** Three palindromic strings: "a", "b", "c".

**Example 2**

- **Input:** s = "aaa"
- **Output:** 6
- **Explanation:** Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".

---

## 2. Brute Force Solution

For each (i,j), check if s[i:j+1] is palindrome. Count. O(n³).

### Python (expand around center)

**Idea:** For each center (odd and even), expand and count palindromes. Each expansion adds 1 to count.

```python
def count_substrings(s: str) -> int:
    def expand(l, r):
        count = 0
        while l >= 0 and r < len(s) and s[l] == s[r]:
            count += 1
            l -= 1
            r += 1
        return count
    total = 0
    for i in range(len(s)):
        total += expand(i, i)
        total += expand(i, i + 1)
    return total
```

### Golang

```go
func countSubstrings(s string) int {
	expand := func(l, r int) int {
		count := 0
		for l >= 0 && r < len(s) && s[l] == s[r] {
			count++
			l--
			r++
		}
		return count
	}
	total := 0
	for i := 0; i < len(s); i++ {
		total += expand(i, i)
		total += expand(i, i+1)
	}
	return total
}
```

- **Time:** O(n²)  
- **Space:** O(1)

---

## 4. Walkthrough

s="aaa". Centers: (0,0)→1, (0,1)→1, (1,1)→1, (1,2)→1, (2,2)→1. (0,1) expand: aaa→1 more. Total 6.

---

## 5. Edge Cases

1. **Single char:** 1.  
2. **All same:** n*(n+1)/2.  
3. **No repeated:** n (each char).  
4. **Empty:** 0.

---

## 6. Pitfalls

1. **Odd + even:** Both contribute.  
2. **Count per expansion:** Each valid (l,r) is one substring.  
3. **Overlap:** Different centers can produce same substring? No — each center produces unique set.

---

## 7. Follow-ups

1. **Longest palindromic substring?** Same expand; track max len.  
2. **Distinct palindromic substrings?** Store in set.  
3. **DP approach?** dp[i][j] = is s[i:j+1] palindrome; O(n²) time and space.

---

## 8. When to Use

- **Count palindromic substrings:** Expand around center; count each.  
- **Pattern:** Same as "Longest Palindromic Substring" but count instead of max.
