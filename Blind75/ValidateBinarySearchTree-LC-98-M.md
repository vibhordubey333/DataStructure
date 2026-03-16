# Validate Binary Search Tree — LeetCode 98 (M)

## 1. What is the problem?

**LeetCode 98 — Validate Binary Search Tree**

Given the root of a binary tree, determine if it is a valid BST. Left subtree < root < right subtree; all descendants must satisfy.

**Example 1**

- **Input:** root = [2,1,3]
- **Output:** true

**Example 2**

- **Input:** root = [5,1,4,null,null,3,6]
- **Output:** false
- **Explanation:** Root 5; right subtree has 4, but 3 < 5.

---

## 2. Brute Force / Standard Solution

**Idea:** Pass (lo, hi) bounds. For each node, lo < node.val < hi. Recurse left with (lo, node.val), right with (node.val, hi).

### Python

```python
def is_valid_bst(root: Optional["TreeNode"]) -> bool:
    def valid(node, lo, hi):
        if not node:
            return True
        if not (lo < node.val < hi):
            return False
        return valid(node.left, lo, node.val) and valid(node.right, node.val, hi)

    return valid(root, float('-inf'), float('inf'))
```

### Golang

```go
func isValidBST(root *TreeNode) bool {
	var valid func(*TreeNode, int64, int64) bool
	valid = func(node *TreeNode, lo, hi int64) bool {
		if node == nil {
			return true
		}
		v := int64(node.Val)
		if v <= lo || v >= hi {
			return false
		}
		return valid(node.Left, lo, v) && valid(node.Right, v, hi)
	}
	return valid(root, math.MinInt64, math.MaxInt64)
}
```

- **Time:** O(n)  
- **Space:** O(h)

---

## 4. Walkthrough

Root 5: (min, max). Left 1: (min, 5) ✓. Right 4: (5, max). But 4 < 5, so invalid. Actually 4 is in right subtree; 3 is left of 4. So 3 < 5 violates. Valid checks: 4 in (5, inf) ✓, 3 in (5, 4) ✗. Return false.

---

## 5. Edge Cases

1. **Empty:** true.  
2. **Single node:** true.  
3. **Duplicate values:** Usually invalid (strict < and >).  
4. **Large values:** Use int64 for bounds.

---

## 6. Pitfalls

1. **Not just parent:** Must satisfy all ancestors.  
2. **Use int64:** To handle node.val = 2^31-1 or -2^31.  
3. **Inorder alternative:** Inorder should be strictly increasing.

---

## 7. Follow-ups

1. **Recover BST (two swapped)?** Inorder; find two inversions.  
2. **Convert to BST?** Sort and build.  
3. **Inorder iterative?** Check prev < curr.

---

## 8. When to Use

- **Validate BST:** Pass bounds (lo, hi) down.  
- **Pattern:** Same bounds idea in "Construct BST" problems.
