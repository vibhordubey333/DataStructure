# Binary Tree Level Order Traversal — LeetCode 102 (M)

## 1. What is the problem?

**LeetCode 102 — Binary Tree Level Order Traversal**

Given the root of a binary tree, return the level-order traversal of its nodes' values (left to right, level by level).

**Example 1**

- **Input:** root = [3,9,20,null,null,15,7]
- **Output:** [[3],[9,20],[15,7]]

**Example 2**

- **Input:** root = [1]
- **Output:** [[1]]

---

## 2. Brute Force / Standard Solution

BFS with a queue. Process level by level: for each level, take current queue size, poll that many nodes, collect values, add non-nil children. Correct and optimal.

### Python

```python
from collections import deque
def level_order(root: Optional["TreeNode"]) -> list[list[int]]:
    if not root:
        return []
    q = deque([root])
    out = []
    while q:
        level = []
        for _ in range(len(q)):
            node = q.popleft()
            level.append(node.val)
            if node.left:
                q.append(node.left)
            if node.right:
                q.append(node.right)
        out.append(level)
    return out
```

### Golang

```go
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	var out [][]int
	for len(q) > 0 {
		var level []int
		next := []*TreeNode{}
		for _, node := range q {
			level = append(level, node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		out = append(out, level)
		q = next
	}
	return out
}
```

- **Time:** O(n)  
- **Space:** O(w) queue

---

## 3. Optimized Solution

BFS level-order is the standard and optimal. DFS with level index (append to result[level]) is also O(n).

---

## 4. Walkthrough

Root 3 → level [3]; then 9, 20 → [9,20]; then 15, 7 → [15,7]. Result [[3],[9,20],[15,7]].

---

## 5. Edge Cases

1. **Empty tree:** [].  
2. **Single node:** [[1]].  
3. **Skewed:** One node per level.

---

## 6. Pitfalls

1. **Level boundary:** Use for _ in range(len(q)) so we process exactly one level per iteration.  
2. **Nil children:** Don’t push nil.

---

## 7. Follow-ups

1. **Level order II (reverse levels)?** BFS then reverse result or use stack.  
2. **Zigzag level order?** Reverse every other level.  
3. **Bottom-up?** Reverse the output list.

---

## 8. When to Use

- **Level-by-level traversal:** BFS with level-sized batches.  
- **Pattern:** Same in “average of levels”, “largest value per level”, etc.
