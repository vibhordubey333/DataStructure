# Merge Intervals — LeetCode 56 (M)

## 1. What is the problem?

**LeetCode 56 — Merge Intervals**

Given an array of intervals where intervals[i] = [start_i, end_i], merge all overlapping intervals and return non-overlapping intervals that cover the same set.

**Example 1**

- **Input:** intervals = [[1,3],[2,6],[8,10],[15,18]]
- **Output:** [[1,6],[8,10],[15,18]]

**Example 2**

- **Input:** intervals = [[1,4],[4,5]]
- **Output:** [[1,5]]

---

## 2. Brute Force Solution

Compare every pair of intervals; if they overlap, merge and replace with merged interval; repeat until no merges. O(n²) or more. Or: sort by start, then traverse and merge adjacent if overlapping. The sort-and-merge is the standard and optimal.

### Python (sort + merge — optimal)

```python
def merge(intervals: list[list[int]]) -> list[list[int]]:
    if not intervals:
        return []
    intervals.sort(key=lambda x: x[0])
    out = [intervals[0]]
    for start, end in intervals[1:]:
        if start <= out[-1][1]:
            out[-1][1] = max(out[-1][1], end)
        else:
            out.append([start, end])
    return out
```

### Golang (optimized)

```go
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	out := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		start, end := intervals[i][0], intervals[i][1]
		if start <= out[len(out)-1][1] {
			if end > out[len(out)-1][1] {
				out[len(out)-1][1] = end
			}
		} else {
			out = append(out, intervals[i])
		}
	}
	return out
}
```

- **Time:** O(n log n)  
- **Space:** O(log n) sort; O(1) or O(n) for output

---

## 3. Optimized Solution

Sort by start then one pass merge is optimal. No better than O(n log n) for comparison-based sort.

---

## 4. Walkthrough

Sorted: [1,3], [2,6], [8,10], [15,18]. out = [[1,3]]. [2,6]: 2 <= 3 → merge to [1,6]. [8,10]: 8 > 6 → append [8,10]. [15,18]: append [15,18]. Result [[1,6],[8,10],[15,18]].

---

## 5. Edge Cases

1. **Empty list:** return [].  
2. **Single interval:** return as is.  
3. **Fully contained:** [1,5], [2,3] → [1,5].  
4. **Touch at endpoint:** [1,4],[4,5] → [1,5] (overlap by “<=”).  
5. **Negative numbers:** Same logic.

---

## 6. Pitfalls

1. **Overlap condition:** start_i <= out[-1][1] (not strictly less) so [1,4] and [4,5] merge.  
2. **Updating end:** Use max(current_end, new_end) when merging.  
3. **Sort key:** Sort by start (or by start then end).

---

## 7. Follow-ups

1. **Insert interval (LC 57)?** Binary search for position, merge overlapping.  
2. **Non-overlapping intervals (LC 435)?** Sort by end; greedy remove overlapping.  
3. **Meeting rooms (LC 252)?** Sort and check adjacent for overlap.  
4. **Meeting rooms II (LC 253)?** Min-heap of end times or sweep line.

---

## 8. When to Use

- **Overlapping intervals → merge:** Sort by start, then one pass merge.  
- **Pattern:** Interval sorting; same in insert interval, erase overlap, meeting rooms.
