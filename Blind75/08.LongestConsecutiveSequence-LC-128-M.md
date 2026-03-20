# Longest Consecutive Sequence — LeetCode 128 (M)

## 1. What is the problem?

**LeetCode 128 — Longest Consecutive Sequence**

Given an unsorted integer array nums, return the length of the longest consecutive elements sequence (numbers that appear in nums and form consecutive integers). O(n) time required.

**Example 1**

- **Input:** nums = [100, 4, 200, 1, 3, 2]
- **Output:** 4 (sequence [1,2,3,4])

**Example 2**

- **Input:** nums = [0, 3, 7, 2, 5, 8, 4, 6, 0, 1]
- **Output:** 9

---

## 2. Brute Force Solution

For each number x, check if x+1, x+2, … are in the array (using a set); track length. Repeat for each x. O(n²) if we scan for every starting point. Or sort and scan for longest run: O(n log n). Correct but not O(n).

### Python (sort — O(n log n))

```python
def longest_consecutive_brute(nums: list[int]) -> int:
    if not nums:
        return 0
    nums = sorted(set(nums))
    best = cur = 1
    for i in range(1, len(nums)):
        if nums[i] == nums[i-1] + 1:
            cur += 1
        else:
            cur = 1
        best = max(best, cur)
    return best
```

### Golang (sort)

```go
func longestConsecutiveBrute(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	best, cur := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			cur++
		} else if nums[i] != nums[i-1] {
			cur = 1
		}
		if cur > best {
			best = cur
		}
	}
	return best
}
```

- **Time:** O(n log n)  
- **Space:** O(1) if sort in place

---

## 3. Optimized Solution (Set + “start of sequence” only)

**Idea:** Put all numbers in a set. A number x is the **start** of a consecutive sequence if x-1 is not in the set. For each start, extend by checking x+1, x+2, … and count length. Each number is part of exactly one sequence and we only extend from starts, so each element is visited at most twice → O(n).

**Steps:**

1. s = set(nums).  
2. best = 0. For each x in s: If x-1 not in s (x is a start): cur = 0, while x+cur in s: cur += 1. best = max(best, cur).  
3. Return best.

### Python (optimized)

```python
def longest_consecutive(nums: list[int]) -> int:
    s = set(nums)
    best = 0
    for x in s:
        if x - 1 not in s:
            cur = 0
            while x + cur in s:
                cur += 1
            best = max(best, cur)
    return best
```

### Golang (optimized)

```go
func longestConsecutive(nums []int) int {
	s := make(map[int]bool)
	for _, x := range nums {
		s[x] = true
	}
	best := 0
	for x := range s {
		if !s[x-1] {
			cur := 0
			for s[x+cur] {
				cur++
			}
			if cur > best {
				best = cur
			}
		}
	}
	return best
}
```

- **Time:** O(n)  
- **Space:** O(n)

---

## 4. Walkthrough

nums = [100, 4, 200, 1, 3, 2]. s = {1,2,3,4,100,200}. Starts: 1 (no 0), 100 (no 99), 200 (no 199). From 1: 1,2,3,4 → 4. From 100: 1. From 200: 1. Best = 4.

---

## 5. Edge Cases

1. **Empty array:** 0.  
2. **Single element:** 1.  
3. **All same:** 1 (no consecutive increase).  
4. **Duplicates:** Set removes them; correct.  
5. **Negative numbers:** Set and arithmetic work.

---

## 6. Pitfalls

1. **Only extending from every number:** Then we’d count same sequence many times; only extend from “start” (x-1 not in set).  
2. **Sorting:** O(n log n); set + start check is O(n).  
3. **Sequence of one:** Length 1 is valid.

---

## 7. Follow-ups

1. **Return the sequence?** Track start of best run and build list.  
2. **Consecutive in sorted order but not necessarily adjacent in array?** Same problem.  
3. **Longest arithmetic subsequence?** Different: DP with difference as key.

---

## 8. When to Use

- **Longest consecutive in unsorted array:** Set + only extend from sequence start (x-1 not in set).  
- **Pattern:** Set for O(1) lookup; “start of chain” to get O(n) total work.
