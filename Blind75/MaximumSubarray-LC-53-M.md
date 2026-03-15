# Maximum Subarray — LeetCode 53 (M) — Kadane's Algorithm

## 1. What is the problem?

**LeetCode 53 — Maximum Subarray**

Given an integer array `nums`, find the contiguous subarray (containing at least one number) that has the largest sum and return its sum.

**Example 1**

- **Input:** `nums = [-2, 1, -3, 4, -1, 2, 1, -5, 4]`
- **Output:** `6`
- **Explanation:** Subarray `[4, -1, 2, 1]` has sum 6.

**Example 2**

- **Input:** `nums = [1]`
- **Output:** `1`

---

## 2. Brute Force Solution

Try every subarray `[i..j]`, compute sum, and keep the maximum. Use prefix sum to get range sum in O(1) per subarray.

### Python (brute force)

```python
def max_subarray_brute(nums: list[int]) -> int:
    n = len(nums)
    best = nums[0]
    for i in range(n):
        s = 0
        for j in range(i, n):
            s += nums[j]
            best = max(best, s)
    return best
```

### Golang (brute force)

```go
func maxSubArrayBrute(nums []int) int {
	best := nums[0]
	for i := 0; i < len(nums); i++ {
		s := 0
		for j := i; j < len(nums); j++ {
			s += nums[j]
			if s > best {
				best = s
			}
		}
	}
	return best
}
```

- **Time:** O(n²)  
- **Space:** O(1)

---

## 3. Optimized Solution (Kadane's Algorithm)

**Idea:** For each position, the best sum ending **here** is either (best sum ending at previous + this element) or just this element. So: `cur = max(nums[i], cur + nums[i])`, then `best = max(best, cur)`.

**Steps:**

1. Initialize `cur = nums[0]`, `best = nums[0]`.
2. For `i` from 1 to n−1: `cur = max(nums[i], cur + nums[i])`, `best = max(best, cur)`.
3. Return `best`.

### Python (optimized)

```python
def max_subarray(nums: list[int]) -> int:
    cur = best = nums[0]
    for x in nums[1:]:
        cur = max(x, cur + x)
        best = max(best, cur)
    return best
```

### Golang (optimized)

```go
func maxSubArray(nums []int) int {
	cur, best := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		x := nums[i]
		if x > cur+x {
			cur = x
		} else {
			cur += x
		}
		if cur > best {
			best = cur
		}
	}
	return best
}
```

- **Time:** O(n)  
- **Space:** O(1)

---

## 4. Walkthrough

**Example 1:** `nums = [-2, 1, -3, 4, -1, 2, 1, -5, 4]`

| i | x  | cur (max(x, cur+x)) | best |
|---|----|----------------------|------|
| 0 | -2 | -2                   | -2   |
| 1 | 1  | 1                    | 1    |
| 2 | -3 | -2                   | 1    |
| 3 | 4  | 4                    | 4    |
| 4 | -1 | 3                    | 4    |
| 5 | 2  | 5                    | 5    |
| 6 | 1  | 6                    | 6    |
| 7 | -5 | 1                    | 6    |
| 8 | 4  | 5                    | 6    |

Output: `6`.

---

## 5. Edge Cases

1. **All negative:** e.g. `[-1,-2]` → -1 (single element).  
2. **Single element:** return that element.  
3. **All positive:** whole array.  
4. **One positive in negatives:** that single positive.  
5. **Zero:** can be part of optimal subarray.

---

## 6. Pitfalls

1. **Initializing best to 0:** Fails when all numbers are negative; use `nums[0]`.  
2. **Resetting cur only when negative:** Logic is “extend or start fresh” every step.  
3. **Returning cur instead of best:** We need the maximum over all ending positions.

---

## 7. Follow-ups

1. **Return the subarray indices?** Track start/end when you update `best`.  
2. **Maximum product subarray?** Track both max and min (LeetCode 152).  
3. **Circular array?** Max subarray either doesn’t wrap (Kadane) or wraps (total − min subarray).  
4. **At most k elements?** Sliding window or DP.

---

## 8. When to Use

- **Maximum contiguous sum:** Kadane (one pass, O(n)).  
- **Maximum product subarray:** Same idea but keep max and min (negative × negative = positive).  
- **Pattern:** DP where state is “best sum ending at current index.”
