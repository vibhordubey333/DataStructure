# Missing Number — LeetCode 268 (E)

## 1. What is the problem?

**LeetCode 268 — Missing Number**

Given an array `nums` containing n distinct numbers in the range [0, n], return the only number in [0, n] that is not in `nums`. (n = nums.length.)

**Example 1**

- **Input:** nums = [3, 0, 1]
- **Output:** 2

**Example 2**

- **Input:** nums = [0, 1]
- **Output:** 2

---

## 2. Brute Force Solution

For each k in 0..n, check if k is in the array (e.g. with a set or linear scan). Return the one that is missing.

### Python (brute force)

```python
def missing_number_brute(nums: list[int]) -> int:
    n = len(nums)
    for k in range(n + 1):
        if k not in nums:
            return k
    return n
```

### Golang (brute force)

```go
func missingNumberBrute(nums []int) int {
	n := len(nums)
	seen := make(map[int]bool)
	for _, x := range nums {
		seen[x] = true
	}
	for k := 0; k <= n; k++ {
		if !seen[k] {
			return k
		}
	}
	return n
}
```

- **Time:** O(n) with set  
- **Space:** O(n)

---

## 3. Optimized Solution

**Idea:** Sum of 0..n is n*(n+1)/2. Sum of nums is sum(nums). Missing = (n*(n+1)/2) - sum(nums). Or: use XOR — xor all indices 0..n and all values in nums; duplicates cancel, missing number remains.

**Steps (math):**

1. n = len(nums). expected_sum = n * (n + 1) // 2.
2. return expected_sum - sum(nums).

**Steps (XOR, O(1) extra space):**

1. x = 0. For i in 0..n-1: x ^= i ^ nums[i]. Then x ^= n.
2. Return x.

### Python (optimized — sum)

```python
def missing_number(nums: list[int]) -> int:
    n = len(nums)
    return n * (n + 1) // 2 - sum(nums)
```

### Python (XOR)

```python
def missing_number_xor(nums: list[int]) -> int:
    n = len(nums)
    x = 0
    for i in range(n):
        x ^= i ^ nums[i]
    return x ^ n
```

### Golang (optimized)

```go
func missingNumber(nums []int) int {
	n := len(nums)
	sum := 0
	for _, x := range nums {
		sum += x
	}
	return n*(n+1)/2 - sum
}
```

```go
func missingNumberXOR(nums []int) int {
	x := 0
	for i, v := range nums {
		x ^= i ^ v
	}
	return x ^ len(nums)
}
```

- **Time:** O(n)  
- **Space:** O(1)

---

## 4. Walkthrough

nums = [3, 0, 1], n = 3. Expected sum = 6. Actual sum = 4. Missing = 2.

---

## 5. Edge Cases

1. **Missing 0:** e.g. [1,2] → 0.  
2. **Missing n:** e.g. [0,1] → 2.  
3. **Single element [0]:** n=1, missing 1.  
4. **Overflow:** n*(n+1)/2 may overflow in some languages; Python is fine; Go use int64 if needed.

---

## 6. Pitfalls

1. **Including n in range:** Numbers are in [0, n], so there are n+1 numbers and n elements; one missing.  
2. **Sum overflow:** For very large n, use XOR.  
3. **Modifying array:** Not needed; O(1) space with sum or XOR.

---

## 7. Follow-ups

1. **Two missing numbers?** Sum and sum of squares give two equations.  
2. **Array not distinct?** Different problem; use cycle sort or count array.  
3. **Stream of numbers?** Keep running sum and expected sum (need n known).

---

## 8. When to Use

- **One missing in 0..n:** Sum formula or XOR.  
- **Pattern:** Math (sum) or bit (XOR) to find missing/single/duplicate.
