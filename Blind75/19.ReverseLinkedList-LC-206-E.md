# Reverse Linked List — LeetCode 206 (E)

## 1. What is the problem?

**LeetCode 206 — Reverse Linked List**

Given the head of a singly linked list, reverse the list and return the new head.

**Example 1**

- **Input:** head = [1, 2, 3, 4, 5]
- **Output:** [5, 4, 3, 2, 1]

**Example 2**

- **Input:** head = [1, 2]
- **Output:** [2, 1]

---

## 2. Brute Force Solution

Collect all node values into an array, reverse the array, then rebuild the list (or create new nodes with reversed values). Correct but uses O(n) extra space. “In-place” reversal is preferred.

### Python (brute — rebuild)

```python
def reverse_list_brute(head: Optional["ListNode"]) -> Optional["ListNode"]:
    vals = []
    while head:
        vals.append(head.val)
        head = head.next
    vals.reverse()
    dummy = ListNode(0)
    p = dummy
    for v in vals:
        p.next = ListNode(v)
        p = p.next
    return dummy.next
```

### Golang (brute)

```go
func reverseListBrute(head *ListNode) *ListNode {
	var vals []int
	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}
	for i, j := 0, len(vals)-1; i < j; i, j = i+1, j-1 {
		vals[i], vals[j] = vals[j], vals[i]
	}
	dummy := &ListNode{}
	p := dummy
	for _, v := range vals {
		p.Next = &ListNode{Val: v}
		p = p.Next
	}
	return dummy.Next
}
```

- **Time:** O(n)  
- **Space:** O(n)

---

## 3. Optimized Solution (Iterative In-Place)

**Idea:** Maintain a “reversed” list that we build by detaching the front node and prepending it. prev = nil; while head: next = head.next, head.next = prev, prev = head, head = next. Return prev.

**Steps:**

1. prev = nil.  
2. While head != nil: next = head.Next; head.Next = prev; prev = head; head = next.  
3. Return prev.

### Python (optimized)

```python
def reverse_list(head: Optional["ListNode"]) -> Optional["ListNode"]:
    prev = None
    while head:
        nxt = head.next
        head.next = prev
        prev = head
        head = nxt
    return prev
```

### Golang (optimized)

```go
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
```

- **Time:** O(n)  
- **Space:** O(1)

---

## 4. Walkthrough

head = 1→2→3→nil. Step 1: prev=1→nil, head=2→3. Step 2: prev=2→1→nil, head=3. Step 3: prev=3→2→1→nil, head=nil. Return prev = 3.

---

## 5. Edge Cases

1. **Empty list:** head=nil → return nil.  
2. **Single node:** return that node.  
3. **Two nodes:** swap and return new head.  
4. **Don’t lose reference:** Save head.next before changing head.next.

---

## 6. Pitfalls

1. **Losing the rest:** Always save next = head.next before head.next = prev.  
2. **Returning head:** After loop, head is nil; return prev.  
3. **Cycle:** Problem assumes no cycle; otherwise need to break cycle.

---

## 7. Follow-ups

1. **Reverse between positions (LC 92)?** Reverse the sublist then reconnect.  
2. **Reverse in groups of k (LC 25)?** Reverse each k block and link.  
3. **Recursive solution?** reverseList(head.next); head.next.next = head; head.next = nil; return newHead (from base case).

---

## 8. When to Use

- **Reverse a linked list:** Iterative three-pointer (prev, curr, next) or recursive.  
- **Pattern:** Fundamental; used in “palindrome list”, “reorder list”, etc.
