# Valid Anagram — LeetCode 242 (E)

## 1. What is the problem?

**LeetCode 242 — Valid Anagram**

Given two strings s and t, return true if t is an anagram of s (same characters with same frequencies).

**Example 1**

- **Input:** s = "anagram", t = "nagaram"
- **Output:** true

**Example 2**

- **Input:** s = "rat", t = "car"
- **Output:** false

---

## 2. Brute Force Solution

Sort both strings and compare. If sorted(s) == sorted(t) then anagram. Correct. O(n log n) time.

### Python (brute force)

```python
def is_anagram_brute(s: str, t: str) -> bool:
    return sorted(s) == sorted(t)
```

### Golang (brute force)

```go
func isAnagramBrute(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	bs := []byte(s)
	bt := []byte(t)
	sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })
	sort.Slice(bt, func(i, j int) bool { return bt[i] < bt[j] })
	return string(bs) == string(bt)
}
```

- **Time:** O(n log n)  
- **Space:** O(n) for sorted copies

---

## 3. Optimized Solution (Hash / Count)

**Idea:** If len(s) != len(t) return false. Count frequency of each character in s (increment) and in t (decrement). If all counts zero, anagram. O(n) time.

**Steps:**

1. If len(s) != len(t): return False.  
2. count = [0]*26 (or defaultdict). For c in s: count[ord(c)-ord('a')] += 1. For c in t: count[ord(c)-ord('a')] -= 1.  
3. Return all(c == 0 for c in count) or equivalent.

### Python (optimized)

```python
def is_anagram(s: str, t: str) -> bool:
    if len(s) != len(t):
        return False
    count = [0] * 26
    for c in s:
        count[ord(c) - ord('a')] += 1
    for c in t:
        count[ord(c) - ord('a')] -= 1
    return all(c == 0 for c in count)
```

### Golang (optimized)

```go
func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var count [26]int
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}
	for _, c := range count {
		if c != 0 {
			return false
		}
	}
	return true
}
```

- **Time:** O(n)  
- **Space:** O(1) (fixed 26 or charset size)

---

## 4. Walkthrough

s = "anagram", t = "nagaram". Lengths equal. Count for s: a=3, n=1, g=1, r=1, m=1. After t: all 0. True.

---

## 5. Edge Cases

1. **Different lengths:** false.  
2. **Empty both:** true.  
3. **Unicode:** Use map instead of array if not just 'a'-'z'.  
4. **Single char:** same char → true.

---

## 6. Pitfalls

1. **Assuming lowercase:** Problem says lowercase English; use ord(c)-'a'. For full Unicode use map.  
2. **Only one loop:** Need to count both strings (or one increment, one decrement).  
3. **Sorting:** O(n log n); count is O(n).

---

## 7. Follow-ups

1. **Group anagrams (LC 49)?** Map: sorted tuple of chars (or count signature) → list of strings.  
2. **Anagram with unicode?** Use map[rune]int.  
3. **All anagram pairs?** Group by signature, then for each group with k strings, pairs = k*(k-1)/2.

---

## 8. When to Use

- **Anagram / same multiset:** Count array or map; or sort.  
- **Pattern:** Frequency count; same idea in “ransom note”, “group anagrams”.
