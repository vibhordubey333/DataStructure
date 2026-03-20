# Meeting Rooms II — LeetCode 253 (M, Premium)

## 1. What is the problem?

**LeetCode 253 — Meeting Rooms II**

Given an array of meeting time intervals where intervals[i] = [start_i, end_i], return the minimum number of conference rooms required.

**Example 1**

- **Input:** intervals = [[0,30],[5,10],[15,20]]
- **Output:** 2
- **Explanation:** [0,30] and [5,10] overlap; need 2 rooms.

**Example 2**

- **Input:** intervals = [[7,10],[2,4]]
- **Output:** 1

---

## 2. Brute Force / Standard Solution

**Sweep line:** Split into (time, +1) for start and (time, -1) for end. Sort by time; if tie, process ends first. Track count, max count = rooms.

**Alternative:** Min-heap of end times. Sort by start. For each interval: if heap min <= start, pop (room freed). Push end. Size of heap = rooms needed.

### Python (min-heap)

```python
import heapq

def min_meeting_rooms(intervals: list[list[int]]) -> int:
    intervals.sort(key=lambda x: x[0])
    heap = []
    for s, e in intervals:
        if heap and heap[0] <= s:
            heapq.heappop(heap)
        heapq.heappush(heap, e)
    return len(heap)
```

### Golang

```go
import "container/heap"

type minHeap []int
func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	h := &minHeap{}
	heap.Init(h)
	for _, iv := range intervals {
		s, e := iv[0], iv[1]
		if h.Len() > 0 && (*h)[0] <= s {
			heap.Pop(h)
		}
		heap.Push(h, e)
	}
	return h.Len()
}
```

- **Time:** O(n log n)  
- **Space:** O(n)

---

## 4. Walkthrough

Sort: [0,30],[5,10],[15,20]. 0: push 30, heap=[30], rooms=1. 5: 30>5, push 10, heap=[10,30], rooms=2. 15: 10<=15, pop, push 20, heap=[20,30], rooms=2. Max=2.

---

## 5. Edge Cases

1. **Empty:** 0.  
2. **Single:** 1.  
3. **All overlap:** n rooms.  
4. **No overlap:** 1 room.

---

## 6. Pitfalls

1. **Tie break:** End before start at same time (room freed first).  
2. **Heap stores end times:** Min = earliest ending meeting.  
3. **Sweep line:** Alternative; sort events, +1/-1.

---

## 7. Follow-ups

1. **Return schedule?** Assign room per interval.  
2. **Meeting Rooms I?** Just check overlap.  
3. **Maximum overlap count?** Sweep line, track max.

---

## 8. When to Use

- **Min rooms for intervals:** Min-heap of end times or sweep line.  
- **Pattern:** Same as "Merge Intervals" style; heap for "earliest free" room.
