# Pacific Atlantic Water Flow — LeetCode 417 (M)

## 1. What is the problem?

**LeetCode 417 — Pacific Atlantic Water Flow**

Given an m×n matrix `heights` (height of each cell), rain flows to adjacent cells with height ≤ current. Return list of cells that can flow to **both** Pacific (top/left) and Atlantic (bottom/right) oceans.

**Example 1**

- **Input:** heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]
- **Output:** [[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]

**Example 2**

- **Input:** heights = [[1]]
- **Output:** [[0,0]]

---

## 2. Brute Force Solution

For each cell (r,c), DFS to see if it can reach Pacific (top/left border) and DFS to see if it can reach Atlantic (bottom/right border). O((m*n)^2) in worst case. Correct but slow.

### Python (optimized — reverse DFS from oceans)

**Idea:** Instead of "can (r,c) reach ocean?", do "which cells can ocean reach?" — DFS from Pacific border (cells with r=0 or c=0) going **uphill** (to higher or equal); mark reachable. Same from Atlantic border. Cells reachable from both are the answer.

### Python (optimized)

```python
def pacific_atlantic(heights: list[list[int]]) -> list[list[int]]:
    if not heights or not heights[0]:
        return []
    rows, cols = len(heights), len(heights[0])
    pacific = [[False] * cols for _ in range(rows)]
    atlantic = [[False] * cols for _ in range(rows)]
    def dfs(r, c, ocean):
        ocean[r][c] = True
        for dr, dc in [(0,1),(0,-1),(1,0),(-1,0)]:
            nr, nc = r + dr, c + dc
            if 0 <= nr < rows and 0 <= nc < cols and not ocean[nr][nc] and heights[nr][nc] >= heights[r][c]:
                dfs(nr, nc, ocean)
    for c in range(cols):
        dfs(0, c, pacific)
        dfs(rows - 1, c, atlantic)
    for r in range(rows):
        dfs(r, 0, pacific)
        dfs(r, cols - 1, atlantic)
    return [[r, c] for r in range(rows) for c in range(cols) if pacific[r][c] and atlantic[r][c]]
```

### Golang (optimized)

```go
func pacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return nil
	}
	rows, cols := len(heights), len(heights[0])
	pacific := make([][]bool, rows)
	atlantic := make([][]bool, rows)
	for i := range pacific {
		pacific[i] = make([]bool, cols)
		atlantic[i] = make([]bool, cols)
	}
	var dfs func(int, int, [][]bool)
	dfs = func(r, c int, ocean [][]bool) {
		ocean[r][c] = true
		for _, d := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && !ocean[nr][nc] && heights[nr][nc] >= heights[r][c] {
				dfs(nr, nc, ocean)
			}
		}
	}
	for c := 0; c < cols; c++ {
		dfs(0, c, pacific)
		dfs(rows-1, c, atlantic)
	}
	for r := 0; r < rows; r++ {
		dfs(r, 0, pacific)
		dfs(r, cols-1, atlantic)
	}
	var out [][]int
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if pacific[r][c] && atlantic[r][c] {
				out = append(out, []int{r, c})
			}
		}
	}
	return out
}
```

- **Time:** O(m * n)  
- **Space:** O(m * n)

---

## 4. Walkthrough

Start DFS from Pacific border (top row, left col) going uphill. Mark all reachable. Same from Atlantic (bottom, right). Intersection = answer.

---

## 5. Edge Cases

1. **1×1:** [[0,0]].  
2. **All same height:** All cells reach both.  
3. **Strictly increasing from coast:** Only border cells.  
4. **Empty matrix:** [].

---

## 6. Pitfalls

1. **Direction:** Water flows **downhill** (to lower/equal). Reverse: from ocean we go **uphill** (to higher/equal).  
2. **Both oceans:** Must reach Pacific AND Atlantic.  
3. **Borders:** Pacific = top + left; Atlantic = bottom + right.

---

## 7. Follow-ups

1. **Trapping rain water II?** Min-heap from border; different.  
2. **Longest increasing path?** DFS with memo.  
3. **Count cells?** Same approach; count intersection.

---

## 8. When to Use

- **Multi-source reachability:** DFS/BFS from multiple sources (borders).  
- **Pattern:** "Reverse" flow — start from destinations, go backward.
