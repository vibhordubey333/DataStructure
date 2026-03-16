# Graph Valid Tree — LeetCode 261 (M, Premium)

## 1. What is the problem?

**LeetCode 261 — Graph Valid Tree**

Given n nodes labeled 0 to n-1 and a list of undirected edges, write a function to check whether these edges make up a valid tree. A tree is a connected graph with no cycles.

**Example 1**

- **Input:** n = 5, edges = [[0,1],[0,2],[0,3],[1,4]]
- **Output:** true

**Example 2**

- **Input:** n = 5, edges = [[0,1],[1,2],[2,3],[1,3],[1,4]]
- **Output:** false
- **Explanation:** Cycle 1-2-3-1.

---

## 2. Brute Force / Standard Solution

**Valid tree:** (1) Exactly n-1 edges. (2) Connected (one component). Use Union-Find: if edge connects already-connected nodes, cycle → false. Finally check one component.

### Python (Union-Find)

```python
def valid_tree(n: int, edges: list[list[int]]) -> bool:
    if len(edges) != n - 1:
        return False
    parent = list(range(n))

    def find(x):
        if parent[x] != x:
            parent[x] = find(parent[x])
        return parent[x]

    for u, v in edges:
        pu, pv = find(u), find(v)
        if pu == pv:
            return False  # cycle
        parent[pu] = pv
    return True
```

### Golang

```go
func validTree(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	for _, e := range edges {
		pu, pv := find(e[0]), find(e[1])
		if pu == pv {
			return false
		}
		parent[pu] = pv
	}
	return true
}
```

- **Time:** O(n + E)  
- **Space:** O(n)

---

## 4. Walkthrough

n=5, edges=[[0,1],[0,2],[0,3],[1,4]]. 4 edges = n-1 ✓. Union all: no cycle. One component ✓. True.

---

## 5. Edge Cases

1. **n=1, edges=[]:** true (single node).  
2. **n=2, edges=[[0,1]]:** true.  
3. **Extra edge creates cycle:** false.  
4. **Too few edges:** Disconnected, false.

---

## 6. Pitfalls

1. **n-1 edges:** Necessary for tree. With n-1 edges and no cycle, must be connected.  
2. **Cycle detection:** Same component before union.  
3. **DFS alternative:** Build graph, DFS with parent; back edge = cycle.

---

## 7. Follow-ups

1. **Return MST?** Different.  
2. **Directed?** Different (DAG with one source).  
3. **Find cycle?** Return edges in cycle.

---

## 8. When to Use

- **Check if graph is tree:** n-1 edges + no cycle (Union-Find).  
- **Pattern:** Union-Find for cycle detection; same as "Number of Connected Components".
