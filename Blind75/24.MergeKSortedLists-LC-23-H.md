# Merge k Sorted Lists — LeetCode 23 (H)

## 1. What is the problem?

**LeetCode 23 — Merge k Sorted Lists**

Given an array of k linked lists, each sorted in ascending order, merge them into one sorted list.

**Example 1**

- **Input:** lists = [[1,4,5],[1,3,4],[2,6]]
- **Output:** [1,1,2,3,4,4,5,6]

**Example 2**

- **Input:** lists = []
- **Output:** []

---

## 2. Brute Force Solution

Collect all node values, sort, rebuild one list. O(N log N) where N = total nodes. Correct but ignores sorted structure.

### Python (brute force)

```python
def merge_k_lists_brute(lists: list[Optional["ListNode"]]) -> Optional["ListNode"]:
    vals = []
    for head in lists:
        while head:
            vals.append(head.val)
            head = head.next
    vals.sort()
    dummy = ListNode(0)
    p = dummy
    for v in vals:
        p.next = ListNode(v)
        p = p.next
    return dummy.next
```

### Golang (brute force)

```go
func mergeKListsBrute(lists []*ListNode) *ListNode {
	var vals []int
	for _, head := range lists {
		for head != nil {
			vals = append(vals, head.Val)
			head = head.Next
		}
	}
	sort.Ints(vals)
	dummy := &ListNode{}
	p := dummy
	for _, v := range vals {
		p.Next = &ListNode{Val: v}
		p = p.Next
	}
	return dummy.Next
}
```

- **Time:** O(N log N)  
- **Space:** O(N)

---

## 3. Optimized Solution (Min-Heap)

**Idea:** Maintain the current head of each list in a min-heap. Repeatedly pop the smallest, append to result, push its next if not nil. O(N log k) where k = number of lists.

**Steps:**

1. Push all non-nil heads into a min-heap (by node value; break ties by list index if needed).  
2. While heap not empty: pop smallest node, append to result list, if node.next push it.  
3. Return merged list.

### Python (optimized — heap)

```python
import heapq
def merge_k_lists(lists: list[Optional["ListNode"]]) -> Optional["ListNode"]:
    heap = []
    for i, head in enumerate(lists):
        if head:
            heapq.heappush(heap, (head.val, i, head))
    dummy = ListNode(0)
    p = dummy
    while heap:
        _, i, node = heapq.heappop(heap)
        p.next = node
        p = p.next
        if node.next:
            heapq.heappush(heap, (node.next.val, i, node.next))
    return dummy.next
```

(Note: In Python 3, tuple (val, i, node) is used because ListNode is not comparable; i breaks ties.)

### Golang (optimized — container/heap)

```go
type minHeap []*ListNode
func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h minHeap) Swap(i, j int)     { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)       { *h = append(*h, x.(*ListNode)) }
func (h *minHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}
func mergeKLists(lists []*ListNode) *ListNode {
	h := &minHeap{}
	heap.Init(h)
	for _, head := range lists {
		if head != nil {
			heap.Push(h, head)
		}
	}
	dummy := &ListNode{}
	p := dummy
	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		p.Next = node
		p = p.Next
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}
	return dummy.Next
}
```

- **Time:** O(N log k)  
- **Space:** O(k) for heap

---

## 4. Walkthrough

lists = [[1,4,5],[1,3,4],[2,6]]. Heap: (1,0),(1,1),(2,2). Pop 1 from list0, push (4,…); pop 1 from list1, push (3,…); … Result: 1,1,2,3,4,4,5,6.

---

## 5. Edge Cases

1. **lists = []:** return nil.  
2. **All lists empty:** return nil.  
3. **One list:** return that list.  
4. **Many empty lists:** Filter nil heads before heap.

---

## 6. Pitfalls

1. **Comparing nodes in heap:** Python: use (val, id, node) so nodes are comparable.  
2. **Pushing nil:** Check node.next before pushing.  
3. **Divide and conquer:** Merge pairs (LC 21) repeatedly; also O(N log k) but different constant.

---

## 7. Follow-ups

1. **External sort (huge lists on disk)?** Same: heap of k heads.  
2. **Merge two sorted arrays?** Two pointers; for k arrays use heap.  
3. **Iterator interface?** Same heap approach with “next” returning next value.

---

## 8. When to Use

- **Merge k sorted streams:** Min-heap of k heads.  
- **Pattern:** K-way merge; also divide-and-conquer merge (merge pairs until one list).
