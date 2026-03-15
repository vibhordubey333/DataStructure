# Contains Duplicate — LeetCode 217 (E)

## 1. What is the problem?

**LeetCode 217 — Contains Duplicate**

Given an integer array `nums`, return `true` if any value appears at least twice, otherwise `false`.

**Example 1**

- **Input:** `nums = [1, 2, 3, 1]`
- **Output:** `true`

**Example 2**

- **Input:** `nums = [1, 2, 3, 4]`
- **Output:** `false`

---

## 2. Brute Force Solution

For each index `i`, check if `nums[i]` appears at any index `j > i`. Return true as soon as a duplicate is found.

### Python (brute force)

```python
def contains_duplicate_brute(nums: list[int]) -> bool:
    n = len(nums)
    for i in range(n):
        for j in range(i + 1, n):
            if nums[i] == nums[j]:
                return True
    return False
```

### Golang (brute force)

```go
func containsDuplicateBrute(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}
```

- **Time:** O(n²)  
- **Space:** O(1)

---

## 3. Optimized Solution

**Idea:** Use a set (hash set). For each element, if it’s already in the set, return true; otherwise add it. One pass.

**Steps:**

1. Create an empty set.
2. For each `x` in `nums`: if `x` is in the set, return true; else add `x`.
3. Return false.

### Python (optimized)

```python
def contains_duplicate(nums: list[int]) -> bool:
    seen = set()
    for x in nums:
        if x in seen:
            return True
        seen.add(x)
    return False
```

### Golang (optimized)

```go
func containsDuplicate(nums []int) bool {
	seen := make(map[int]struct{})
	for _, x := range nums {
		if _, ok := seen[x]; ok {
			return true
		}
		seen[x] = struct{}{}
	}
	return false
}
```

- **Time:** O(n)  
- **Space:** O(n)

---

## 4. Walkthrough

**Example 1:** `nums = [1, 2, 3, 1]`

| Step | x | seen      | in seen? |
|------|---|-----------|----------|
| 1    | 1 | {}        | no → add |
| 2    | 2 | {1}       | no → add |
| 3    | 3 | {1,2}     | no → add |
| 4    | 1 | {1,2,3}   | yes → return true |

---

## 5. Edge Cases

1. **Empty array:** `[]` → false.  
2. **Single element:** `[1]` → false.  
3. **Two same:** `[1, 1]` → true.  
4. **All same:** `[2,2,2,2]` → true.  
5. **Large n:** Hash set avoids O(n²).

---

## 6. Pitfalls

1. **Sorting:** Sort then adjacent check is O(n log n); set is O(n) and often preferred.  
2. **Using map for count:** Only need “seen or not”; set is enough.  
3. **Off-by-one:** No index math needed with set.

---

## 7. Follow-ups

1. **Contains duplicate within distance k?** Sliding window + set of size k.  
2. **Within value difference t?** Bucket by `x // (t+1)` or sorted order + sliding window.  
3. **Return all duplicates?** Use map: count frequency, return keys with count &gt; 1.

---

## 8. When to Use

- **Uniqueness / “seen before”?** → Set.  
- **Frequency?** → Map (value → count).  
- **Sorted + adjacent?** When O(n log n) is acceptable and you need order.
