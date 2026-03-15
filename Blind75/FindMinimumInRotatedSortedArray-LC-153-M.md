# Find Minimum in Rotated Sorted Array — LeetCode 153 (M)

## 1. What is the problem?

**LeetCode 153 — Find Minimum in Rotated Sorted Array**

Suppose an array of length `n` sorted in ascending order is rotated between 1 and n times. Given the sorted rotated array `nums` (unique elements), return the minimum element. O(log n) required.

**Example 1**

- **Input:** `nums = [3, 4, 5, 1, 2]`
- **Output:** `1`

**Example 2**

- **Input:** `nums = [4, 5, 6, 7, 0, 1, 2]`
- **Output:** `0`

---

## 2. Brute Force Solution

Scan the array and return the minimum. Correct because the minimum exists somewhere in the array.

### Python (brute force)

```python
def find_min_brute(nums: list[int]) -> int:
    return min(nums)
```

### Golang (brute force)

```go
func findMinBrute(nums []int) int {
	min := nums[0]
	for _, x := range nums[1:] {
		if x < min {
			min = x
		}
	}
	return min
}
```

- **Time:** O(n)  
- **Space:** O(1)

---

## 3. Optimized Solution (Binary Search)

**Idea:** The array has two sorted segments: [large…] then [small…]. The minimum is the first element of the second segment. Compare `mid` with `right`: if `nums[mid] > nums[right]`, the pivot is in the right half (mid+1..right); else the minimum is in the left half including mid (left..mid).

**Steps:**

1. Let `left = 0`, `right = len(nums)-1`.
2. While `left < right`:  
   - `mid = left + (right - left) // 2`  
   - If `nums[mid] > nums[right]`: `left = mid + 1`  
   - Else: `right = mid`  
3. Return `nums[left]`.

### Python (optimized)

```python
def find_min(nums: list[int]) -> int:
    left, right = 0, len(nums) - 1
    while left < right:
        mid = (left + right) // 2
        if nums[mid] > nums[right]:
            left = mid + 1
        else:
            right = mid
    return nums[left]
```

### Golang (optimized)

```go
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}
```

- **Time:** O(log n)  
- **Space:** O(1)

---

## 4. Walkthrough

**Example 2:** `nums = [4, 5, 6, 7, 0, 1, 2]`

| left | right | mid | nums[mid] vs nums[right] | action        |
|------|-------|-----|---------------------------|---------------|
| 0    | 6     | 3   | 7 > 2 → min in right      | left = 4      |
| 4    | 6     | 5   | 1 < 2 → min at or left of 5 | right = 5   |
| 4    | 5     | 4   | 0 < 2 → right = 4         | right = 4     |
| 4    | 4     | -   | stop                       | return nums[4]=0 |

---

## 5. Edge Cases

1. **Not rotated:** `[1,2,3,4]` → 1; `nums[mid] <= nums[right]` keeps moving right to mid.  
2. **Single element:** return that element.  
3. **Two elements:** `[2,1]` → 1; mid=0, nums[0]>nums[1] → left=1.  
4. **Rotated n-1:** `[2,3,4,1]`; still works.  
5. **All same:** problem says unique; no duplicates.

---

## 6. Pitfalls

1. **Comparing with nums[0] or nums[left]:** Comparing with `nums[right]` correctly separates “still in first segment” vs “in second segment”.  
2. **Using left <= right:** Use `left < right` so we stop when one element; that element is the min.  
3. **right = mid - 1:** Would skip the minimum when it’s at mid; use `right = mid`.

---

## 7. Follow-ups

1. **Search in rotated sorted array (LC 33)?** Same idea: determine which half is sorted and whether target is in that range.  
2. **Duplicates allowed (LC 154)?** When `nums[mid] == nums[right]`, do `right--` to avoid O(n) in all-equal case.  
3. **Return index of min?** Return `left` instead of `nums[left]`.

---

## 8. When to Use

- **Sorted rotated array / find pivot/min:** Binary search by comparing mid with the right end.  
- **Pattern:** “Rotated sorted” → one half is always sorted; use that to discard half.
