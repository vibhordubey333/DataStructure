# Combination Sum IV — LeetCode 377 (M)

## 1. What is the problem?

**LeetCode 377 — Combination Sum IV**

Given an array of distinct integers `nums` and a target `target`, return the number of possible combinations that add up to `target`. Same element may be used multiple times. Order matters (1+2 and 2+1 count as different).

**Example 1**

- **Input:** nums = [1, 2, 3], target = 4
- **Output:** 7

**Example 2**

- **Input:** nums = [9], target = 3
- **Output:** 0

---

## 2. Brute Force Solution

Recurse: for each coin c, if target>=c, add count(1 + f(target-c)). Base: target==0 → 1; target<0 → 0. Exponential without memo.

### Python (brute force)

```python
def combination_sum4_brute(nums: list[int], target: int) -> int:
    def f(t):
        if t == 0:
            return 1
        if t < 0:
            return 0
        total = 0
        for c in nums:
            total += f(t - c)
        return total
    return f(target)
```

### Golang (brute force)

```go
func combinationSum4Brute(nums []int, target int) int {
	var f func(int) int
	f = func(t int) int {
		if t == 0 {
			return 1
		}
		if t < 0 {
			return 0
		}
		total := 0
		for _, c := range nums {
			total += f(t - c)
		}
		return total
	}
	return f(target)
}
```

- **Time:** O(|nums|^target)  
- **Space:** O(target) stack

---

## 3. Optimized Solution (DP)

**Idea:** dp[t] = number of combinations that sum to t. dp[0]=1. dp[t] = sum over c in nums of dp[t-c] (if t>=c).

### Python (optimized)

```python
def combination_sum4(nums: list[int], target: int) -> int:
    dp = [0] * (target + 1)
    dp[0] = 1
    for t in range(1, target + 1):
        for c in nums:
            if c <= t:
                dp[t] += dp[t - c]
    return dp[target]
```

### Golang (optimized)

```go
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for t := 1; t <= target; t++ {
		for _, c := range nums {
			if c <= t {
				dp[t] += dp[t-c]
			}
		}
	}
	return dp[target]
}
```

- **Time:** O(target * len(nums))  
- **Space:** O(target)

---

## 4. Walkthrough

nums=[1,2,3], target=4. dp[0]=1. dp[1]=dp[0]=1. dp[2]=dp[1]+dp[0]=2. dp[3]=dp[2]+dp[1]+dp[0]=4. dp[4]=dp[3]+dp[2]+dp[1]=4+2+1=7. Return 7.

---

## 5. Edge Cases

1. **target=0:** 1 (empty combination).  
2. **No way:** e.g. nums=[9], target=3 → 0.  
3. **Overflow:** Problem says answer fits in 32-bit; use mod if needed.  
4. **Single coin:** dp[t]=1 if t divisible by coin else 0.

---

## 6. Pitfalls

1. **Order matters:** This is permutations, not combinations. So we iterate t first, then coins (all ways to reach t).  
2. **Combination Sum I (order doesn't matter):** Different: use "start index" to avoid duplicates.  
3. **Negative numbers:** Problem says distinct positive; no negatives.

---

## 7. Follow-ups

1. **Return actual combinations?** Backtrack and store paths.  
2. **Combination Sum I (unique combinations)?** Sort, use start index, no reuse of same order.  
3. **Mod 10^9+7?** Add mod to avoid overflow.

---

## 8. When to Use

- **Count ways to form target (order matters):** DP over target; dp[t] += dp[t-c].  
- **Pattern:** Unbounded knapsack count; different from "min coins" or "unique combinations".
