# House Robber II — LeetCode 213 (M)

## 1. What is the problem?

**LeetCode 213 — House Robber II**

Same as House Robber but houses are arranged in a **circle** (first and last are adjacent). Return max money without robbing two adjacent houses.

**Example 1**

- **Input:** nums = [2, 3, 2]
- **Output:** 3 (rob house 1 only)

**Example 2**

- **Input:** nums = [1, 2, 3, 1]
- **Output:** 4 (rob house 1 and 3)

---

## 2. Brute Force Solution

Same as House Robber but with constraint: if we rob house 0, we cannot rob house n-1. Try both cases: rob(0..n-2) and rob(1..n-1); take max. Each subproblem is standard House Robber.

### Python (brute — reuse House Robber)

```python
def rob_brute(nums: list[int]) -> int:
    def rob_linear(arr):
        prev2, prev1 = 0, 0
        for x in arr:
            cur = max(prev1, x + prev2)
            prev2, prev1 = prev1, cur
        return prev1
    if len(nums) <= 1:
        return nums[0] if nums else 0
    return max(rob_linear(nums[:-1]), rob_linear(nums[1:]))
```

### Golang (brute)

```go
func robBrute(nums []int) int {
	robLinear := func(arr []int) int {
		prev2, prev1 := 0, 0
		for _, x := range arr {
			cur := prev1
			if x+prev2 > cur {
				cur = x + prev2
			}
			prev2, prev1 = prev1, cur
		}
		return prev1
	}
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	a := robLinear(nums[:len(nums)-1])
	b := robLinear(nums[1:])
	if a > b {
		return a
	}
	return b
}
```

- **Time:** O(n)  
- **Space:** O(1)

---

## 3. Optimized Solution

**Idea:** Circular constraint → either we exclude first house (rob nums[1:]) or exclude last house (rob nums[:-1]). Run House Robber DP on both and return max.

### Python (optimized)

```python
def rob(nums: list[int]) -> int:
    def rob_linear(arr):
        prev2, prev1 = 0, 0
        for x in arr:
            prev2, prev1 = prev1, max(prev1, x + prev2)
        return prev1
    if len(nums) <= 1:
        return nums[0] if nums else 0
    return max(rob_linear(nums[:-1]), rob_linear(nums[1:]))
```

### Golang (optimized)

```go
func rob(nums []int) int {
	robLinear := func(arr []int) int {
		prev2, prev1 := 0, 0
		for _, x := range arr {
			cur := prev1
			if x+prev2 > cur {
				cur = x + prev2
			}
			prev2, prev1 = prev1, cur
		}
		return prev1
	}
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	a := robLinear(nums[:len(nums)-1])
	b := robLinear(nums[1:])
	if a > b {
		return a
	}
	return b
}
```

- **Time:** O(n)  
- **Space:** O(1)

---

## 4. Walkthrough

nums=[2,3,2]. nums[:-1]=[2,3] → rob=3. nums[1:]=[3,2] → rob=3. max=3.

nums=[1,2,3,1]. nums[:-1]=[1,2,3] → rob=4. nums[1:]=[2,3,1] → rob=4. max=4.

---

## 5. Edge Cases

1. **Single house:** return that value.  
2. **Two houses:** max(nums[0], nums[1]).  
3. **All same:** depends on parity.  
4. **First and last both high:** Must choose one; two passes handle it.

---

## 6. Pitfalls

1. **Including both first and last:** Breaks circular constraint.  
2. **Three passes:** Only need two linear passes.  
3. **Empty array:** return 0.

---

## 7. Follow-ups

1. **House Robber III (tree)?** DP on tree: (rob root, skip root).  
2. **K adjacent constraint?** More complex DP.  
3. **Print chosen houses?** Backtrack with parent pointers.

---

## 8. When to Use

- **Circular constraint:** Split into two linear problems (exclude first or last).  
- **Pattern:** Same as House Robber; reduce circle to two lines.
