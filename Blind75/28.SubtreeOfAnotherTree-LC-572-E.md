# Subtree of Another Tree — LeetCode 572 (E)

## 1. What is the problem?

**LeetCode 572 — Subtree of Another Tree**

Given the roots of two binary trees root and subRoot, return true if there is a subtree of root with the same structure and node values as subRoot.

**Example 1**

- **Input:** root = [3,4,5,1,2], subRoot = [4,1,2]
- **Output:** true

**Example 2**

- **Input:** root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
- **Output:** false

---

## 2. Brute Force / Standard Solution

For each node in root, check if subtree rooted at that node equals subRoot. Use helper `same_tree(a, b)` and `is_subtree(root, sub)`.

### Python

```python
def is_subtree(root: Optional["TreeNode"], sub_root: Optional["TreeNode"]) -> bool:
    def same(a, b):
        if not a and not b:
            return True
        if not a or not b or a.val != b.val:
            return False
        return same(a.left, b.left) and same(a.right, b.right)

    if not root:
        return not sub_root
    if same(root, sub_root):
        return True
    return is_subtree(root.left, sub_root) or is_subtree(root.right, sub_root)
```

### Golang

```go
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return subRoot == nil
	}
	if same(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func same(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil || a.Val != b.Val {
		return false
	}
	return same(a.Left, b.Left) && same(a.Right, b.Right)
}
```

- **Time:** O(m * n) — m nodes in root, n in subRoot  
- **Space:** O(h)

---

## 4. Walkthrough

At root 3: same(3,4)? No. Recurse left (4): same(4,4)? Yes. Return true.

---

## 5. Edge Cases

1. **subRoot empty:** return true (empty is subtree).  
2. **root empty, subRoot non-empty:** false.  
3. **subRoot single node:** Check any matching node.  
4. **Identical trees:** true.

---

## 6. Pitfalls

1. **Structure + values:** Same tree means same structure and values.  
2. **Subtree:** Must be contiguous from some node down.  
3. **Merkle hash:** O(m+n) with tree hashing; more complex.

---

## 7. Follow-ups

1. **Count occurrences?** Count when same() returns true.  
2. **Overlap allowed?** Different problem.  
3. **KMP on serialization?** Serialize both, find substring.

---

## 8. When to Use

- **Find subtree match:** Same-tree check at each node.  
- **Pattern:** Tree recursion; same as "Same Tree" + "Find node".
