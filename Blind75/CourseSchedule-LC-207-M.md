# Course Schedule — LeetCode 207 (M)

## 1. What is the problem?

**LeetCode 207 — Course Schedule**

There are `numCourses` courses labeled 0 to numCourses-1. `prerequisites[i] = [a, b]` means you must take b before a. Return true if you can finish all courses.

**Example 1**

- **Input:** numCourses = 2, prerequisites = [[1,0]]
- **Output:** true (take 0 then 1)

**Example 2**

- **Input:** numCourses = 2, prerequisites = [[1,0],[0,1]]
- **Output:** false (cycle)

---

## 2. Brute Force Solution

Build directed graph. Detect cycle (DFS with recursion stack / three colors). If cycle exists, return false. Correct but "brute" is same as optimal for cycle detection.

### Python (DFS cycle detection)

```python
def can_finish(num_courses: int, prerequisites: list[list[int]]) -> bool:
    graph = [[] for _ in range(num_courses)]
    for a, b in prerequisites:
        graph[a].append(b)
    state = [0] * num_courses  # 0=unvisited, 1=visiting, 2=done
    def has_cycle(v):
        if state[v] == 1:
            return True
        if state[v] == 2:
            return False
        state[v] = 1
        for u in graph[v]:
            if has_cycle(u):
                return True
        state[v] = 2
        return False
    for v in range(num_courses):
        if has_cycle(v):
            return False
    return True
```

### Golang (DFS)

```go
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for _, p := range prerequisites {
		graph[p[0]] = append(graph[p[0]], p[1])
	}
	state := make([]int, numCourses)
	var hasCycle func(int) bool
	hasCycle = func(v int) bool {
		if state[v] == 1 {
			return true
		}
		if state[v] == 2 {
			return false
		}
		state[v] = 1
		for _, u := range graph[v] {
			if hasCycle(u) {
				return true
			}
		}
		state[v] = 2
		return false
	}
	for v := 0; v < numCourses; v++ {
		if hasCycle(v) {
			return false
		}
	}
	return true
}
```

- **Time:** O(V + E)  
- **Space:** O(V)

---

## 3. Optimized Solution

Same: topological sort or cycle detection. Kahn's algorithm (BFS with in-degree) also works: repeatedly remove nodes with in-degree 0; if any remain, cycle exists.

---

## 4. Walkthrough

numCourses=2, pre=[[1,0]]. Graph: 1→0. DFS from 1: visit 1, then 0. No back edge. Return true.

pre=[[1,0],[0,1]]. 1→0, 0→1. Cycle. Return false.

---

## 5. Edge Cases

1. **No prerequisites:** true.  
2. **Single course:** true.  
3. **Self-loop (a,b) with a==b:** Invalid; problem may not have.  
4. **Disconnected components:** Check each; cycle in any → false.

---

## 6. Pitfalls

1. **Graph direction:** prerequisites [a,b] means b→a (b before a). So edge a→b in "depends on" terms: a depends on b, so edge from a to b.  
2. **Three states:** Unvisited, visiting (in stack), done. Back edge to visiting = cycle.  
3. **Kahn's:** Build in-degree; queue nodes with 0; decrement neighbors; if count removed < n, cycle.

---

## 7. Follow-ups

1. **Course Schedule II (return order)?** Topological sort; return order or empty if cycle.  
2. **Course Schedule III (duration + deadline)?** Greedy; different problem.  
3. **Parallel courses?** Level-by-level BFS; return max level.

---

## 8. When to Use

- **DAG / dependency order:** Cycle detection (DFS) or topological sort (Kahn).  
- **Pattern:** Prerequisites, build order, "can complete" → detect cycle in directed graph.
