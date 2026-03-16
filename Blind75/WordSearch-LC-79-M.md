# Word Search — LeetCode 79 (M)

## 1. What is the problem?

**LeetCode 79 — Word Search**

Given an m×n character grid and a string word, return true if word exists in the grid (adjacent cells horizontally or vertically). Same cell cannot be used more than once.

**Example 1**

- **Input:** board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
- **Output:** true

**Example 2**

- **Input:** board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
- **Output:** false

---

## 2. Brute Force Solution

For each cell (r,c), DFS: if board[r][c]==word[i], mark visited, recurse for word[i+1] on neighbors. Unmark on backtrack. Return true if any path completes. O(m*n*4^len(word)).

### Python

```python
def exist(board: list[list[str]], word: str) -> bool:
    if not board or not board[0] or not word:
        return False
    rows, cols = len(board), len(board[0])
    def dfs(r, c, i):
        if i == len(word):
            return True
        if r < 0 or r >= rows or c < 0 or c >= cols or board[r][c] != word[i]:
            return False
        tmp = board[r][c]
        board[r][c] = '#'
        for dr, dc in [(0,1),(0,-1),(1,0),(-1,0)]:
            if dfs(r+dr, c+dc, i+1):
                return True
        board[r][c] = tmp
        return False
    for r in range(rows):
        for c in range(cols):
            if dfs(r, c, 0):
                return True
    return False
```

### Golang

```go
func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 || len(word) == 0 {
		return false
	}
	rows, cols := len(board), len(board[0])
	var dfs func(int, int, int) bool
	dfs = func(r, c, i int) bool {
		if i == len(word) {
			return true
		}
		if r < 0 || r >= rows || c < 0 || c >= cols || board[r][c] != word[i] {
			return false
		}
		board[r][c] = '#'
		defer func() { board[r][c] = word[i] }()
		for _, d := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			if dfs(r+d[0], c+d[1], i+1) {
				return true
			}
		}
		return false
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if dfs(r, c, 0) {
				return true
			}
		}
	}
	return false
}
```

- **Time:** O(m * n * 4^L)  
- **Space:** O(L) recursion

---

## 4. Walkthrough

Start at (0,0) with 'A'. DFS: A→B→C→C→E→D. Mark visited with '#'; restore on backtrack.

---

## 5. Edge Cases

1. **Single char word:** Check any cell.  
2. **Word longer than grid:** false.  
3. **Empty word:** true (or false per problem).  
4. **Repeated chars:** Backtrack must unmark.

---

## 6. Pitfalls

1. **Unmark on backtrack:** Must restore board[r][c] before returning false.  
2. **Index bounds:** Check before accessing.  
3. **Word Search II:** Use Trie + DFS; different.

---

## 7. Follow-ups

1. **Word Search II (multiple words)?** Build Trie from words; DFS and prune.  
2. **Return all paths?** Collect paths and backtrack.  
3. **8-direction?** Add diagonal moves.

---

## 8. When to Use

- **Find path in grid matching string:** DFS + backtrack.  
- **Pattern:** Grid DFS; same in "number of islands", "path sum".
