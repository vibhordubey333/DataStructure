# Same Tree — LeetCode 100 (E)

## 1. What is the problem?

**LeetCode 100 — Same Tree**

Given the roots of two binary trees p and q, return true if they are structurally identical and nodes have the same values.

**Example 1**

- **Input:** p = [1,2,3], q = [1,2,3]
- **Output:** true

**Example 2**

- **Input:** p = [1,2], q = [1,null,2]
- **Output:** false

---

## 2. Brute Force / Direct Solution

Recurse: both nil → true; one nil → false; values differ → false; else return is_same(p.left, q.left) and is_same(p.right, q.right). This is the standard and optimal.

### Python

```python
def is_same_tree(p: Optional["TreeNode"], q: Optional["TreeNode"]) -> bool:
    if p is None and q is None:
        return True
    if p is None or q is None:
        return False
    if p.val != q.val:
        return False
    return is_same_tree(p.left, q.left) and is_same_tree(p.right, q.right)
```

### Golang

```go
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
```

- **Time:** O(n)  
- **Space:** O(h) stack

---

## 3. Optimized Solution

Same recursion is already O(n). Iterative: use two queues (or one queue with pairs) and compare level by level. Still O(n).

---

## 4. Walkthrough

p = [1,2,3], q = [1,2,3]. Compare 1==1, then left (2,2) and right (3,3); all match → true.  
p = [1,2], q = [1,null,2]: 1==1; left: (2, nil) → one nil → false.

---

## 5. Edge Cases

1. **Both nil:** true.  
2. **One nil:** false.  
3. **Same structure, different val:** false.  
4. **Different structure:** false (e.g. one has left child, other has right).

---

## 6. Pitfalls

1. **Checking value before nil:** Check structure (both nil / one nil) first to avoid nil dereference.  
2. **Short-circuit:** and is correct; if left differs we don’t need to check right.

---

## 7. Follow-ups

1. **Subtree of another tree (LC 572)?** For each node of main tree, check if subtree rooted there equals given subtree.  
2. **Symmetric tree (LC 101)?** Same logic with (left, right) and (right, left).  
3. **Clone tree?** Same structure, create new nodes with same values.

---

## 8. When to Use

- **Compare two trees:** Recursive structure + value check.  
- **Pattern:** Base for subtree, symmetric, and serialization comparison.
