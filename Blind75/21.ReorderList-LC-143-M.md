# Reorder List — LeetCode 143 (M)

## 1. What is the problem?

**LeetCode 143 — Reorder List**

Reorder a linked list: L0→L1→…→Ln-1 → L0→Ln-1→L1→Ln-2→… (first, last, second, second-last, …).

**Example 1**

- **Input:** head = [1,2,3,4]
- **Output:** [1,4,2,3]

**Example 2**

- **Input:** head = [1,2,3,4,5]
- **Output:** [1,5,2,4,3]

---

## 2. Brute Force Solution

Store all nodes in array. Reorder by alternating first/last. Rebuild list. O(n) time and space. Correct.

### Python (array approach)

```python
def reorder_list_brute(head: Optional["ListNode"]) -> None:
    nodes = []
    while head:
        nodes.append(head)
        head = head.next
    n = len(nodes)
    for i in range(n // 2):
        nodes[i].next = nodes[n - 1 - i]
        nodes[n - 1 - i].next = nodes[i + 1] if i + 1 < n - 1 - i else None
    if n % 2 == 1:
        nodes[n // 2].next = None
```

### Golang (optimal — find mid, reverse second half, merge)

**Idea:** Find middle (slow/fast). Reverse second half. Merge first half with reversed second half (alternating).

### Python (optimized)

```python
def reorder_list(head: Optional["ListNode"]) -> None:
    if not head or not head.next:
        return
    slow, fast = head, head
    while fast.next and fast.next.next:
        slow, fast = slow.next, fast.next.next
    second = slow.next
    slow.next = None
    prev = None
    while second:
        nxt = second.next
        second.next = prev
        prev, second = second, nxt
    first, second = head, prev
    while second:
        n1, n2 = first.next, second.next
        first.next = second
        second.next = n1
        first, second = n1, n2
```

### Golang (optimized)

```go
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	second := slow.Next
	slow.Next = nil
	var prev *ListNode
	for second != nil {
		next := second.Next
		second.Next = prev
		prev, second = second, next
	}
	first, second := head, prev
	for second != nil {
		n1, n2 := first.Next, second.Next
		first.Next = second
		second.Next = n1
		first, second = n1, n2
	}
}
```

- **Time:** O(n)  
- **Space:** O(1)

---

## 4. Walkthrough

[1,2,3,4]. Mid at 2. Second half [3,4] → reverse to [4,3]. Merge: 1→4→2→3, 3→nil. Result [1,4,2,3].

---

## 5. Edge Cases

1. **Single node:** no change.  
2. **Two nodes:** 1→2 stays.  
3. **Odd length:** Middle node ends list.  
4. **In-place:** Modify in place; no return.

---

## 6. Pitfalls

1. **Break at middle:** slow.next = None before reverse.  
2. **Merge order:** first, second, first.next, second.next.  
3. **Second half shorter or equal:** When second is nil, first may have one left; it's already correct.

---

## 7. Follow-ups

1. **Reorder in groups of k?** Similar; reverse k nodes at a time.  
2. **Palindrome list?** Same: reverse second half, compare.  
3. **Return new list?** Copy nodes instead of in-place.

---

## 8. When to Use

- **Reorder list (first, last, second, …):** Find mid, reverse second half, merge.  
- **Pattern:** Same as "palindrome linked list"; find mid + reverse + merge.
