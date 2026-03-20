# Minimum Window Substring — LeetCode 76 (H)

## 1. What is the problem?

**LeetCode 76 — Minimum Window Substring**

Given two strings s and t, return the minimum window substring of s such that every character in t (including duplicates) is included. If no such substring exists, return "".

**Example 1**

- **Input:** s = "ADOBECODEBANC", t = "ABC"
- **Output:** "BANC"
- **Explanation:** The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.

**Example 2**

- **Input:** s = "a", t = "a"
- **Output:** "a"

---

## 2. Brute Force Solution

For each (i,j), check if s[i:j+1] contains all chars of t with correct counts. O(n² * m). Optimize with sliding window.

### Python (sliding window)

```python
def min_window(s: str, t: str) -> str:
    if not t or not s:
        return ""
    need = {}
    for c in t:
        need[c] = need.get(c, 0) + 1
    have = 0
    required = len(need)
    freq = {}
    best_len = float('inf')
    best_start = 0
    l = 0
    for r, c in enumerate(s):
        if c in need:
            freq[c] = freq.get(c, 0) + 1
            if freq[c] == need[c]:
                have += 1
        while have == required:
            if r - l + 1 < best_len:
                best_len = r - l + 1
                best_start = l
            if s[l] in need:
                freq[s[l]] -= 1
                if freq[s[l]] < need[s[l]]:
                    have -= 1
            l += 1
    return s[best_start:best_start + best_len] if best_len != float('inf') else ""
```

### Golang

```go
func minWindow(s string, t string) string {
	if len(t) == 0 || len(s) == 0 {
		return ""
	}
	need := make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}
	required := len(need)
	freq := make(map[byte]int)
	have := 0
	bestLen, bestStart := len(s)+1, 0
	l := 0
	for r := 0; r < len(s); r++ {
		c := s[r]
		if need[c] > 0 {
			freq[c]++
			if freq[c] == need[c] {
				have++
			}
		}
		for have == required {
			if r-l+1 < bestLen {
				bestLen = r - l + 1
				bestStart = l
			}
			if need[s[l]] > 0 {
				freq[s[l]]--
				if freq[s[l]] < need[s[l]] {
					have--
				}
			}
			l++
		}
	}
	if bestLen > len(s) {
		return ""
	}
	return s[bestStart : bestStart+bestLen]
}
```

- **Time:** O(n + m)  
- **Space:** O(1) (bounded by alphabet)

---

## 4. Walkthrough

Expand r until window has all chars of t. When valid, shrink l and track min window. When have < required, stop shrinking.

---

## 5. Edge Cases

1. **t longer than s:** return "".  
2. **t has duplicate chars:** Count matters.  
3. **s = t:** return s.  
4. **No valid window:** return "".

---

## 6. Pitfalls

1. **have vs required:** have counts how many distinct chars in t are satisfied; required = len(unique chars in t).  
2. **Shrink condition:** Shrink while have == required.  
3. **Substring vs subsequence:** Must be contiguous.

---

## 7. Follow-ups

1. **All anagrams of t in s?** Sliding window of len(t); check if freq matches.  
2. **Minimum size k?** Add constraint.  
3. **Multiple t?** Return first/min.

---

## 8. When to Use

- **Minimum window containing all chars (with counts):** Sliding window + freq maps.  
- **Pattern:** Expand until valid; shrink to minimize; track best.
