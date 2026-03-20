# Meeting Rooms — LeetCode 252 (E, Premium)

## 1. What is the problem?

**LeetCode 252 — Meeting Rooms**

Given an array of meeting time intervals where intervals[i] = [start_i, end_i], determine if a person could attend all meetings.

**Example 1**

- **Input:** intervals = [[0,30],[5,10],[15,20]]
- **Output:** false
- **Explanation:** [0,30] and [5,10] overlap.

**Example 2**

- **Input:** intervals = [[7,10],[2,4]]
- **Output:** true

---

## 2. Brute Force / Standard Solution

Sort by start. Check if intervals[i].end > intervals[i+1].start for any i. If so, overlap → false.

### Python

```python
def can_attend_meetings(intervals: list[list[int]]) -> bool:
    intervals.sort(key=lambda x: x[0])
    for i in range(len(intervals) - 1):
        if intervals[i][1] > intervals[i + 1][0]:
            return False
    return True
```

### Golang

```go
func canAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][1] > intervals[i+1][0] {
			return false
		}
	}
	return true
}
```

- **Time:** O(n log n)  
- **Space:** O(1)

---

## 4. Walkthrough

Sort: [2,4], [7,10]. 4 <= 7 ✓. True.

---

## 5. Edge Cases

1. **Empty:** true.  
2. **Single interval:** true.  
3. **Touching [0,5],[5,10]:** No overlap (end is exclusive per typical spec).  
4. **All overlap:** false.

---

## 6. Pitfalls

1. **Sort by start:** Ensures we only check adjacent pairs.  
2. **Boundary:** [a,b] and [b,c] — usually non-overlapping if end exclusive.  
3. **Equal start:** Sort by end as tiebreaker if needed.

---

## 7. Follow-ups

1. **Meeting Rooms II (min rooms)?** Sweep line or min-heap.  
2. **Return conflicting pairs?** Compare all pairs or use sweep.  
3. **Merge overlapping?** Merge Intervals.

---

## 8. When to Use

- **Check if intervals overlap:** Sort by start; check adjacent.  
- **Pattern:** Same as "Merge Intervals", "Non-overlapping Intervals".
