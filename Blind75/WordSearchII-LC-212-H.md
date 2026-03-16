# Word Search II — LeetCode 212 (H)

## 1. What is the problem?

**LeetCode 212 — Word Search II**

Given an m×n board of characters and a list of words, return all words that exist on the board. Words can be constructed from adjacent cells (horizontal/vertical). Same cell cannot be used more than once per word.

**Example 1**

- **Input:** board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]
- **Output:** ["eat","oath"]

**Example 2**

- **Input:** board = [["a","b"]], words = ["abcb"]
- **Output:** []

---

## 2. Brute Force / Standard Solution

**Trie + DFS:** Build Trie from words. For each cell, DFS: if current path is prefix in Trie, continue; if word found, add to result (and optionally remove from Trie to avoid duplicates). Mark cell visited during DFS, unmark on backtrack.

### Python

```python
class TrieNode:
    def __init__(self):
        self.children = {}
        self.word = None  # store word at end

def find_words(board: list[list[str]], words: list[str]) -> list[str]:
    root = TrieNode()
    for w in words:
        node = root
        for c in w:
            if c not in node.children:
                node.children[c] = TrieNode()
            node = node.children[c]
        node.word = w

    rows, cols = len(board), len(board[0])
    res = []

    def dfs(r, c, node):
        if r < 0 or r >= rows or c < 0 or c >= cols:
            return
        ch = board[r][c]
        if ch not in node.children:
            return
        node = node.children[ch]
        if node.word:
            res.append(node.word)
            node.word = None  # avoid duplicate
        board[r][c] = '#'
        for dr, dc in [(0,1),(0,-1),(1,0),(-1,0)]:
            dfs(r+dr, c+dc, node)
        board[r][c] = ch

    for r in range(rows):
        for c in range(cols):
            dfs(r, c, root)
    return res
```

### Golang

```go
type TrieNode struct {
	children [26]*TrieNode
	word     string
}

func findWords(board [][]byte, words []string) []string {
	root := &TrieNode{}
	for _, w := range words {
		node := root
		for _, c := range w {
			idx := c - 'a'
			if node.children[idx] == nil {
				node.children[idx] = &TrieNode{}
			}
			node = node.children[idx]
		}
		node.word = w
	}
	rows, cols := len(board), len(board[0])
	var res []string
	var dfs func(int, int, *TrieNode)
	dfs = func(r, c int, node *TrieNode) {
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return
		}
		ch := board[r][c]
		if ch == '#' {
			return
		}
		idx := ch - 'a'
		if node.children[idx] == nil {
			return
		}
		node = node.children[idx]
		if node.word != "" {
			res = append(res, node.word)
			node.word = ""
		}
		board[r][c] = '#'
		for _, d := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			dfs(r+d[0], c+d[1], node)
		}
		board[r][c] = ch
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			dfs(r, c, root)
		}
	}
	return res
}
```

- **Time:** O(m*n*4^L + total_word_chars) — Trie build + DFS  
- **Space:** O(total_word_chars + L)

---

## 4. Walkthrough

Trie: oath, pea, eat, rain. DFS from (0,0) 'o': o→a→t→h found "oath". From (1,0) 'e': e→a→t found "eat". Prune when node has no children.

---

## 5. Edge Cases

1. **Empty words:** return [].  
2. **Word not on board:** not in result.  
3. **Duplicate words:** Add once; set node.word = None after found.  
4. **Single cell word:** Check that cell.

---

## 6. Pitfalls

1. **Remove word after found:** Avoid adding same word from different paths.  
2. **Prune Trie:** Remove nodes with no children after word found (optional optimization).  
3. **Backtrack:** Restore board[r][c] after DFS.

---

## 7. Follow-ups

1. **Return positions?** Store (r,c) path.  
2. **Count occurrences?** Don't remove node.word.  
3. **Large word list?** Trie prunes early; better than separate DFS per word.

---

## 8. When to Use

- **Multiple words in grid:** Trie + DFS; prune when prefix not in Trie.  
- **Pattern:** Trie for prefix; same as "Word Search" + "Implement Trie".
