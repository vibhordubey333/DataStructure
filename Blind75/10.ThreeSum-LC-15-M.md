# 3Sum — LeetCode 15 (M)

## 1. What is the problem?

**LeetCode 15 — 3Sum**

Given an integer array `nums`, return all triplets `[nums[i], nums[j], nums[k]]` such that `i != j`, `i != k`, `j != k` and `nums[i] + nums[j] + nums[k] == 0`. Solution set must not contain duplicate triplets.

**Example 1**

- **Input:** `nums = [-1, 0, 1, 2, -1, -4]`
- **Output:** `[[-1, -1, 2], [-1, 0, 1]]`

**Example 2**

- **Input:** `nums = [0, 1, 1]`
- **Output:** `[]`

---

## 2. Brute Force Solution

Enumerate all triplets (i, j, k) with i < j < k, check if sum is 0, and deduplicate (e.g. by sorting each triplet and using a set). Correct but slow.

### Python (brute force)

```python
def three_sum_brute(nums: list[int]) -> list[list[int]]:
    n = len(nums)
    seen = set()
    out = []
    for i in range(n):
        for j in range(i + 1, n):
            for k in range(j + 1, n):
                if nums[i] + nums[j] + nums[k] == 0:
                    key = tuple(sorted([nums[i], nums[j], nums[k]]))
                    if key not in seen:
                        seen.add(key)
                        out.append([nums[i], nums[j], nums[k]])
    return out
```

### Golang (brute force)

```go
func threeSumBrute(nums []int) [][]int {
	var out [][]int
	seen := make(map[[3]int]bool)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					a := [3]int{nums[i], nums[j], nums[k]}
					sort.Ints(a[:])
					if !seen[a] {
						seen[a] = true
						out = append(out, []int{nums[i], nums[j], nums[k]})
					}
				}
			}
		}
	}
	return out
}
```

- **Time:** O(n³) + dedup  
- **Space:** O(1) + set size

---

## 3. Optimized Solution (Sort + Two Pointers)

**Idea:** Sort the array. Fix the first number `nums[i]`, then find two numbers in the rest that sum to `-nums[i]` using two pointers (j from i+1, k from end). Skip duplicates for i, j, and k.

**Steps:**

1. Sort `nums`.
2. For each index `i` from 0 to n−3:  
   - If `i > 0` and `nums[i] == nums[i-1]`, skip (duplicate first element).  
   - Let `target = -nums[i]`. Two pointers: `j = i+1`, `k = n-1`.  
   - While `j < k`:  
     - If `nums[j] + nums[k] == target`: add `[nums[i], nums[j], nums[k]]`; then advance j and k, skipping duplicates.  
     - Else if `nums[j] + nums[k] < target`: j++  
     - Else: k--  
3. Return collected triplets.

### Python (optimized)

```python
def three_sum(nums: list[int]) -> list[list[int]]:
    nums = sorted(nums)
    n = len(nums)
    out = []
    for i in range(n - 2):
        if i > 0 and nums[i] == nums[i - 1]:
            continue
        target = -nums[i]
        j, k = i + 1, n - 1
        while j < k:
            s = nums[j] + nums[k]
            if s == target:
                out.append([nums[i], nums[j], nums[k]])
                j += 1
                while j < k and nums[j] == nums[j - 1]:
                    j += 1
                k -= 1
                while j < k and nums[k] == nums[k + 1]:
                    k -= 1
            elif s < target:
                j += 1
            else:
                k -= 1
    return out
```

### Golang (optimized)

```go
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var out [][]int
	n := len(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		j, k := i+1, n-1
		for j < k {
			s := nums[j] + nums[k]
			if s == target {
				out = append(out, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			} else if s < target {
				j++
			} else {
				k--
			}
		}
	}
	return out
}
```

- **Time:** O(n²) (sort O(n log n) + outer loop × two pointers)  
- **Space:** O(1) excluding output; O(log n) for sort if not in-place

---

## 4. Walkthrough

**Example 1:** `nums = [-1, 0, 1, 2, -1, -4]` → sorted: `[-4, -1, -1, 0, 1, 2]`

- i=0: nums[i]=-4, target=4; j=1,k=5: -1+2=1<4 → j=2; -1+2=1<4 → j=3; 0+2=2<4 → j=4; 1+2=3<4 → j=5; no triplet.
- i=1: nums[i]=-1, target=1; j=2,k=5: -1+2=1 → add [-1,-1,2]; skip dup j; j=4,k=4 done. j=3,k=5: 0+2=2>1 → k=4; 0+1=1 → add [-1,0,1]. 
- Result: `[[-1,-1,2], [-1,0,1]]`.

---

## 5. Edge Cases

1. **Fewer than 3 elements:** return [].  
2. **All zeros:** one triplet [0,0,0].  
3. **No solution:** return [].  
4. **Duplicate triplets:** skip duplicate i, and after finding a pair skip duplicate j/k.  
5. **Negative and positive:** sort makes two-pointer valid.

---

## 6. Pitfalls

1. **Duplicate triplets:** Must skip duplicate first element (i) and after adding a pair skip duplicate j and k.  
2. **Using hash map for two sum:** Works but need to avoid duplicates; two pointers after sort is cleaner.  
3. **Forgetting to sort:** Two-pointer for “two sum” in rest array requires sorted array.

---

## 7. Follow-ups

1. **4Sum (LC 18)?** Fix two indices, then two pointers for the remaining two.  
2. **k-Sum?** Reduce to (k-1)-Sum recursively; base case 2Sum with two pointers.  
3. **Closest sum (3Sum closest)?** Track closest sum and update during two-pointer.  
4. **Return indices?** Store indices in triplets instead of values (and handle dup by index).

---

## 8. When to Use

- **Three sum / k sum:** Sort + fix (k-2) numbers + two pointers for the last two.  
- **Pattern:** “Unique triplets” + “sum to target” → sort + two pointers + duplicate skipping.
