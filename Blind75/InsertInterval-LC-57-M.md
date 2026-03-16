# Insert Interval — LeetCode 57 (M)

## 1. What is the problem?

**LeetCode 57 — Insert Interval**

Given non-overlapping intervals sorted by start, and a new interval, insert it and merge any overlapping intervals. Return the merged list.

**Example 1**

- **Input:** intervals = [[1,3],[6,9]], newInterval = [2,5]
- **Output:** [[1,5],[6,9]]

**Example 2**

- **Input:** intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
- **Output:** [[1,2],[3,10],[12,16]]

---

## 2. Brute Force Solution

Append newInterval to intervals, sort by start, then merge (same as Merge Intervals). O(n log n). Correct.

### Python (brute — sort and merge)

```python
def insert_brute(intervals: list[list[int]], new_interval: list[int]) -> list[list[int]]:
    intervals = intervals + [new_interval]
    intervals.sort(key=lambda x: x[0])
    out = [intervals[0]]
    for start, end in intervals[1:]:
        if start <= out[-1][1]:
            out[-1][1] = max(out[-1][1], end)
        else:
            out.append([start, end])
    return out
```

### Golang (brute)

```go
func insertBrute(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
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
- **Space:** O(n)

---

## 3. Optimized Solution (One Pass)

**Idea:** Intervals are sorted. Three phases: (1) add all intervals ending before newInterval starts; (2) merge overlapping with newInterval; (3) add remaining. O(n).

### Python (optimized)

```python
def insert(intervals: list[list[int]], new_interval: list[int]) -> list[list[int]]:
    out = []
    i, n = 0, len(intervals)
    start, end = new_interval
    while i < n and intervals[i][1] < start:
        out.append(intervals[i])
        i += 1
    while i < n and intervals[i][0] <= end:
        start = min(start, intervals[i][0])
        end = max(end, intervals[i][1])
        i += 1
    out.append([start, end])
    while i < n:
        out.append(intervals[i])
        i += 1
    return out
```

### Golang (optimized)

```go
func insert(intervals [][]int, newInterval []int) [][]int {
	var out [][]int
	start, end := newInterval[0], newInterval[1]
	i := 0
	for i < len(intervals) && intervals[i][1] < start {
		out = append(out, intervals[i])
		i++
	}
	for i < len(intervals) && intervals[i][0] <= end {
		if intervals[i][0] < start {
			start = intervals[i][0]
		}
		if intervals[i][1] > end {
			end = intervals[i][1]
		}
		i++
	}
	out = append(out, []int{start, end})
	for i < len(intervals) {
		out = append(out, intervals[i])
		i++
	}
	return out
}
```

- **Time:** O(n)  
- **Space:** O(n)

---

## 4. Walkthrough

intervals=[[1,3],[6,9]], new=[2,5]. Add [1,3] (3>=2, overlap). Merge: start=1, end=5. Add [6,9]. Result [[1,5],[6,9]].

---

## 5. Edge Cases

1. **Empty intervals:** return [newInterval].  
2. **New before all:** prepend.  
3. **New after all:** append.  
4. **New overlaps all:** single merged interval.

---

## 6. Pitfalls

1. **Overlap condition:** intervals[i][0] <= end (current merged end).  
2. **Merge:** Expand start and end during merge phase.  
3. **Don't sort:** Intervals already sorted; use one pass.

---

## 7. Follow-ups

1. **Insert multiple?** Sort all, then merge.  
2. **Remove interval?** Split/merge logic.  
3. **Range sum?** Different structure.

---

## 8. When to Use

- **Insert into sorted intervals:** One pass: before, merge, after.  
- **Pattern:** Same as Merge Intervals but optimized for single insert.
