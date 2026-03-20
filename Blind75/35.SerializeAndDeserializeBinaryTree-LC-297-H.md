# Serialize and Deserialize Binary Tree — LeetCode 297 (H)

## 1. What is the problem?

**LeetCode 297 — Serialize and Deserialize Binary Tree**

Design an algorithm to serialize and deserialize a binary tree. Serialization: convert tree to string. Deserialization: convert string back to tree. No restriction on format.

**Example 1**

- **Input:** root = [1,2,3,null,null,4,5]
- **Output:** [1,2,3,null,null,4,5]
- **Explanation:** Your serialized string could be "1,2,3,null,null,4,5" (preorder with nulls).

**Example 2**

- **Input:** root = []
- **Output:** []

---

## 2. Brute Force / Standard Solution

**Preorder with null markers:** Serialize: preorder, emit "null" for None. Deserialize: consume tokens, build tree recursively.

### Python

```python
class Codec:
    def serialize(self, root: Optional["TreeNode"]) -> str:
        def pre(node):
            if not node:
                return ["null"]
            return [str(node.val)] + pre(node.left) + pre(node.right)
        return ",".join(pre(root))

    def deserialize(self, data: str) -> Optional["TreeNode"]:
        tokens = data.split(",")
        i = [0]
        def build():
            if tokens[i[0]] == "null":
                i[0] += 1
                return None
            node = TreeNode(int(tokens[i[0]]))
            i[0] += 1
            node.left = build()
            node.right = build()
            return node
        return build() if tokens and tokens[0] != "null" else None
```

### Golang

```go
type Codec struct{}

func Constructor() Codec { return Codec{} }

func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "null"
	}
	return strconv.Itoa(root.Val) + "," + c.serialize(root.Left) + "," + c.serialize(root.Right)
}

func (c *Codec) deserialize(data string) *TreeNode {
	tokens := strings.Split(data, ",")
	i := 0
	var build func() *TreeNode
	build = func() *TreeNode {
		if tokens[i] == "null" {
			i++
			return nil
		}
		val, _ := strconv.Atoi(tokens[i])
		i++
		node := &TreeNode{Val: val}
		node.Left = build()
		node.Right = build()
		return node
	}
	if len(tokens) == 0 || tokens[0] == "null" {
		return nil
	}
	return build()
}
```

- **Time:** O(n) serialize and deserialize  
- **Space:** O(n) for string and recursion

---

## 4. Walkthrough

Serialize: 1,2,null,null,3,4,null,null,5,null,null. Deserialize: read 1, build left (2,null,null), right (3,4,null,null,5,null,null).

---

## 5. Edge Cases

1. **Empty tree:** "null" or "".  
2. **Single node:** "1,null,null".  
3. **Skewed tree:** Many nulls.  
4. **Large values:** Handle in string.

---

## 6. Pitfalls

1. **Index in deserialize:** Use mutable ref (list, pointer) to advance.  
2. **Split on comma:** Ensure no comma in numbers.  
3. **BFS alternative:** Level-order with nulls; different format.

---

## 7. Follow-ups

1. **BST?** Inorder gives sorted; can compress.  
2. **N-ary tree?** Store child count or use brackets.  
3. **Compact binary?** Bit-packed format.

---

## 8. When to Use

- **Tree ↔ string:** Preorder + null markers is simple and robust.  
- **Pattern:** Same idea for "Construct from preorder" style.
