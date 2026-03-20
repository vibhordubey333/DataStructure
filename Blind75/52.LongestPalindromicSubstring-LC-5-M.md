# Longest Palindromic Substring — LeetCode 5 (M)

## 1. What is the problem?

**LeetCode 5 — Longest Palindromic Substring**

Given a string s, return the longest palindromic substring in s.

**Example 1**

- **Input:** s = "babad"
- **Output:** "bab"
- **Explanation:** "aba" is also a valid answer.

**Example 2**

- **Input:** s = "cbbd"
- **Output:** "bb"

---

## 2. Brute Force Solution

For each (i,j), check if s[i:j+1] is palindrome. Track max. O(n³).

### Python (expand around center)

**Idea:** For each center (single char or between two chars), expand outward while palindrome. Track longest.

```python
def longest_palindrome(s: str) -> str:
    def expand(l, r):
        while l >= 0 and r < len(s) and s[l] == s[r]:
            l -= 1
            r += 1
        return s[l + 1:r]
    best = ""
    for i in range(len(s)):
        odd = expand(i, i)
        even = expand(i, i + 1)
        best = max(best, odd, even, key=len)
    return best
```

### Golang

```go
func longestPalindrome(s string) string {
	expand := func(l, r int) string {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		return s[l+1 : r]
	}
	best := ""
	for i := 0; i < len(s); i++ {
		odd := expand(i, i)
		even := expand(i, i+1)
		if len(odd) > len(best) {
			best = odd
		}
		if len(even) > len(best) {
			best = even
		}
	}
	return best
}
```

- **Time:** O(n²)  
- **Space:** O(1) excluding output

---

## 4. Walkthrough

s="babad". Center 0: "b". Center 1: "bab". Center 2: "aba". Center 3: "a". Center 4: "d". Best = "bab" or "aba".

---

## 5. Edge Cases

1. **Single char:** return s.  
2. **All same:** "aaa" → "aaa".  
3. **No palindrome > 1:** return first char.  
4. **Empty:** return "".

---

## 6. Pitfalls

1. **Odd vs even:** Odd: expand(i,i). Even: expand(i,i+1).  
2. **Expand returns:** s[l+1:r] because l,r are one past after loop.  
3. **Manacher's:** O(n) but more complex; expand is usually enough.

---

## 7. Follow-ups

1. **Palindromic Substrings (count)?** Same expand; count each.  
2. **Longest palindromic subsequence?** DP.  
3. **Manacher's algorithm?** O(n) with radius array.

---

## 8. When to Use

- **Longest palindrome substring:** Expand around center (odd/even).  
- **Pattern:** Same idea in "Valid Palindrome", "Palindromic Substrings".
