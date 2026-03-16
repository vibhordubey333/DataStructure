# Lowest Common Ancestor of a Binary Search Tree — LeetCode 235 (E)

## 1. What is the problem?

**LeetCode 235 — Lowest Common Ancestor of a Binary Search Tree**

Given a BST and two nodes p and q, return the lowest common ancestor (LCA). LCA is the lowest node that has both p and q as descendants.

**Example 1**

- **Input:** root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
- **Output:** 6
- **Explanation:** LCA of 2 and 8 is 6.

**Example 2**

- **Input:** root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
- **Output:** 2

---

## 2. Brute Force / Standard Solution

**BST property:** If p.val < root.val and q.val < root.val, LCA in left. If both > root.val, LCA in right. Otherwise root is LCA.

### Python

```python
def lowest_common_ancestor(root: "TreeNode", p: "TreeNode", q: "TreeNode") -> "TreeNode":
    if p.val < root.val and q.val < root.val:
        return lowest_common_ancestor(root.left, p, q)
    if p.val > root.val and q.val > root.val:
        return lowest_common_ancestor(root.right, p, q)
    return root
```

### Golang (iterative)

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}
```

- **Time:** O(h)  
- **Space:** O(1) iterative, O(h) recursive

---

## 4. Walkthrough

p=2, q=8. Root 6: 2<6 and 8>6 → root is LCA. p=2, q=4: both <6 → left. Root 2: 2<=2, 4>2 → root 2 is LCA.

---

## 5. Edge Cases

1. **p or q is root:** Root is LCA.  
2. **p is ancestor of q:** p is LCA.  
3. **Same node:** Return that node.  
4. **BST property:** Assumed.

---

## 6. Pitfalls

1. **BST vs binary tree:** LCA in binary tree (LC 236) needs different approach (postorder).  
2. **Iterative:** No stack; just traverse.  
3. **Assume p,q exist:** Per problem.

---

## 7. Follow-ups

1. **LCA in general binary tree?** Postorder; return node when both found.  
2. **With parent pointer?** Move to same depth, then go up.  
3. **Multiple queries?** Preprocess with binary lifting.

---

## 8. When to Use

- **LCA in BST:** Use BST property; go left/right until split.  
- **Pattern:** Simpler than general tree LCA.
