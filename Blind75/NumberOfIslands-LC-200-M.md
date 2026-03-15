# Number of Islands — LeetCode 200 (M)

## 1. What is the problem?

**LeetCode 200 — Number of Islands**

Given a 2D grid of '1's (land) and '0's (water), return the number of islands (connected 4-directional land).

**Example 1**

- **Input:** grid = [["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]
- **Output:** 1

**Example 2**

- **Input:** grid = [["1","1","0","0","0"],["1","1","0","0","0"],["0","0","1","0","0"],["0","0","0","1","1"]]
- **Output:** 3

---

## 2. Brute Force Solution

For each cell (i,j), if grid[i][j]=='1', do a DFS or BFS to mark entire island (e.g. set to '0' or use visited), increment count. Each '1' is visited once. Correct; the “brute” is just trying every cell and flooding — same as standard approach.

### Python (DFS)

```python
def num_islands(grid: list[list[str]]) -> int:
    if not grid or not grid[0]:
        return 0
    rows, cols = len(grid), len(grid[0])
    def dfs(r, c):
        if r < 0 or r >= rows or c < 0 or c >= cols or grid[r][c] != '1':
            return
        grid[r][c] = '0'
        for dr, dc in [(1,0),(-1,0),(0,1),(0,-1)]:
            dfs(r+dr, c+dc)
    count = 0
    for r in range(rows):
        for c in range(cols):
            if grid[r][c] == '1':
                count += 1
                dfs(r, c)
    return count
```

### Golang (DFS)

```go
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	var dfs func(int, int)
	dfs = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != '1' {
			return
		}
		grid[r][c] = '0'
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}
	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				count++
				dfs(r, c)
			}
		}
	}
	return count
}
```

- **Time:** O(rows * cols)  
- **Space:** O(rows * cols) recursion or queue

---

## 3. Optimized Solution

Same DFS/BFS is optimal. Alternative: Union-Find (disjoint set). For each '1', union with neighbors; number of distinct roots = number of islands. Same O(R*C) with path compression.

**Steps (Union-Find):**

1. Initialize parent[i] = i for each cell i = r*cols+c.  
2. For each (r,c) with grid[r][c]=='1', union with each neighbor that is '1'.  
3. Count roots (parent[i]==i) that correspond to '1'. Return count.

### Python (Union-Find sketch)

Use a 1D parent array; find with path compression; union by rank or simple. Iterate grid, union adjacent '1's, then count distinct roots for '1' cells.

---

## 4. Walkthrough

Example 2: Three separate components of '1's → 3 islands. DFS from first '1' marks one island; from next unvisited '1' marks second; then third.

---

## 5. Edge Cases

1. **Empty grid:** 0.  
2. **All water:** 0.  
3. **All land:** 1 (if connected).  
4. **Single cell '1':** 1.  
5. **Grid with 1 row or 1 column:** Same logic.

---

## 6. Pitfalls

1. **Modifying grid:** Mark visited by setting to '0' (or use separate visited); avoid infinite loop.  
2. **8-direction:** Problem says 4-directional; don’t add diagonals.  
3. **Type:** Grid may be byte or string; compare with '1'.

---

## 7. Follow-ups

1. **Largest island area?** DFS returns area (count); take max.  
2. **Number of closed islands?** Only count islands that don’t touch the border.  
3. **Islands that disappear when water rises?** Sort cells by value, process in order and use Union-Find.  
4. **Pacific Atlantic (LC 417)?** Two DFS from borders.

---

## 8. When to Use

- **Connected components in 2D grid:** DFS/BFS or Union-Find.  
- **Pattern:** Flood fill; same in “max area of island”, “walls and gates”, etc.
