# Construct Binary Tree from Preorder and Inorder — LeetCode 105 (M)

## 1. What is the problem?

**LeetCode 105 — Construct Binary Tree from Preorder and Inorder Traversal**

Given two integer arrays preorder and inorder, construct and return the binary tree. Assume no duplicates.

**Example 1**

- **Input:** preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
- **Output:** [3,9,20,null,null,15,7]

**Example 2**

- **Input:** preorder = [-1], inorder = [-1]
- **Output:** [-1]

---

## 2. Brute Force / Standard Solution

**Idea:** preorder[0] is root. Find root in inorder; left of it is left subtree, right is right subtree. Recurse with preorder[1:1+left_size], inorder[:idx] and preorder[1+left_size:], inorder[idx+1:].

### Python

```python
def build_tree(preorder: list[int], inorder: list[int]) -> Optional["TreeNode"]:
    idx = {v: i for i, v in enumerate(inorder)}  # value -> inorder index

    def build(pl, pr, il, ir):
        if pl > pr:
            return None
        root_val = preorder[pl]
        root = TreeNode(root_val)
        i = idx[root_val]
        left_size = i - il
        root.left = build(pl + 1, pl + left_size, il, i - 1)
        root.right = build(pl + left_size + 1, pr, i + 1, ir)
        return root

    return build(0, len(preorder) - 1, 0, len(inorder) - 1)
```

### Golang

```go
func buildTree(preorder, inorder []int) *TreeNode {
	idx := make(map[int]int)
	for i, v := range inorder {
		idx[v] = i
	}
	var build func(pl, pr, il, ir int) *TreeNode
	build = func(pl, pr, il, ir int) *TreeNode {
		if pl > pr {
			return nil
		}
		rootVal := preorder[pl]
		i := idx[rootVal]
		leftSize := i - il
		return &TreeNode{
			Val:   rootVal,
			Left:  build(pl+1, pl+leftSize, il, i-1),
			Right: build(pl+leftSize+1, pr, i+1, ir),
		}
	}
	return build(0, len(preorder)-1, 0, len(inorder)-1)
}
```

- **Time:** O(n)  
- **Space:** O(n) for map and recursion

---

## 4. Walkthrough

pre=[3,9,20,15,7], in=[9,3,15,20,7]. Root=3, idx=1. Left: pre[1:2], in[0:1] → 9. Right: pre[2:5], in[2:5] → 20,15,7. Recurse.

---

## 5. Edge Cases

1. **Empty:** return None.  
2. **Single node:** Both have one element.  
3. **Skewed:** Left or right empty.  
4. **No duplicates:** Required for unique index.

---

## 6. Pitfalls

1. **Index bounds:** pl+left_size, pr, il, ir.  
2. **left_size = i - il:** Number of nodes in left subtree.  
3. **Postorder + inorder:** Similar; root is last in postorder.

---

## 7. Follow-ups

1. **Preorder + Postorder?** Ambiguous; multiple trees possible.  
2. **With duplicates?** Not possible to uniquely determine.  
3. **Iterative?** Use stack.

---

## 8. When to Use

- **Build tree from traversals:** Root from pre/post; split inorder by root.  
- **Pattern:** Same for "Construct from Inorder and Postorder" (LC 106).
