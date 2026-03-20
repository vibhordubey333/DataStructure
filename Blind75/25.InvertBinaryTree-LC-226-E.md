# Invert Binary Tree — LeetCode 226 (E)

## 1. What is the problem?

**LeetCode 226 — Invert Binary Tree**

Given the root of a binary tree, invert the tree (swap left and right of every node) and return the root.

**Example 1**

- **Input:** root = [4,2,7,1,3,6,9]
- **Output:** [4,7,2,9,6,3,1]

**Example 2**

- **Input:** root = [2,1,3]
- **Output:** [2,3,1]

---

## 2. Brute Force / Direct Solution

Recurse: if root is nil return nil. Swap root.left and root.right (after recursing or before). Invert(left), invert(right), then swap; or swap then invert both. Correct and optimal.

### Python

```python
def invert_tree(root: Optional["TreeNode"]) -> Optional["TreeNode"]:
    if not root:
        return None
    root.left, root.right = root.right, root.left
    invert_tree(root.left)
    invert_tree(root.right)
    return root
```

### Golang

```go
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
```

- **Time:** O(n)  
- **Space:** O(h) stack

---

## 3. Optimized Solution

Same; no asymptotically better approach. Iterative: use a stack or queue, for each node swap left/right and push children. O(n) time, O(w) space for queue.

### Python (iterative BFS)

```python
def invert_tree(root: Optional["TreeNode"]) -> Optional["TreeNode"]:
    if not root:
        return None
    q = [root]
    while q:
        node = q.pop(0)
        node.left, node.right = node.right, node.left
        if node.left:
            q.append(node.left)
        if node.right:
            q.append(node.right)
    return root
```

---

## 4. Walkthrough

Root 4, left 2, right 7. Swap → left=7, right=2. Invert(7): 7’s children 6,9 → swap to 9,6. Invert(2): 1,3 → 3,1. Result: 4 with left 7 (9,6), right 2 (3,1).

---

## 5. Edge Cases

1. **Empty tree:** return nil.  
2. **Single node:** no change.  
3. **Only one child:** swap makes it the other side.

---

## 6. Pitfalls

1. **Order of swap and recurse:** Swap first then recurse (or recurse first then swap); both correct.  
2. **Losing reference:** Use simultaneous swap so we don’t lose a child.  
3. **Return value:** Return root for chaining; problem asks for new root (same root in place).

---

## 7. Follow-ups

1. **Mirror image (symmetric)?** Invert and compare with original (or recursive left vs right mirror check).  
2. **Clone and invert?** Create new tree with swapped structure.  
3. **Level-order print after invert?** BFS after inverting.

---

## 8. When to Use

- **Swap left/right everywhere:** One recursive or iterative pass.  
- **Pattern:** Simple tree mutation; same idea in “mirror” problems.
