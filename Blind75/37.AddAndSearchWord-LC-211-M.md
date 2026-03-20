# Design Add and Search Words Data Structure — LeetCode 211 (M)

## 1. What is the problem?

**LeetCode 211 — Design Add and Search Words Data Structure**

Design a data structure that supports addWord and search. search(word) returns true if word is in the structure. word may contain '.' which matches any letter.

**Example 1**

- **Input:** ["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
  [[],["bad"],["dad"],["mad"],[["pad"]],[["bad"]],[[".ad"]],[["b.."]]]
- **Output:** [null,null,null,null,false,true,true,true]

---

## 2. Brute Force / Standard Solution

**Trie + DFS for search:** AddWord: standard Trie insert. Search: when char is '.', try all 26 children; otherwise go to specific child.

### Python

```python
class TrieNode:
    def __init__(self):
        self.children = {}
        self.is_end = False

class WordDictionary:
    def __init__(self):
        self.root = TrieNode()

    def addWord(self, word: str) -> None:
        node = self.root
        for c in word:
            if c not in node.children:
                node.children[c] = TrieNode()
            node = node.children[c]
        node.is_end = True

    def search(self, word: str) -> bool:
        def dfs(node, i):
            if i == len(word):
                return node.is_end
            c = word[i]
            if c == '.':
                for child in node.children.values():
                    if dfs(child, i + 1):
                        return True
                return False
            if c not in node.children:
                return False
            return dfs(node.children[c], i + 1)
        return dfs(self.root, 0)
```

### Golang

```go
type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{root: &TrieNode{}}
}

func (w *WordDictionary) AddWord(word string) {
	node := w.root
	for _, c := range word {
		idx := c - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &TrieNode{}
		}
		node = node.children[idx]
	}
	node.isEnd = true
}

func (w *WordDictionary) Search(word string) bool {
	var dfs func(*TrieNode, int) bool
	dfs = func(node *TrieNode, i int) bool {
		if i == len(word) {
			return node.isEnd
		}
		c := word[i]
		if c == '.' {
			for _, child := range node.children {
				if child != nil && dfs(child, i+1) {
					return true
				}
			}
			return false
		}
		idx := c - 'a'
		if node.children[idx] == nil {
			return false
		}
		return dfs(node.children[idx], i+1)
	}
	return dfs(w.root, 0)
}
```

- **Time:** Add O(m), Search O(26^dots * m) worst  
- **Space:** O(total chars)

---

## 4. Walkthrough

Search ".ad": at root, try all children for '.'. 'b'→'a'→'d', is_end → true. Search "b..": 'b'→ try all for '.', 'a'→ try all for '.', 'd' is_end → true.

---

## 5. Edge Cases

1. **No dots:** Same as Trie search.  
2. **All dots:** Try all paths.  
3. **Empty word:** Check is_end at root.  
4. **Word not added:** false.

---

## 6. Pitfalls

1. **DFS for '.':** Must try all children.  
2. **Early return:** Return true as soon as one path matches.  
3. **Optimization:** Cache by pattern if many repeated searches.

---

## 7. Follow-ups

1. **Only one '.'?** Simpler; try 26 at one level.  
2. **Regex support?** Full NFA/DFA.  
3. **Delete word?** Trie delete with pruning.

---

## 8. When to Use

- **Word search with wildcard:** Trie + DFS on '.'.  
- **Pattern:** Extends "Implement Trie"; same in "Word Search II" (grid + Trie).
