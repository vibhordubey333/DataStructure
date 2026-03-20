# Non-overlapping Intervals — LeetCode 435 (M)

## 1. What is the problem?

**LeetCode 435 — Non-overlapping Intervals**

Given an array of intervals, return the **minimum number of intervals to remove** so the remaining intervals are non-overlapping.

**Example 1**

- **Input:** intervals = [[1,2],[2,3],[3,4],[1,3]]
- **Output:** 1 (remove [1,3])

**Example 2**

- **Input:** intervals = [[1,2],[1,2],[1,2]]
- **Output:** 2

---

## 2. Brute Force Solution

Try all subsets; find largest subset of non-overlapping intervals. O(2^n). Or: sort by end, greedy keep if doesn't overlap last kept. Greedy is optimal.

### Python (greedy — optimal)

**Idea:** Sort by end time. Keep interval if its start >= end of last kept. Count removals = n - kept.

### Python (optimized)

```python
def erase_overlap_intervals(intervals: list[list[int]]) -> int:
    if not intervals:
        return 0
    intervals.sort(key=lambda x: x[1])
    kept = 1
    last_end = intervals[0][1]
    for start, end in intervals[1:]:
        if start >= last_end:
            kept += 1
            last_end = end
    return len(intervals) - kept
```

### Golang (optimized)

```go
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
	kept := 1
	lastEnd := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		start, end := intervals[i][0], intervals[i][1]
		if start >= lastEnd {
			kept++
			lastEnd = end
		}
	}
	return len(intervals) - kept
}
```

- **Time:** O(n log n)  
- **Space:** O(1)

---

## 3. Optimized Solution

Same greedy: sort by end, keep non-overlapping. Maximizing kept = minimizing removed.

---

## 4. Walkthrough

Sort by end: [[1,2],[2,3],[1,3],[3,4]]. Keep [1,2], last_end=2. [2,3]: 2>=2, keep. [1,3]: 1<2, skip. [3,4]: 3>=2, keep. Kept=3, remove=1.

---

## 5. Edge Cases

1. **Empty:** 0.  
2. **Single interval:** 0 removals.  
3. **All overlap:** Remove n-1.  
4. **All same:** Remove n-1.

---

## 6. Pitfalls

1. **Sort by end:** Greedy works when sorted by end (earliest finish first).  
2. **Sort by start:** Would need different logic.  
3. **Touching [1,2],[2,3]:** Not overlapping; keep both.

---

## 7. Follow-ups

1. **Maximum intervals in non-overlapping set?** Same: return kept.  
2. **Minimum interval cover?** Different problem.  
3. **Meeting rooms II (min rooms)?** Min-heap of end times.

---

## 8. When to Use

- **Interval scheduling / activity selection:** Sort by end, greedy keep.  
- **Pattern:** "Maximize non-overlapping" = "minimize removals".
