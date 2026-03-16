# Find Median from Data Stream — LeetCode 295 (H)

## 1. What is the problem?

**LeetCode 295 — Find Median from Data Stream**

Implement MedianFinder: addNum(num) adds a number, findMedian() returns the median. If count is even, median is average of two middle elements.

**Example 1**

- **Input:** ["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]
  [[],[1],[2],[],[3],[]]
- **Output:** [null,null,null,1.5,null,2.0]

---

## 2. Brute Force Solution

Store all numbers in array. addNum: append, O(1). findMedian: sort and pick middle, O(n log n). Or maintain sorted list: addNum O(n), findMedian O(1).

### Python (two heaps)

**Idea:** Maintain lo (max-heap) and hi (min-heap). Keep len(lo) >= len(hi), at most 1 more. addNum: add to lo, pop max to hi; if len(lo) < len(hi), pop min from hi to lo. Median: lo[0] if odd, (lo[0]+hi[0])/2 if even.

```python
import heapq

class MedianFinder:
    def __init__(self):
        self.lo = []   # max-heap (negate for min-heap)
        self.hi = []   # min-heap

    def addNum(self, num: int) -> None:
        heapq.heappush(self.lo, -num)
        heapq.heappush(self.hi, -heapq.heappop(self.lo))
        if len(self.lo) < len(self.hi):
            heapq.heappush(self.lo, -heapq.heappop(self.hi))

    def findMedian(self) -> float:
        if len(self.lo) > len(self.hi):
            return -self.lo[0]
        return (-self.lo[0] + self.hi[0]) / 2
```

### Golang

```go
import "container/heap"

type maxHeap []int
func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MedianFinder struct {
	lo maxHeap
	hi minHeap
}

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

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (m *MedianFinder) AddNum(num int) {
	heap.Push(&m.lo, num)
	heap.Push(&m.hi, heap.Pop(&m.lo).(int))
	if m.lo.Len() < m.hi.Len() {
		heap.Push(&m.lo, heap.Pop(&m.hi).(int))
	}
}

func (m *MedianFinder) FindMedian() float64 {
	if m.lo.Len() > m.hi.Len() {
		return float64(m.lo[0])
	}
	return float64(m.lo[0]+m.hi[0]) / 2
}
```

- **Time:** addNum O(log n), findMedian O(1)  
- **Space:** O(n)

---

## 4. Walkthrough

add 1: lo=[1], hi=[]. add 2: lo=[1], hi=[2]; balance: lo=[1], hi=[2]. Median: (1+2)/2=1.5. add 3: lo=[1,3]? Push 3 to lo, pop to hi: lo=[1], hi=[2,3]. lo<hi, pop hi to lo: lo=[1,2], hi=[3]. Median: lo.top=2.

---

## 5. Edge Cases

1. **Single element:** Median = that element.  
2. **Two elements:** Average.  
3. **Large numbers:** Use float for average.  
4. **Negative numbers:** Heaps work.

---

## 6. Pitfalls

1. **Balance:** lo has >= hi; at most 1 more.  
2. **Python heapq:** Min-heap only; negate for max.  
3. **Golang:** Implement heap.Interface for max heap.

---

## 7. Follow-ups

1. **Find 99th percentile?** Similar; adjust heap sizes.  
2. **Sliding window median?** Two heaps + lazy delete or multiset.  
3. **Integers only?** Return float for even case.

---

## 8. When to Use

- **Streaming median:** Two heaps (max for lower half, min for upper).  
- **Pattern:** Same for "Sliding Window Median", "IPO" (k largest).
