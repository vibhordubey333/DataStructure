# Maximum Depth of Binary Tree — LeetCode 104 (E)

## 1. What is the problem?

**LeetCode 104 — Maximum Depth of Binary Tree**

Given the root of a binary tree, return its maximum depth (number of nodes along the longest path from root to leaf).

**Example 1**

- **Input:** root = [3,9,20,null,null,15,7]
- **Output:** 3

**Example 2**

- **Input:** root = [1,null,2]
- **Output:** 2

---

## 2. Brute Force Solution

Recurse: depth(root) = 0 if root is nil, else 1 + max(depth(root.left), depth(root.right)). This is already the standard and optimal approach; “brute” would be same with no memoization (no need for memo here).

### Python (recursive)

```python
def max_depth(root: Optional["TreeNode"]) -> int:
    if not root:
        return 0
    return 1 + max(max_depth(root.left), max_depth(root.right))
```

### Golang (recursive)

```go
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return 1 + left
	}
	return 1 + right
}
```

- **Time:** O(n)  
- **Space:** O(h) stack, h = height

---

## 3. Optimized Solution (BFS for level count)

**Idea:** BFS level-order: count levels. Or iterative DFS with stack storing (node, depth); track max depth. Both O(n).

**Steps (BFS):**

1. If not root return 0. Queue = [root], depth = 0.  
2. While queue: depth += 1; for each node in current level, add non-nil children to next level.  
3. Return depth.

### Python (BFS)

```python
from collections import deque
def max_depth(root: Optional["TreeNode"]) -> int:
    if not root:
        return 0
    q = deque([root])
    depth = 0
    while q:
        depth += 1
        for _ in range(len(q)):
            node = q.popleft()
            if node.left:
                q.append(node.left)
            if node.right:
                q.append(node.right)
    return depth
```

### Golang (BFS)

```go
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	depth := 0
	for len(q) > 0 {
		depth++
		next := []*TreeNode{}
		for _, node := range q {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		q = next
	}
	return depth
}
```

- **Time:** O(n)  
- **Space:** O(w) queue, w = max width

---

## 4. Walkthrough

Root 3, left 9, right 20 with children 15,7. Recursive: depth(3)=1+max(depth(9),depth(20))=1+max(1,1+max(1,1))=1+2=3. Output 3.

---

## 5. Edge Cases

1. **Empty tree:** 0.  
2. **Single node:** 1.  
3. **Skewed (linked list):** depth = n.  
4. **Balanced:** depth ≈ log n.

---

## 6. Pitfalls

1. **Depth vs number of edges:** Problem says “number of nodes”; root = 1. If “edges” then return 0 for single node.  
2. **Null root:** Return 0.  
3. **Stack overflow:** For very deep tree, use BFS or iterative DFS.

---

## 7. Follow-ups

1. **Minimum depth (LC 111)?** Only count to a leaf; if one child is nil, follow the other.  
2. **Balanced check (LC 110)?** For each node, |depth(left)-depth(right)| <= 1 and both subtrees balanced.  
3. **Diameter (LC 543)?** Max (left_depth + right_depth) over nodes.

---

## 8. When to Use

- **Tree depth / height:** Recursive 1 + max(left, right) or BFS level count.  
- **Pattern:** Base for balanced check, diameter, and many tree problems.
