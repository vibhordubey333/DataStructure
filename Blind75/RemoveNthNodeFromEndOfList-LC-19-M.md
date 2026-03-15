# Remove Nth Node From End of List — LeetCode 19 (M)

## 1. What is the problem?

**LeetCode 19 — Remove Nth Node From End of List**

Given the head of a linked list, remove the n-th node from the end and return the head. (n is valid, 1-indexed from end.)

**Example 1**

- **Input:** head = [1,2,3,4,5], n = 2
- **Output:** [1,2,3,5]

**Example 2**

- **Input:** head = [1], n = 1
- **Output:** []

---

## 2. Brute Force Solution

Two passes: first pass count length L; second pass skip the (L - n + 1)-th node from the start. Correct.

### Python (two pass)

```python
def remove_nth_from_end_brute(head: Optional["ListNode"], n: int) -> Optional["ListNode"]:
    length = 0
    p = head
    while p:
        length += 1
        p = p.next
    if n == length:
        return head.next
    p = head
    for _ in range(length - n - 1):
        p = p.next
    p.next = p.next.next
    return head
```

### Golang (two pass)

```go
func removeNthFromEndBrute(head *ListNode, n int) *ListNode {
	length := 0
	for p := head; p != nil; p = p.Next {
		length++
	}
	if n == length {
		return head.Next
	}
	p := head
	for i := 0; i < length-n-1; i++ {
		p = p.Next
	}
	p.Next = p.Next.Next
	return head
}
```

- **Time:** O(L)  
- **Space:** O(1)

---

## 3. Optimized Solution (One Pass — Two Pointers)

**Idea:** Place two pointers n+1 apart: first and second. When first reaches the end (past last node), second is right before the node to delete. Use dummy so we can remove the head cleanly.

**Steps:**

1. dummy = ListNode(0); dummy.next = head; first = second = dummy.  
2. Advance first n+1 times (so first is n+1 steps ahead).  
3. While first: first, second = first.next, second.next.  
4. second.next = second.next.next. Return dummy.next.

### Python (optimized)

```python
def remove_nth_from_end(head: Optional["ListNode"], n: int) -> Optional["ListNode"]:
    dummy = ListNode(0, head)
    first = second = dummy
    for _ in range(n + 1):
        first = first.next
    while first:
        first = first.next
        second = second.next
    second.next = second.next.next
    return dummy.next
```

### Golang (optimized)

```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	first, second := dummy, dummy
	for i := 0; i <= n; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}
```

- **Time:** O(L) one pass  
- **Space:** O(1)

---

## 4. Walkthrough

head = [1,2,3,4,5], n = 2. dummy→1→2→3→4→5. first moves to 3 (n+1=3 steps). Then first and second advance until first is past 5; second at 3. Remove 4: 3.next = 5. Return dummy.next = 1.

---

## 5. Edge Cases

1. **Remove head (n = length):** first reaches nil after n+1 steps; second still at dummy; remove head.  
2. **Single node, n=1:** return nil.  
3. **n = 1 (remove last):** second ends at second-to-last; remove last.  
4. **Dummy node:** Handles head removal without special case.

---

## 6. Pitfalls

1. **Off-by-one:** Advance first n+1 times so second lands one before the node to delete.  
2. **Not using dummy:** If n = length, we need to remove head; dummy simplifies.  
3. **Null pointer:** Guaranteed n <= size; second.next and second.next.next are valid when we delete.

---

## 7. Follow-ups

1. **Remove from start?** Trivial; remove head.  
2. **Remove every n-th from end in one pass?** Similar two-pointer with different gap.  
3. **Return the removed node?** Save second.next before deleting.

---

## 8. When to Use

- **“K-th from end” in one pass:** Two pointers with gap k (or k+1 for deletion).  
- **Pattern:** Dummy node + lead/lag pointers; same idea for “middle of list”.
