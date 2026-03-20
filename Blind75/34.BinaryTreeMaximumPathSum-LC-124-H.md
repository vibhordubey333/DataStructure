# Binary Tree Maximum Path Sum — LeetCode 124 (H)

## 1. What is the problem?

**LeetCode 124 — Binary Tree Maximum Path Sum**

A path is a sequence of nodes where each pair of adjacent nodes has an edge. Return the maximum path sum. A path can start and end at any node (must include at least one node).

**Example 1**

- **Input:** root = [1,2,3]
- **Output:** 6
- **Explanation:** Path 2-1-3 has sum 6.

**Example 2**

- **Input:** root = [-10,9,20,null,null,15,7]
- **Output:** 42
- **Explanation:** Path 15-20-7 has sum 42.

---

## 2. Brute Force / Standard Solution

**Idea:** For each node, max path through it = node.val + max(0, left_gain) + max(0, right_gain). "Gain" = max sum of path going down from node (to one child). Recurse: return max_gain (for parent), and track global max.

### Python

```python
def max_path_sum(root: Optional["TreeNode"]) -> int:
    best = float('-inf')

    def gain(node):
        nonlocal best
        if not node:
            return 0
        left = max(0, gain(node.left))
        right = max(0, gain(node.right))
        best = max(best, node.val + left + right)
        return node.val + max(left, right)

    gain(root)
    return best
```

### Golang

```go
func maxPathSum(root *TreeNode) int {
	best := math.MinInt64
	var gain func(*TreeNode) int
	gain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(0, gain(node.Left))
		right := max(0, gain(node.Right))
		if node.Val+left+right > best {
			best = node.Val + left + right
		}
		return node.Val + max(left, right)
	}
	gain(root)
	return best
}
```

- **Time:** O(n)  
- **Space:** O(h)

---

## 4. Walkthrough

At node 20: left=15, right=7. Path through 20 = 15+20+7=42. Return 20+15=35 for parent. Global best=42.

---

## 5. Edge Cases

1. **Single node:** Return node.val.  
2. **All negative:** Return least negative (single node).  
3. **Large values:** Use int.  
4. **Skewed:** One path through root.

---

## 6. Pitfalls

1. **max(0, gain):** Avoid negative gains; we can "stop" at a node.  
2. **Path through node:** Must be at most one turn (left-up-right).  
3. **Return value:** Single path down (for parent to extend).

---

## 7. Follow-ups

1. **Return the path?** Track nodes during recursion.  
2. **At most k turns?** More complex DP.  
3. **Diameter (max nodes)?** Same structure; count nodes.

---

## 8. When to Use

- **Max path through tree:** Gain per node; path = node + left_gain + right_gain.  
- **Pattern:** Same as "Diameter of Binary Tree" (sum instead of count).
