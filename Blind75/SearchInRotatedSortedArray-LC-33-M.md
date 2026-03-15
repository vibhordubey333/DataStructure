# Search in Rotated Sorted Array — LeetCode 33 (M)

## 1. What is the problem?

**LeetCode 33 — Search in Rotated Sorted Array**

There is an integer array `nums` sorted in ascending order (distinct values), rotated at an unknown pivot. Given `target`, return its index if it is in `nums`, else -1. O(log n) required.

**Example 1**

- **Input:** `nums = [4, 5, 6, 7, 0, 1, 2]`, `target = 0`
- **Output:** `4`

**Example 2**

- **Input:** `nums = [4, 5, 6, 7, 0, 1, 2]`, `target = 3`
- **Output:** `-1`

---

## 2. Brute Force Solution

Linear scan: if `nums[i] == target` return `i`; else return -1.

### Python (brute force)

```python
def search_brute(nums: list[int], target: int) -> int:
    for i, x in enumerate(nums):
        if x == target:
            return i
    return -1
```

### Golang (brute force)

```go
func searchBrute(nums []int, target int) int {
	for i, x := range nums {
		if x == target {
			return i
		}
	}
	return -1
}
```

- **Time:** O(n)  
- **Space:** O(1)

---

## 3. Optimized Solution (Binary Search)

**Idea:** One half of the array is always sorted. Compare `target` with the sorted half: if target is in that range, search there; else search the other half.

**Steps:**

1. `left, right = 0, len(nums)-1`.
2. While `left <= right`:  
   - `mid = left + (right-left)/2`  
   - If `nums[mid] == target`: return mid  
   - If `nums[left] <= nums[mid]`: left half is sorted. If `nums[left] <= target < nums[mid]`, search left; else right.  
   - Else: right half is sorted. If `nums[mid] < target <= nums[right]`, search right; else left.  
3. Return -1.

### Python (optimized)

```python
def search(nums: list[int], target: int) -> int:
    left, right = 0, len(nums) - 1
    while left <= right:
        mid = (left + right) // 2
        if nums[mid] == target:
            return mid
        if nums[left] <= nums[mid]:
            if nums[left] <= target < nums[mid]:
                right = mid - 1
            else:
                left = mid + 1
        else:
            if nums[mid] < target <= nums[right]:
                left = mid + 1
            else:
                right = mid - 1
    return -1
```

### Golang (optimized)

```go
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
```

- **Time:** O(log n)  
- **Space:** O(1)

---

## 4. Walkthrough

**Example 1:** `nums = [4, 5, 6, 7, 0, 1, 2]`, `target = 0`

| left | right | mid | nums[mid] | left half sorted? | action           |
|------|-------|-----|-----------|---|------------------|
| 0    | 6     | 3   | 7         | yes (4≤7)        | 0 in [4,7)? No → left=4 |
| 4    | 6     | 5   | 1         | no               | 0 in (1,2]? No → right=4 |
| 4    | 4     | 4   | 0         | -                | match → return 4 |

---

## 5. Edge Cases

1. **Not rotated:** standard binary search.  
2. **Single element:** handle in loop.  
3. **Target at pivot:** e.g. min element; right half sorted.  
4. **Target not in array:** exit loop, return -1.  
5. **Two elements:** e.g. `[1,3]`, target 3; left half sorted, 3 not in [1,3) → left=1, then mid=1, match.

---

## 6. Pitfalls

1. **Using `nums[left] < nums[mid]`:** With two elements `[3,1]`, left=0, mid=0, so left half “sorted”; use `<=`.  
2. **Including mid in range:** Use `target < nums[mid]` for left half to avoid infinite loop.  
3. **Forgetting `left <= right`:** We need to check the last remaining element.

---

## 7. Follow-ups

1. **Duplicates allowed (LC 81)?** When `nums[left]==nums[mid]==nums[right]`, shrink (e.g. left++, right--) and may degrade to O(n).  
2. **Find pivot first then binary search?** Two passes; one modified binary search is simpler.  
3. **How many times rotated?** Pivot index = number of rotations (if 1..n rotations).

---

## 8. When to Use

- **Rotated sorted array + search:** Determine which half is sorted and whether target lies in that range.  
- **Pattern:** “Sorted but rotated” → one half always sorted; binary search with conditional branch.
