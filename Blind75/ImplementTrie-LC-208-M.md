# Implement Trie (Prefix Tree) — LeetCode 208 (M)

## 1. What is the problem?

**LeetCode 208 — Implement Trie (Prefix Tree)**

Implement a Trie with insert, search, and startsWith. Assume lowercase letters a-z.

**Example 1**

- **Input:** ["Trie","insert","search","search","startsWith","insert","search"]
  [[],["apple"],[["apple"]],[["app"]],[["app"]],[["app"]],[["app"]]]
- **Output:** [null,null,true,false,true,null,true]

---

## 2. Brute Force / Standard Solution

**Trie node:** children[26], is_end. Insert: traverse, create nodes, mark end. Search: traverse, check is_end. startsWith: traverse, no need for is_end.

### Python

```python
class TrieNode:
    def __init__(self):
        self.children = {}
        self.is_end = False

class Trie:
    def __init__(self):
        self.root = TrieNode()

    def insert(self, word: str) -> None:
        node = self.root
        for c in word:
            if c not in node.children:
                node.children[c] = TrieNode()
            node = node.children[c]
        node.is_end = True

    def search(self, word: str) -> bool:
        node = self.root
        for c in word:
            if c not in node.children:
                return False
            node = node.children[c]
        return node.is_end

    def startsWith(self, prefix: str) -> bool:
        node = self.root
        for c in prefix:
            if c not in node.children:
                return False
            node = node.children[c]
        return True
```

### Golang

```go
type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{root: &TrieNode{}}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, c := range word {
		idx := c - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &TrieNode{}
		}
		node = node.children[idx]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, c := range word {
		idx := c - 'a'
		if node.children[idx] == nil {
			return false
		}
		node = node.children[idx]
	}
	return node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, c := range prefix {
		idx := c - 'a'
		if node.children[idx] == nil {
			return false
		}
		node = node.children[idx]
	}
	return true
}
```

- **Time:** O(m) per operation (m = word/prefix length)  
- **Space:** O(total chars)

---

## 4. Walkthrough

Insert "apple": a→p→p→l→e, mark end. Search "app": a→p→p, not end → false. startsWith "app": a→p→p → true. Insert "app", search "app" → true.

---

## 5. Edge Cases

1. **Empty word:** insert "" marks root? Usually not.  
2. **Prefix of existing:** "app" prefix of "apple"; insert "app" adds end at p.  
3. **Duplicate insert:** Idempotent; just mark end.  
4. **Search empty:** Return false or true per spec.

---

## 6. Pitfalls

1. **is_end:** Distinguishes "app" (word) from prefix of "apple".  
2. **children:** Use array [26] for a-z or dict.  
3. **Delete:** More complex; prune unused nodes.

---

## 7. Follow-ups

1. **Word Search II?** Trie + DFS on grid.  
2. **Add and Search Word (.)?** DFS on '.' to try all children.  
3. **Autocomplete?** DFS from prefix, collect all words.

---

## 8. When to Use

- **Prefix/word lookup:** Trie for O(m) insert/search/startsWith.  
- **Pattern:** Base for "Word Search II", "Design Add and Search Words".
