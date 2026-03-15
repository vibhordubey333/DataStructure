# House Robber — LeetCode 198 (M)

## 1. What is the problem?

**LeetCode 198 — House Robber**

You are given an array `nums` where nums[i] is the amount of money at the i-th house. You cannot rob two adjacent houses. Return the maximum amount you can rob.

**Example 1**

- **Input:** nums = [1, 2, 3, 1]
- **Output:** 4 (rob house 0 and 2)

**Example 2**

- **Input:** nums = [2, 7, 9, 3, 1]
- **Output:** 12 (rob 2, 9, 1)

---

## 2. Brute Force Solution

Recurse: at each index i, either rob i (nums[i] + f(i+2)) or skip (f(i+1)); return max. Exponential without memo.

### Python (brute)

```python
def rob_brute(nums: list[int]) -> int:
    def f(i):
        if i >= len(nums):
            return 0
        return max(nums[i] + f(i + 2), f(i + 1))
    return f(0)
```

### Golang (brute)

```go
func robBrute(nums []int) int {
	var f func(int) int
	f = func(i int) int {
		if i >= len(nums) {
			return 0
		}
		a := nums[i] + f(i+2)
		b := f(i + 1)
		if a > b {
			return a
		}
		return b
	}
	return f(0)
}
```

- **Time:** O(2^n)  
- **Space:** O(n) stack

---

## 3. Optimized Solution (DP)

**Idea:** dp[i] = max money from houses 0..i (we may or may not rob i). dp[i] = max(dp[i-1], nums[i] + dp[i-2]). Can use two variables: prev2, prev1 for i-2 and i-1.

**Steps:**

1. prev2 = 0, prev1 = 0.  
2. For each x in nums: cur = max(prev1, x + prev2); prev2, prev1 = prev1, cur.  
3. Return prev1.

### Python (optimized)

```python
def rob(nums: list[int]) -> int:
    prev2 = prev1 = 0
    for x in nums:
        cur = max(prev1, x + prev2)
        prev2, prev1 = prev1, cur
    return prev1
```

### Golang (optimized)

```go
func rob(nums []int) int {
	prev2, prev1 := 0, 0
	for _, x := range nums {
		cur := prev1
		if x+prev2 > cur {
			cur = x + prev2
		}
		prev2, prev1 = prev1, cur
	}
	return prev1
}
```

- **Time:** O(n)  
- **Space:** O(1)

---

## 4. Walkthrough

nums = [2, 7, 9, 3, 1]. prev2=0, prev1=0 → cur=2; prev2=0, prev1=2 → cur=7; prev2=2, prev1=7 → cur=11 (9+2); prev2=7, prev1=11 → cur=11 (3+7 vs 11); prev2=11, prev1=11 → cur=12 (1+11). Return 12.

---

## 5. Edge Cases

1. **Empty:** 0.  
2. **One house:** nums[0].  
3. **Two houses:** max(nums[0], nums[1]).  
4. **All same:** depends on parity.

---

## 6. Pitfalls

1. **Base case:** prev2=0, prev1=0 so first iteration gives nums[0].  
2. **Adjacent constraint:** Only use prev2 when robbing current.  
3. **House Robber II (circular):** Run DP on [0..n-2] and [1..n-1]; take max.

---

## 7. Follow-ups

1. **House Robber II (LC 213)?** Circular: max(rob(nums[1:]), rob(nums[:-1])).  
2. **Tree (LC 337)?** DP on tree: (rob root, skip root).  
3. **Print chosen houses?** Store parent in DP.

---

## 8. When to Use

- **Non-adjacent selection for max sum:** Linear DP with two previous states.  
- **Pattern:** Classic 1D DP; same idea in “max sum no two consecutive”.
