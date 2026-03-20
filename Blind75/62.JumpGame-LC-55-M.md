# Jump Game — LeetCode 55 (M)

## 1. What is the problem?

**LeetCode 55 — Jump Game**

Given array `nums` where nums[i] is max jump length from index i, return true if you can reach the last index from index 0.

**Example 1**

- **Input:** nums = [2, 3, 1, 1, 4]
- **Output:** true

**Example 2**

- **Input:** nums = [3, 2, 1, 0, 4]
- **Output:** false

---

## 2. Brute Force Solution

Recurse: from i, try jumping 1 to nums[i] steps; if any reaches end, return true. Exponential.

### Python (brute force)

```python
def can_jump_brute(nums: list[int]) -> bool:
    def f(i: int) -> bool:
        if i >= len(nums) - 1:
            return True
        for j in range(1, nums[i] + 1):
            if f(i + j):
                return True
        return False
    return f(0)
```

### Golang (brute force)

```go
func canJumpBrute(nums []int) bool {
	var f func(int) bool
	f = func(i int) bool {
		if i >= len(nums)-1 {
			return true
		}
		for j := 1; j <= nums[i] && i+j < len(nums); j++ {
			if f(i + j) {
				return true
			}
		}
		return false
	}
	return f(0)
}
```

- **Time:** O(2^n)  
- **Space:** O(n) stack

---

## 3. Optimized Solution (Greedy)

**Idea:** Track the farthest index we can reach. For each i, if i > farthest, we're stuck → false. Else update farthest = max(farthest, i + nums[i]). If farthest >= n-1, true.

### Python (optimized)

```python
def can_jump(nums: list[int]) -> bool:
    farthest = 0
    for i in range(len(nums)):
        if i > farthest:
            return False
        farthest = max(farthest, i + nums[i])
        if farthest >= len(nums) - 1:
            return True
    return True
```

### Golang (optimized)

```go
func canJump(nums []int) bool {
	farthest := 0
	for i := 0; i < len(nums); i++ {
		if i > farthest {
			return false
		}
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}
		if farthest >= len(nums)-1 {
			return true
		}
	}
	return true
}
```

- **Time:** O(n)  
- **Space:** O(1)

---

## 4. Walkthrough

nums=[2,3,1,1,4]. i=0: farthest=2. i=1: farthest=4 ≥ 4 → true.

nums=[3,2,1,0,4]. i=0: farthest=3. i=1: farthest=3. i=2: farthest=3. i=3: farthest=3, i=3>3? No, 3+0=3. i=4: i=4>3 → false.

---

## 5. Edge Cases

1. **Single element:** true.  
2. **nums[0]=0, n>1:** false.  
3. **All zeros except last:** false if n>1.  
4. **nums[0]>=n-1:** true immediately.

---

## 6. Pitfalls

1. **Check i > farthest first:** Before using nums[i].  
2. **Early exit:** When farthest >= n-1.  
3. **Jump Game II (min jumps):** Different: BFS or greedy with "level".

---

## 7. Follow-ups

1. **Jump Game II (min jumps)?** BFS or greedy: track current level end, next level end.  
2. **Jump Game III (can reach index with value 0)?** DFS/BFS from start.  
3. **Jump Game IV?** BFS with same-value jumps.

---

## 8. When to Use

- **Reachability with max step:** Greedy "farthest we can reach".  
- **Pattern:** One-pass greedy; same idea in "jump game II" for min steps.
