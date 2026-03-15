# Group Anagrams — LeetCode 49 (M)

## 1. What is the problem?

**LeetCode 49 — Group Anagrams**

Given an array of strings strs, group the anagrams together. An anagram is a string with the same characters and frequencies. Return answer in any order.

**Example 1**

- **Input:** strs = ["eat","tea","tan","ate","nat","bat"]
- **Output:** [["bat"],["nat","tan"],["ate","eat","tea"]]

**Example 2**

- **Input:** strs = [""]
- **Output:** [[""]]

---

## 2. Brute Force Solution

For each string, check against every other string (or every group) whether they are anagrams (e.g. by sorting both or by count array). Build groups. O(n² * m) or similar. Correct but slow.

### Python (optimal — group by key)

The standard approach is to use a **signature** (key) for each string so that anagrams get the same key, then group by that key. Signature = sorted(s) as tuple, or tuple of character counts. This is already optimal.

---

## 3. Optimized Solution

**Idea:** Map each string to a key that is the same for all anagrams. Two options: (1) key = tuple(sorted(s)), (2) key = tuple of 26 counts. Then group: dict[key].append(s).

**Steps:**

1. groups = {}.  
2. For each s in strs: key = tuple(sorted(s)) (or count signature); groups.setdefault(key, []).append(s).  
3. Return list(groups.values()).

### Python (sorted key)

```python
def group_anagrams(strs: list[str]) -> list[list[str]]:
    groups = {}
    for s in strs:
        key = tuple(sorted(s))
        groups.setdefault(key, []).append(s)
    return list(groups.values())
```

### Python (count key — O(n * m))

```python
def group_anagrams(strs: list[str]) -> list[list[str]]:
    groups = {}
    for s in strs:
        count = [0] * 26
        for c in s:
            count[ord(c) - ord('a')] += 1
        key = tuple(count)
        groups.setdefault(key, []).append(s)
    return list(groups.values())
```

### Golang (sorted key)

```go
func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, s := range strs {
		b := []byte(s)
		sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
		key := string(b)
		groups[key] = append(groups[key], s)
	}
	out := make([][]string, 0, len(groups))
	for _, v := range groups {
		out = append(out, v)
	}
	return out
}
```

- **Time:** O(n * m log m) with sort key; O(n * m) with count key (m = max string length).  
- **Space:** O(n * m) for output

---

## 4. Walkthrough

"eat" → key (a,e,t) → group ["eat"]; "tea" → same key → ["eat","tea"]; "tan" → (a,n,t); "ate" → (a,e,t) → ["eat","tea","ate"]; etc. Result: [["bat"],["nat","tan"],["ate","eat","tea"]].

---

## 5. Edge Cases

1. **Empty string:** key = () or (0)*26; group [""].  
2. **Single string:** one group.  
3. **All same:** one group.  
4. **No anagrams:** each in its own group.

---

## 6. Pitfalls

1. **Mutable key:** In Python, use tuple(sorted(s)), not list.  
2. **Go map key:** String from sorted bytes is a valid key.  
3. **Order:** Problem says any order; order within group often input order.

---

## 7. Follow-ups

1. **Return indices?** Store indices in groups.  
2. **Count groups?** len(groups).  
3. **Largest group?** max(len(v) for v in groups.values()).

---

## 8. When to Use

- **Group by “same multiset”:** Key = sorted string or count tuple.  
- **Pattern:** Hash by signature; same in “valid anagram”, “group shifted strings”.
