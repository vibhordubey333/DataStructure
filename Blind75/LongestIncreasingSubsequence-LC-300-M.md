# Longest Increasing Subsequence — LeetCode 300 (M)

## 1. What is the problem?

**LeetCode 300 — Longest Increasing Subsequence**

Given an integer array `nums`, return the length of the longest strictly increasing subsequence (not necessarily contiguous).

**Example 1**

- **Input:** nums = [10, 9, 2, 5, 3, 7, 101, 18]
- **Output:** 4 (e.g. [2, 3, 7, 101])

**Example 2**

- **Input:** nums = [0, 1, 0, 3, 2, 3]
- **Output:** 4

---

## 2. Brute Force Solution

For each element, either include or exclude and recurse (with “last taken” value to enforce increasing). O(2^n). Or enumerate all subsequences. Correct but slow.

### Python (brute force — recursive)

```python
def length_of_lis_brute(nums: list[int]) -> int:
    def f(i: int, last: int) -> int:
        if i == len(nums):
            return 0
        skip = f(i + 1, last)
        take = 0
        if nums[i] > last:
            take = 1 + f(i + 1, nums[i])
        return max(skip, take)
    return f(0, -10**9 - 1)
```

### Golang (brute force)

```go
func lengthOfLISBrute(nums []int) int {
	var f func(int, int) int
	f = func(i, last int) int {
		if i == len(nums) {
			return 0
		}
		skip := f(i+1, last)
		take := 0
		if nums[i] > last {
			take = 1 + f(i+1, nums[i])
		}
		if skip > take {
			return skip
		}
		return take
	}
	return f(0, -1e9-1)
}
```

- **Time:** O(2^n)  
- **Space:** O(n) stack

---

## 3. Optimized Solution

**Idea (O(n²) DP):** dp[i] = length of LIS ending at index i. dp[i] = 1 + max(dp[j]) over all j < i with nums[j] < nums[i]. Answer = max(dp).

**Idea (O(n log n)):** Maintain a “tails” array: tails[len] = smallest ending value of an increasing subsequence of length len. For each x, binary search: replace first element >= x (or append). Length of tails is the answer.

**Steps (O(n log n)):**

1. tails = [].  
2. For each x in nums: find leftmost index where tails[i] >= x (bisect_left). If i == len(tails): append x; else tails[i] = x.  
3. Return len(tails).

### Python (O(n²) DP)

```python
def length_of_lis(nums: list[int]) -> int:
    n = len(nums)
    dp = [1] * n
    for i in range(n):
        for j in range(i):
            if nums[j] < nums[i]:
                dp[i] = max(dp[i], dp[j] + 1)
    return max(dp) if dp else 0
```

### Python (O(n log n) — greedy + binary search)

```python
import bisect
def length_of_lis_nlogn(nums: list[int]) -> int:
    tails = []
    for x in nums:
        i = bisect.bisect_left(tails, x)
        if i == len(tails):
            tails.append(x)
        else:
            tails[i] = x
    return len(tails)
```

### Golang (O(n²))

```go
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
	}
	max := 1
	for _, v := range dp {
		if v > max {
			max = v
		}
	}
	return max
}
```

### Golang (O(n log n))

```go
func lengthOfLIS(nums []int) int {
	tails := []int{}
	for _, x := range nums {
		i := sort.SearchInts(tails, x)
		if i == len(tails) {
			tails = append(tails, x)
		} else {
			tails[i] = x
		}
	}
	return len(tails)
}
```

- **Time:** O(n²) or O(n log n)  
- **Space:** O(n)

---

## 4. Walkthrough

nums = [10, 9, 2, 5, 3, 7, 101, 18]. O(n log n): tails: 10 → 9 → 2 → 2,5 → 2,3 → 2,3,7 → 2,3,7,101 → 2,3,7,18. Length 4.

---

## 5. Edge Cases

1. **Empty array:** 0.  
2. **Single element:** 1.  
3. **Strictly decreasing:** 1.  
4. **Strictly increasing:** n.  
5. **All same:** 1 (strictly increasing means strictly greater).

---

## 6. Pitfalls

1. **Non-strict (<=):** Problem says “strictly increasing”; use <.  
2. **Subarray vs subsequence:** Subsequence can skip elements.  
3. **O(n²) vs O(n log n):** Mention binary search for full credit.

---

## 7. Follow-ups

1. **Print one LIS?** Store parent in DP; backtrack.  
2. **Number of LIS?** DP with count.  
3. **Longest non-decreasing?** Change < to <= and tails logic accordingly.

---

## 8. When to Use

- **LIS:** O(n²) DP or O(n log n) with “tails” + binary search.  
- **Pattern:** Classic DP; “tails” array is a common optimization.
