# Number of Connected Components in an Undirected Graph — LeetCode 323 (M, Premium)

## 1. What is the problem?

**LeetCode 323 — Number of Connected Components in an Undirected Graph**

Given n nodes labeled 0 to n-1 and a list of undirected edges, write a function to find the number of connected components.

**Example 1**

- **Input:** n = 5, edges = [[0,1],[1,2],[3,4]]
- **Output:** 2

**Example 2**

- **Input:** n = 5, edges = [[0,1],[1,2],[2,3],[3,4]]
- **Output:** 1

---

## 2. Brute Force / Standard Solution

**Union-Find (Disjoint Set Union):** For each edge, union the two nodes. Count distinct roots. Or DFS/BFS from each unvisited node; count components.

### Python (Union-Find)

```python
def count_components(n: int, edges: list[list[int]]) -> int:
    parent = list(range(n))

    def find(x):
        if parent[x] != x:
            parent[x] = find(parent[x])
        return parent[x]

    def union(x, y):
        px, py = find(x), find(y)
        if px != py:
            parent[px] = py

    for u, v in edges:
        union(u, v)
    return sum(1 for i in range(n) if find(i) == i)
```

### Golang

```go
func countComponents(n int, edges [][]int) int {
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
	union := func(x, y int) {
		px, py := find(x), find(y)
		if px != py {
			parent[px] = py
		}
	}
	for _, e := range edges {
		union(e[0], e[1])
	}
	count := 0
	for i := 0; i < n; i++ {
		if find(i) == i {
			count++
		}
	}
	return count
}
```

- **Time:** O(n + E * α(n)) ≈ O(n + E)  
- **Space:** O(n)

---

## 4. Walkthrough

n=5, edges=[[0,1],[1,2],[3,4]]. Union 0-1, 1-2 → {0,1,2}. Union 3-4 → {3,4}. Roots: 2, 4. Count=2.

---

## 5. Edge Cases

1. **No edges:** n components.  
2. **Single node:** 1 component.  
3. **All connected:** 1 component.  
4. **Empty edges:** n components.

---

## 6. Pitfalls

1. **Path compression:** find with parent[x]=find(parent[x]).  
2. **Union by rank:** Optional optimization.  
3. **Undirected:** Each edge connects both ways.

---

## 7. Follow-ups

1. **Return components?** Group by root.  
2. **Add/remove edges?** Dynamic connectivity.  
3. **Directed graph?** Strongly connected components (Kosaraju/Tarjan).

---

## 8. When to Use

- **Count connected components:** Union-Find or DFS/BFS.  
- **Pattern:** Same as "Graph Valid Tree", "Number of Islands".
