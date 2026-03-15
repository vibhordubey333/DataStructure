# Maximum Product Subarray — LeetCode 152 (M)

## 1. What is the problem?

**LeetCode 152 — Maximum Product Subarray**

Given an integer array `nums`, find the contiguous subarray that has the largest product and return the product.

**Example 1**

- **Input:** `nums = [2, 3, -2, 4]`
- **Output:** `6`
- **Explanation:** Subarray `[2, 3]` has product 6.

**Example 2**

- **Input:** `nums = [-2, 0, -1]`
- **Output:** `0`

---

## 2. Brute Force Solution

Try every subarray `[i..j]`, compute product, track maximum. O(n²) or O(n³) if computing product from scratch each time (or O(n²) with running product).

### Python (brute force)

```python
def max_product_brute(nums: list[int]) -> int:
    n = len(nums)
    best = nums[0]
    for i in range(n):
        p = 1
        for j in range(i, n):
            p *= nums[j]
            best = max(best, p)
    return best
```

### Golang (brute force)

```go
func maxProductBrute(nums []int) int {
	best := nums[0]
	for i := 0; i < len(nums); i++ {
		p := 1
		for j := i; j < len(nums); j++ {
			p *= nums[j]
			if p > best {
				best = p
			}
		}
	}
	return best
}
```

- **Time:** O(n²)  
- **Space:** O(1)

---

## 3. Optimized Solution

**Idea:** Unlike sum, product can become maximum when we multiply two negatives. So track both **max product** and **min product** ending at current position. For each `x`: new_max = max(x, max*x, min*x), new_min = min(x, max*x, min*x).

**Steps:**

1. Set `cur_max = cur_min = best = nums[0]`.
2. For each `x` in `nums[1:]`:  
   - If `x < 0`, swapping max and min helps: so compute `cur_max = max(x, cur_max*x, cur_min*x)` and `cur_min = min(x, cur_max_old*x, cur_min*x)` (use temp for old cur_max).
   - Or simply: `candidates = (x, cur_max*x, cur_min*x)`, then `cur_max = max(candidates)`, `cur_min = min(candidates)`, `best = max(best, cur_max)`.
3. Return `best`.

### Python (optimized)

```python
def max_product(nums: list[int]) -> int:
    cur_max = cur_min = best = nums[0]
    for x in nums[1:]:
        a, b = cur_max * x, cur_min * x
        cur_max = max(x, a, b)
        cur_min = min(x, a, b)
        best = max(best, cur_max)
    return best
```

### Golang (optimized)

```go
func maxProduct(nums []int) int {
	curMax, curMin, best := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		x := nums[i]
		a, b := curMax*x, curMin*x
		if x > a {
			if x > b {
				curMax = x
			} else {
				curMax = b
			}
		} else {
			if a > b {
				curMax = a
			} else {
				curMax = b
			}
		}
		if x < a {
			if x < b {
				curMin = x
			} else {
				curMin = b
			}
		} else {
			if a < b {
				curMin = a
			} else {
				curMin = b
			}
		}
		if curMax > best {
			best = curMax
		}
	}
	return best
}
```

Cleaner Go using a helper:

```go
func maxProduct(nums []int) int {
	curMax, curMin, best := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		x := nums[i]
		a, b := curMax*x, curMin*x
		curMax = max3(x, a, b)
		curMin = min3(x, a, b)
		if curMax > best {
			best = curMax
		}
	}
	return best
}
func max3(a, b, c int) int {
	if a >= b && a >= c {
		return a
	}
	if b >= c {
		return b
	}
	return c
}
func min3(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= c {
		return b
	}
	return c
}
```

- **Time:** O(n)  
- **Space:** O(1)

---

## 4. Walkthrough

**Example 1:** `nums = [2, 3, -2, 4]`

| i | x  | a=cur_max*x | b=cur_min*x | cur_max | cur_min | best |
|---|----|-------------|-------------|---------|---------|------|
| 0 | 2  | -           | -           | 2       | 2       | 2    |
| 1 | 3  | 6           | 6           | 6       | 6       | 6    |
| 2 | -2 | -12         | -12         | -2      | -12     | 6    |
| 3 | 4  | -8          | -48         | 4       | -48     | 6    |

Output: `6`.

---

## 5. Edge Cases

1. **Single element:** return that element.  
2. **One zero:** product can be 0; min/max both can become 0, then next step uses 0.  
3. **Two negatives:** e.g. `[-2,-3]` → 6.  
4. **Zero in middle:** resets; track max so 0 is considered.  
5. **All negative odd length:** product is negative; best is least negative (or single element).

---

## 6. Pitfalls

1. **Only tracking max:** Min product can become max after multiplying by negative.  
2. **Forgetting to consider “start at x”:** Always include `x` in max/min (so one negative resets max to that negative until next neg).  
3. **Using only positive products:** Fails for arrays with negatives.

---

## 7. Follow-ups

1. **Return the subarray?** Store start/end when updating best.  
2. **Maximum sum subarray?** Kadane (no need for min).  
3. **At most k negatives?** More complex DP.

---

## 8. When to Use

- **Contiguous product:** Track both max and min ending at current index (DP / one pass).  
- **Pattern:** “Product” + “subarray” → consider sign; max and min swap when multiplying by negative.
