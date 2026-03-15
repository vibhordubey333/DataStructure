# Longest Substring Without Repeating Characters — LeetCode 3 (M)

## 1. What is the problem?

**LeetCode 3 — Longest Substring Without Repeating Characters**

Given a string s, find the length of the longest substring without repeating characters.

**Example 1**

- **Input:** s = "abcabcbb"
- **Output:** 3 ("abc")

**Example 2**

- **Input:** s = "bbbbb"
- **Output:** 1

---

## 2. Brute Force Solution

For each starting index i, extend j until we see a duplicate; track max length. Check duplicate with a set for substring s[i:j+1]. O(n²) or O(n³) depending on duplicate check.

### Python (brute force)

```python
def length_of_longest_substring_brute(s: str) -> int:
    n = len(s)
    best = 0
    for i in range(n):
        seen = set()
        for j in range(i, n):
            if s[j] in seen:
                break
            seen.add(s[j])
            best = max(best, j - i + 1)
    return best
```

### Golang (brute force)

```go
func lengthOfLongestSubstringBrute(s string) int {
	best := 0
	for i := 0; i < len(s); i++ {
		seen := make(map[byte]bool)
		for j := i; j < len(s); j++ {
			if seen[s[j]] {
				break
			}
			seen[s[j]] = true
			if j-i+1 > best {
				best = j - i + 1
			}
		}
	}
	return best
}
```

- **Time:** O(n²)  
- **Space:** O(min(n, charset))

---

## 3. Optimized Solution (Sliding Window + Hash)

**Idea:** Maintain a window [left, right] with no duplicate. For each right, if s[right] is already in the window, advance left until the previous occurrence of s[right] is excluded; then add s[right] and update max length.

**Steps:**

1. left = 0; seen = {} (char → index); best = 0.  
2. For right in 0..n-1: If s[right] in seen and seen[s[right]] >= left: left = seen[s[right]] + 1. seen[s[right]] = right; best = max(best, right - left + 1).  
3. Return best.

### Python (optimized)

```python
def length_of_longest_substring(s: str) -> int:
    seen = {}
    left = best = 0
    for right, c in enumerate(s):
        if c in seen and seen[c] >= left:
            left = seen[c] + 1
        seen[c] = right
        best = max(best, right - left + 1)
    return best
```

### Golang (optimized)

```go
func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]int)
	left, best := 0, 0
	for right := 0; right < len(s); right++ {
		c := s[right]
		if idx, ok := seen[c]; ok && idx >= left {
			left = idx + 1
		}
		seen[c] = right
		if right-left+1 > best {
			best = right - left + 1
		}
	}
	return best
}
```

- **Time:** O(n)  
- **Space:** O(min(n, charset))

---

## 4. Walkthrough

s = "abcabcbb". right=0: a, left=0, best=1. right=1: b, best=2. right=2: c, best=3. right=3: a seen at 0, left=1, best=3. right=4: b seen at 1, left=2, best=3. right=5: c seen at 2, left=3, best=3. right=6: b seen at 4, left=5, best=3. right=7: b seen at 6, left=7, best=3. Output 3.

---

## 5. Edge Cases

1. **Empty string:** 0.  
2. **Single character:** 1.  
3. **All same:** 1.  
4. **No duplicate:** entire string length.  
5. **seen[c] >= left:** Only shrink when duplicate is inside current window.

---

## 6. Pitfalls

1. **Only storing “in set”:** We need last index of each char to move left correctly.  
2. **Moving left to right+1:** left should move to seen[c]+1 so the duplicate is excluded.  
3. **Updating seen after updating left:** seen[c]=right is correct; left uses old index.

---

## 7. Follow-ups

1. **Longest substring with at most k distinct?** Sliding window + map count; shrink when distinct > k.  
2. **Longest with at most 2 repeating?** Count of each char in window; shrink when any count > 2.  
3. **Return the substring?** Track start of best window (left when we update best).

---

## 8. When to Use

- **Longest substring satisfying condition:** Sliding window + hash map for “last seen” or counts.  
- **Pattern:** Variable-size window; “no repeat” → map char to last index.
