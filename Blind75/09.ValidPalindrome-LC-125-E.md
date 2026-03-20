# Valid Palindrome — LeetCode 125 (E)

## 1. What is the problem?

**LeetCode 125 — Valid Palindrome**

Given a string s, determine if it is a palindrome. Consider only alphanumeric characters and ignore cases.

**Example 1**

- **Input:** s = "A man, a plan, a canal: Panama"
- **Output:** true
- **Explanation:** "amanaplanacanalpanama" is a palindrome.

**Example 2**

- **Input:** s = "race a car"
- **Output:** false
- **Explanation:** "raceacar" is not a palindrome.

---

## 2. Brute Force Solution

Filter to alphanumeric only, lowercase. Check if reversed equals original. O(n) time, O(n) space.

### Python (two pointers — O(1) extra space)

```python
def is_palindrome(s: str) -> bool:
    def alnum(c):
        return c.isalnum()
    l, r = 0, len(s) - 1
    while l < r:
        while l < r and not alnum(s[l]):
            l += 1
        while l < r and not alnum(s[r]):
            r -= 1
        if l < r and s[l].lower() != s[r].lower():
            return False
        l += 1
        r -= 1
    return True
```

### Golang

```go
func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		for l < r && !isAlnum(s[l]) {
			l++
		}
		for l < r && !isAlnum(s[r]) {
			r--
		}
		if l < r && toLower(s[l]) != toLower(s[r]) {
			return false
		}
		l++
		r--
	}
	return true
}
func isAlnum(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
}
func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 32
	}
	return b
}
```

- **Time:** O(n)  
- **Space:** O(1)

---

## 4. Walkthrough

"A man, a plan..." → skip non-alnum, compare 'a' vs 'a', 'm' vs 'm', etc. Two pointers meet in middle.

---

## 5. Edge Cases

1. **Empty string:** true.  
2. **Single char:** true.  
3. **Only non-alnum:** true (empty after filter).  
4. **Numbers:** "0P" — '0' vs 'P', false.

---

## 6. Pitfalls

1. **Case:** Use .lower() or toLower.  
2. **Skip non-alnum:** Move pointers until alnum.  
3. **Bounds:** l < r after skips.

---

## 7. Follow-ups

1. **Valid Palindrome II (remove at most 1 char)?** Try skip left or right when mismatch.  
2. **Strict (no skip)?** Just reverse and compare.  
3. **Unicode?** Use proper normalization.

---

## 8. When to Use

- **Check palindrome:** Two pointers from both ends; compare.  
- **Pattern:** Same in "Longest Palindromic Substring" (expand from center).
