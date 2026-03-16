# Longest Repeating Character Replacement — LeetCode 424 (M)

## 1. What is the problem?

**LeetCode 424 — Longest Repeating Character Replacement**

Given a string s and an integer k, you can change at most k characters to any letter. Return the length of the longest substring containing the same letter after performing at most k replacements.

**Example 1**

- **Input:** s = "ABAB", k = 2
- **Output:** 4
- **Explanation:** Replace the two 'A's with two 'B's or vice versa.

**Example 2**

- **Input:** s = "AABABBA", k = 1
- **Output:** 4
- **Explanation:** Replace the one 'A' in the middle with 'B' to get "AABBBBA". Substring "BBBB" has length 4.

---

## 2. Brute Force Solution

For each (i,j), count freq of each char in s[i:j+1]. If len - max_freq <= k, valid. Track max len. O(n² * 26).

### Python (optimized — sliding window)

**Idea:** Sliding window. For window [l,r], max_freq = max count of any char. If window_len - max_freq <= k, valid; else shrink l. Track max window size.

```python
def character_replacement(s: str, k: int) -> int:
    freq = {}
    max_freq = 0
    best = 0
    l = 0
    for r, c in enumerate(s):
        freq[c] = freq.get(c, 0) + 1
        max_freq = max(max_freq, freq[c])
        while (r - l + 1) - max_freq > k:
            freq[s[l]] -= 1
            l += 1
        best = max(best, r - l + 1)
    return best
```

### Golang

```go
func characterReplacement(s string, k int) int {
	freq := make(map[byte]int)
	maxFreq, best, l := 0, 0, 0
	for r := 0; r < len(s); r++ {
		freq[s[r]]++
		if freq[s[r]] > maxFreq {
			maxFreq = freq[s[r]]
		}
		for (r-l+1)-maxFreq > k {
			freq[s[l]]--
			l++
		}
		if r-l+1 > best {
			best = r - l + 1
		}
	}
	return best
}
```

- **Time:** O(n)  
- **Space:** O(26) = O(1)

---

## 4. Walkthrough

s="AABABBA", k=1. Window [0,3]: AABA, max_freq=2, len-max=2>1, shrink. [1,4]: ABAB, max=2, len-max=2>1. [2,5]: BABB, max=3, len-max=1<=1, valid, best=4. Continue.

---

## 5. Edge Cases

1. **k >= len(s):** return len(s).  
2. **k=0:** longest run of same char.  
3. **Single char:** return 1.  
4. **Empty s:** return 0.

---

## 6. Pitfalls

1. **max_freq doesn't decrease when shrinking:** We only care about max valid window; not updating max_freq is OK because we only shrink when invalid.  
2. **Window validity:** len - max_freq <= k means we can replace (len - max_freq) chars to make all same.

---

## 7. Follow-ups

1. **Return the substring?** Track best (l,r) and return s[l:r+1].  
2. **At most k different chars?** Different problem (Longest Substring with At Most K Distinct).  
3. **Case insensitive?** Normalize to lower.

---

## 8. When to Use

- **Sliding window + char freq:** Valid when (window_len - max_freq) <= k.  
- **Pattern:** Same as "Minimum Window Substring" style; expand, shrink when invalid.
