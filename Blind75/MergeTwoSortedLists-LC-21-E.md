# Merge Two Sorted Lists — LeetCode 21 (E)

## 1. What is the problem?

**LeetCode 21 — Merge Two Sorted Lists**

Merge two sorted linked lists into one sorted list. Return the head of the merged list.

**Example 1**

- **Input:** l1 = [1,2,4], l2 = [1,3,4]
- **Output:** [1,1,2,3,4,4]

**Example 2**

- **Input:** l1 = [], l2 = []
- **Output:** []

---

## 2. Brute Force Solution

Collect all values from both lists, sort, then build a new list. Correct but O(n log n) and doesn’t use the fact that lists are already sorted. The natural “merge” is already optimal.

### Python (optimal merge — standard approach)

```python
def merge_two_lists(l1: Optional["ListNode"], l2: Optional["ListNode"]) -> Optional["ListNode"]:
    dummy = ListNode(0)
    p = dummy
    while l1 and l2:
        if l1.val <= l2.val:
            p.next = l1
            l1 = l1.next
        else:
            p.next = l2
            l2 = l2.next
        p = p.next
    p.next = l1 or l2
    return dummy.next
```

### Golang (optimal)

```go
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 != nil {
		p.Next = l1
	} else {
		p.Next = l2
	}
	return dummy.Next
}
```

- **Time:** O(n + m)  
- **Space:** O(1)

---

## 3. Optimized Solution

Same as above: two-pointer merge with a dummy node. No asymptotically better approach.

---

## 4. Walkthrough

l1=1→2→4, l2=1→3→4. Attach 1 from l1, then 1 from l2, then 2, 3, 4, 4. Result: 1→1→2→3→4→4.

---

## 5. Edge Cases

1. **One list empty:** return the other.  
2. **Both empty:** return nil.  
3. **One list much longer:** remaining tail attached with p.next = l1 or l2.

---

## 6. Pitfalls

1. **Using new nodes:** Not needed; reuse existing nodes.  
2. **Forgetting dummy:** Simplifies handling empty and first node.  
3. **Strict < vs <=:** Use <= so merge is stable and deterministic.

---

## 7. Follow-ups

1. **Merge k sorted lists (LC 23)?** Min-heap of k heads, or divide and conquer merge.  
2. **Merge two sorted arrays (in place)?** Start from end of larger array.  
3. **Sort list (LC 148)?** Merge sort: find mid, sort halves, merge.

---

## 8. When to Use

- **Merging two sorted sequences:** Two pointers + dummy.  
- **Pattern:** Base for merge sort on linked list and for merge k lists.
