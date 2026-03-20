# Product of Array Except Self — LeetCode 238 (M)

## 1. What is the problem?

**LeetCode 238 — Product of Array Except Self**

Given an integer array `nums`, return an array `answer` such that `answer[i]` equals the product of all elements of `nums` except `nums[i]`. You must write an O(n) solution without using division.

**Example 1**

- **Input:** `nums = [1, 2, 3, 4]`
- **Output:** `[24, 12, 8, 6]`

**Example 2**

- **Input:** `nums = [-1, 1, 0, -3, 3]`
- **Output:** `[0, 0, 9, 0, 0]`

---

## 2. Brute Force Solution

For each index `i`, compute the product of all elements except `nums[i]` with a loop (or multiply all then divide by `nums[i]` if division were allowed). Without division: nested loop for each `i`.

### Python (brute force, no division)

```python
def product_except_self_brute(nums: list[int]) -> list[int]:
    n = len(nums)
    out = [1] * n
    for i in range(n):
        for j in range(n):
            if i != j:
                out[i] *= nums[j]
    return out
```

### Golang (brute force)

```go
func productExceptSelfBrute(nums []int) []int {
	n := len(nums)
	out := make([]int, n)
	for i := 0; i < n; i++ {
		out[i] = 1
		for j := 0; j < n; j++ {
			if i != j {
				out[i] *= nums[j]
			}
		}
	}
	return out
}
```

- **Time:** O(n²)  
- **Space:** O(1) excluding output

---

## 3. Optimized Solution

**Idea:** `answer[i] = (product of nums[0..i-1]) * (product of nums[i+1..n-1])`. So use two passes: left products (prefix) and right products (suffix). Build prefix in one array, then multiply by suffix in place.

**Steps:**

1. Let `n = len(nums)`. Create `ans = [1]*n`.
2. **Left pass:** For `i` from 0 to n−1, set `ans[i] = left_product`, then `left_product *= nums[i]`.
3. **Right pass:** For `i` from n−1 down to 0, set `ans[i] *= right_product`, then `right_product *= nums[i]`.
4. Return `ans`.

### Python (optimized)

```python
def product_except_self(nums: list[int]) -> list[int]:
    n = len(nums)
    ans = [1] * n
    left = 1
    for i in range(n):
        ans[i] = left
        left *= nums[i]
    right = 1
    for i in range(n - 1, -1, -1):
        ans[i] *= right
        right *= nums[i]
    return ans
```

### Golang (optimized)

```go
func productExceptSelf(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	left := 1
	for i := 0; i < n; i++ {
		ans[i] = left
		left *= nums[i]
	}
	right := 1
	for i := n - 1; i >= 0; i-- {
		ans[i] *= right
		right *= nums[i]
	}
	return ans
}
```

- **Time:** O(n)  
- **Space:** O(1) excluding output (output doesn’t count per problem)

---

## 4. Walkthrough

**Example 1:** `nums = [1, 2, 3, 4]`

- Left pass: `ans = [1, 1, 2, 6]`, left becomes 24.
- Right pass: right starts 1; ans[3]*=1 → 6, right=4; ans[2]*=4 → 8, right=12; ans[1]*=12 → 12, right=24; ans[0]*=24 → 24.
- Result: `[24, 12, 8, 6]`.

---

## 5. Edge Cases

1. **Two elements:** `[a, b]` → `[b, a]`.  
2. **One zero:** product except that index is nonzero; others 0.  
3. **Two zeros:** all 0.  
4. **Negative:** product sign changes; logic unchanged.  
5. **Single element:** return `[1]` (product of empty set).

---

## 6. Pitfalls

1. **Using division:** Problem disallows division (e.g. when there is a 0).  
2. **Extra array:** We can do one output array + two running variables for O(1) extra space.  
3. **Order of updates:** In left pass, assign then multiply so current element is excluded.

---

## 7. Follow-ups

1. **Can you do it in one pass?** Yes: maintain left and right product in one loop with two pointers (conceptually same, one pass with both updates).  
2. **With division allowed?** Compute total product, then `ans[i] = total / nums[i]` (handle zero).  
3. **Product of last k elements?** Sliding window product or prefix product.

---

## 8. When to Use

- **“Except self” / prefix and suffix:** Two passes (prefix from left, suffix from right) or one array + two running values.  
- **Pattern:** Prefix/suffix product or sum; same idea in “trapping rain water” style problems.
