# Two Sum — LeetCode 1 (E)

## 1. What is the problem?

**LeetCode 1 — Two Sum**

Given an array of integers `nums` and an integer `target`, return indices of the two numbers that add up to `target`. You may assume exactly one solution exists, and you may not use the same element twice.

**Example 1**

- **Input:** `nums = [2, 7, 11, 15]`, `target = 9`
- **Output:** `[0, 1]`
- **Explanation:** `nums[0] + nums[1] = 2 + 7 = 9`.

**Example 2**

- **Input:** `nums = [3, 2, 4]`, `target = 6`
- **Output:** `[1, 2]`
- **Explanation:** `nums[1] + nums[2] = 2 + 4 = 6`.

---

## 2. Brute Force Solution

Check every pair of indices `(i, j)` with `i < j`. If `nums[i] + nums[j] == target`, return `[i, j]`. This is correct because we exhaust all pairs.

### Python (brute force)

```python
def two_sum_brute(nums: list[int], target: int) -> list[int]:
    n = len(nums)
    for i in range(n):
        for j in range(i + 1, n):
            if nums[i] + nums[j] == target:
                return [i, j]
    return []  # unreachable if solution exists
```

### Golang (brute force)

```go
package main

func twoSumBrute(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
```

- **Time:** O(n²) — two nested loops over `n` elements.
- **Space:** O(1) — only a few variables.

---

## 3. Optimized Solution

**Idea:** For each value `nums[i]`, we need a partner `target - nums[i]`. Use a hash map: key = value seen, value = index. In one pass, for each `x = nums[i]`, if `target - x` is already in the map, return `[map[target-x], i]`; otherwise store `map[x] = i`.

**Steps:**

1. Create an empty map: value → index.
2. For each index `i` and value `x = nums[i]`:
   - If `target - x` exists in the map, return `[map[target-x], i]`.
   - Else set `map[x] = i`.
3. (If guaranteed a solution, no fallback needed.)

### Python (optimized)

```python
def two_sum(nums: list[int], target: int) -> list[int]:
    seen = {}
    for i, x in enumerate(nums):
        need = target - x
        if need in seen:
            return [seen[need], i]
        seen[x] = i
    return []
```

### Golang (optimized)

```go
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, x := range nums {
		need := target - x
		if j, ok := seen[need]; ok {
			return []int{j, i}
		}
		seen[x] = i
	}
	return nil
}
```

- **Time:** O(n) — single pass.
- **Space:** O(n) — map of up to `n` entries.

---

## 4. Walkthrough with Example

**Example 1:** `nums = [2, 7, 11, 15]`, `target = 9`

| Step | i | x  | need = 9−x | seen before     | action              |
|------|---|----|------------|-----------------|---------------------|
| 1    | 0 | 2  | 7          | {}              | 7 not in map → seen[2]=0 |
| 2    | 1 | 7  | 2          | {2→0}           | 2 in map → return [0,1]   |

Output: `[0, 1]`.

---

## 5. Edge Cases to Remember

1. **Two elements:** `[1, 2]`, `target = 3` → `[0, 1]`.
2. **Same value, different indices:** `[3, 3]`, `target = 6` → `[0, 1]` (store index for first 3, second 3 finds it).
3. **Zero and negative:** `[-1, 0, 1]`, `target = 0` → `[0, 2]`.
4. **Target is sum of first and last:** long array, answer at indices 0 and n−1.
5. **Large numbers:** ensure language doesn’t overflow (Python fine; Go int is fine for LeetCode).
6. **Exactly one solution:** no need to handle “no answer” if problem states it exists.

---

## 6. Pitfalls to Avoid

1. **Using the same index twice:** Check partner with `j > i` in brute force; in hash map, we only use indices already stored (from earlier in the array), so we never use the same index twice.
2. **Returning value instead of index:** Problem asks for indices; store index in the map.
3. **One-pass order:** Store current value after checking for partner so we don’t use same element (e.g. `target = 2*nums[i]`).
4. **Map key:** Use the **value** as key (to look up “need”), not the index.
5. **Brute force only:** Mentioning hash map shows you know the standard O(n) approach.

---

## 7. Follow-up Interview Questions

1. **What if the array is sorted?**  
   Two pointers at both ends: if sum &gt; target, decrement right; if sum &lt; target, increment left. O(n) time, O(1) space.

2. **Return all pairs that sum to target (indices or values)?**  
   Hash map: for each value, count how many times `target - x` appeared; handle duplicates (e.g. count frequency, then iterate and count valid pairs).

3. **Three sum / k sum?**  
   Sort and use two pointers for 3Sum; for k-sum, reduce to (k−1)-sum recursively or use two pointers for the last two.

4. **What if there are multiple solutions?**  
   Return one pair (e.g. first found) or adapt to collect all pairs (watch for duplicate pairs when using hash map or two pointers).

5. **Return values instead of indices?**  
   Same logic; return `[nums[i], nums[j]]` or store values in the map if you don’t need indices.

---

## 8. When to Use This Technique?

- **Hash map for “complement”:** Any “find two elements that satisfy a relation” (sum, difference, etc.) in an unsorted array often reduces to “for each element, is there a partner?” → map from value to index or count.
- **Pattern:** “Two Sum” pattern: one pass + hash map for O(n). Related: 3Sum (sort + two pointers), 4Sum, Two Sum II (sorted array → two pointers), Subarray Sum Equals K (prefix sum + map).
