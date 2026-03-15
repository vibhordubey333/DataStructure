# Linked List Cycle — LeetCode 141 (E)

## 1. What is the problem?

**LeetCode 141 — Linked List Cycle**

Given the head of a linked list, determine if the list has a cycle (a node’s next points to an earlier node). Return true if cycle exists, else false.

**Example 1**

- **Input:** head = [3,2,0,-4], pos = 1 (cycle: tail connects to node 1)
- **Output:** true

**Example 2**

- **Input:** head = [1,2], pos = 0
- **Output:** true

---

## 2. Brute Force Solution

Use a set of node addresses (id(node) in Python, pointer in Go). Traverse; if we see a node already in the set, return true; else add and continue. If we reach nil, return false. Correct and O(n) time and space.

### Python (brute force)

```python
def has_cycle_brute(head: Optional["ListNode"]) -> bool:
    seen = set()
    while head:
        if id(head) in seen:
            return True
        seen.add(id(head))
        head = head.next
    return False
```

### Golang (brute force)

```go
func hasCycleBrute(head *ListNode) bool {
	seen := make(map[*ListNode]bool)
	for head != nil {
		if seen[head] {
			return true
		}
		seen[head] = true
		head = head.Next
	}
	return false
}
```

- **Time:** O(n)  
- **Space:** O(n)

---

## 3. Optimized Solution (Floyd’s Cycle Detection)

**Idea:** Slow pointer moves 1 step, fast moves 2 steps. If there is a cycle, they eventually meet. If fast reaches nil, no cycle. O(1) extra space.

**Steps:**

1. slow, fast = head, head.  
2. While fast and fast.next: slow = slow.next; fast = fast.next.next; if slow == fast: return True.  
3. Return False.

### Python (optimized)

```python
def has_cycle(head: Optional["ListNode"]) -> bool:
    slow = fast = head
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
        if slow == fast:
            return True
    return False
```

### Golang (optimized)

```go
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
```

- **Time:** O(n)  
- **Space:** O(1)

---

## 4. Walkthrough

List 3→2→0→-4→(back to 2). slow=3, fast=3 → slow=2, fast=0 → slow=0, fast=2 → slow=-4, fast=-4. Meet → true.

---

## 5. Edge Cases

1. **Empty list:** head=nil → false.  
2. **Single node, no cycle:** false.  
3. **Single node, self-cycle:** true.  
4. **Cycle at head:** slow and fast meet after some steps.  
5. **Check fast.Next before fast.Next.Next:** Avoid nil pointer.

---

## 6. Pitfalls

1. **Initial meet:** Start with slow=fast=head; first iteration they move then compare, so we don’t return true immediately.  
2. **Fast pointer:** Check both fast and fast.next so fast.next.next is safe.  
3. **Modifying list:** Don’t modify; just traverse.

---

## 7. Follow-ups

1. **Return cycle start node (LC 142)?** After meet, put one pointer at head, advance both by 1 until they meet; that’s the start.  
2. **Cycle length?** After meet, keep one fixed, advance the other until meet again; count steps.  
3. **Destroy cycle?** Find start, then traverse to node before start and set next = nil.

---

## 8. When to Use

- **Cycle detection:** Floyd’s slow/fast (tortoise and hare).  
- **Pattern:** O(1) space; same idea finds duplicate in array (if values in 1..n-1 and one duplicate).
