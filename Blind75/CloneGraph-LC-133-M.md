# Clone Graph — LeetCode 133 (M)

## 1. What is the problem?

**LeetCode 133 — Clone Graph**

Given a reference to a node in a connected undirected graph, return a deep copy of the graph. Each node has val and list of neighbors.

**Example 1**

- **Input:** adjList = [[2,4],[1,3],[2,4],[1,3]]
- **Output:** Deep copy of the graph

**Example 2**

- **Input:** adjList = [[]]
- **Output:** [[]]

---

## 2. Brute Force Solution

BFS/DFS: traverse graph, for each node create a copy and map old→new. When processing neighbors, create copy if not seen, add to neighbors. Same as optimal approach.

### Python (BFS)

```python
def clone_graph(node: "Node") -> "Node":
    if not node:
        return None
    from collections import deque
    clone = {node: Node(node.val)}
    q = deque([node])
    while q:
        cur = q.popleft()
        for nb in cur.neighbors:
            if nb not in clone:
                clone[nb] = Node(nb.val)
                q.append(nb)
            clone[cur].neighbors.append(clone[nb])
    return clone[node]
```

### Golang (BFS)

```go
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	clone := map[*Node]*Node{node: {Val: node.Val}}
	q := []*Node{node}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, nb := range cur.Neighbors {
			if clone[nb] == nil {
				clone[nb] = &Node{Val: nb.Val}
				q = append(q, nb)
			}
			clone[cur].Neighbors = append(clone[cur].Neighbors, clone[nb])
		}
	}
	return clone[node]
}
```

- **Time:** O(V + E)  
- **Space:** O(V)

---

## 3. Optimized Solution

Same BFS/DFS with hash map. DFS recursive also works.

### Python (DFS)

```python
def clone_graph(node: "Node") -> "Node":
    if not node:
        return None
    clone = {}
    def dfs(n):
        if n in clone:
            return clone[n]
        c = Node(n.val)
        clone[n] = c
        for nb in n.neighbors:
            c.neighbors.append(dfs(nb))
        return c
    return dfs(node)
```

---

## 4. Walkthrough

Start at node 1. Clone 1, add to map. Process neighbors 2,4. Clone 2, 4; add to queue. Link clone[1].neighbors = [clone[2], clone[4]]. Process 2: clone 3, link. Process 4: 1,3 already cloned, link. Done.

---

## 5. Edge Cases

1. **Empty graph:** return nil.  
2. **Single node:** clone with empty neighbors.  
3. **Self-loop:** Handle in map (node in clone).  
4. **Disconnected:** Problem says connected; one node given.

---

## 6. Pitfalls

1. **Infinite loop:** Must check if node already cloned before recursing.  
2. **Same reference:** Each node must be new; don't reuse.  
3. **Node struct:** Val and Neighbors; Neighbors is list of *Node.

---

## 7. Follow-ups

1. **Clone linked list with random pointer?** Map old→new; two passes.  
2. **Clone tree?** Simpler; no cycle.  
3. **Serialize/deserialize graph?** BFS with node IDs.

---

## 8. When to Use

- **Deep copy of graph:** BFS/DFS + hash map old node → new node.  
- **Pattern:** Same in "copy list with random pointer", any structure with cycles.
