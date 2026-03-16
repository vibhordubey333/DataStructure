# Kth Smallest Element in a BST — LeetCode 230 (M)

## 1. What is the problem?

**LeetCode 230 — Kth Smallest Element in a BST**

Given the root of a BST and an integer k, return the kth smallest value (1-indexed) in the BST.

**Example 1**

- **Input:** root = [3,1,4,null,2], k = 1
- **Output:** 1

**Example 2**

- **Input:** root = [5,3,6,2,4,null,null,1], k = 3
- **Output:** 3

---

## 2. Brute Force Solution

Inorder traversal gives sorted order. Collect all values, return kth. O(n) time and space.

### Python (inorder iterative — stop at k)

```python
def kth_smallest(root: Optional["TreeNode"], k: int) -> int:
    stack = []
    node = root
    while node or stack:
        while node:
            stack.append(node)
            node = node.left
        node = stack.pop()
        k -= 1
        if k == 0:
            return node.val
        node = node.right
    return -1
```

### Golang

```go
func kthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return node.Val
		}
		node = node.Right
	}
	return -1
}
```

- **Time:** O(h + k) — go left to smallest, then k steps  
- **Space:** O(h)

---

## 4. Walkthrough

Inorder: 1,2,3,4. k=1: return 1. k=3: return 3.

---

## 5. Edge Cases

1. **k=1:** Smallest (leftmost).  
2. **k=n:** Largest (rightmost).  
3. **Single node:** Return root.val.  
4. **1 <= k <= n:** Guaranteed.

---

## 6. Pitfalls

1. **1-indexed:** kth means k=1 is smallest.  
2. **Inorder:** Left-root-right gives sorted order in BST.  
3. **Recursive with global:** Alternatively use recursive inorder with counter.

---

## 7. Follow-ups

1. **BST is modified often?** Augment with subtree size; binary search.  
2. **Kth largest?** Reverse inorder (right-root-left) or (n-k+1)th smallest.  
3. **Multiple queries?** Build sorted array once.

---

## 8. When to Use

- **Kth smallest in BST:** Inorder traversal; stop at k.  
- **Pattern:** Inorder = sorted; same for "Validate BST".
