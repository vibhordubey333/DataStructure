# Alien Dictionary — LeetCode 269 (H, Premium)

## 1. What is the problem?

**LeetCode 269 — Alien Dictionary**

There is a new alien language that uses the English alphabet. The order of letters is unknown. You are given a list of words from the dictionary. Derive the order of letters from the words. If invalid (cycle or contradiction), return "".

**Example 1**

- **Input:** words = ["wrt","wrf","er","ett","rftt"]
- **Output:** "wertf"
- **Explanation:** w→e→r→t→f from adjacent word comparisons.

**Example 2**

- **Input:** words = ["z","x"]
- **Output:** "zx"

**Example 3**

- **Input:** words = ["z","x","z"]
- **Output:** ""
- **Explanation:** Cycle z→x→z.

---

## 2. Brute Force / Standard Solution

**Topological sort:** Build graph from adjacent words: first differing char gives edge. Then topological sort (Kahn's or DFS). If cycle, return "".

### Python

```python
def alien_order(words: list[str]) -> str:
    adj = {c: set() for w in words for c in w}
    in_deg = {c: 0 for c in adj}
    for i in range(len(words) - 1):
        a, b = words[i], words[i + 1]
        if len(a) > len(b) and a[:len(b)] == b:
            return ""
        for j in range(min(len(a), len(b))):
            if a[j] != b[j]:
                if b[j] not in adj[a[j]]:
                    adj[a[j]].add(b[j])
                    in_deg[b[j]] += 1
                break
    q = [c for c in adj if in_deg[c] == 0]
    out = []
    while q:
        c = q.pop()
        out.append(c)
        for n in adj[c]:
            in_deg[n] -= 1
            if in_deg[n] == 0:
                q.append(n)
    return "".join(out) if len(out) == len(adj) else ""
```

### Golang

```go
func alienOrder(words []string) string {
	adj := make(map[byte]map[byte]bool)
	inDeg := make(map[byte]int)
	for _, w := range words {
		for _, c := range w {
			b := byte(c)
			if adj[b] == nil {
				adj[b] = make(map[byte]bool)
				inDeg[b] = 0
			}
		}
	}
	for i := 0; i < len(words)-1; i++ {
		a, b := words[i], words[i+1]
		if len(a) > len(b) && a[:len(b)] == b {
			return ""
		}
		for j := 0; j < len(a) && j < len(b); j++ {
			if a[j] != b[j] {
				if !adj[a[j]][b[j]] {
					adj[a[j]][b[j]] = true
					inDeg[b[j]]++
				}
				break
			}
		}
	}
	var q []byte
	for c := range adj {
		if inDeg[c] == 0 {
			q = append(q, c)
		}
	}
	var out []byte
	for len(q) > 0 {
		c := q[len(q)-1]
		q = q[:len(q)-1]
		out = append(out, c)
		for n := range adj[c] {
			inDeg[n]--
			if inDeg[n] == 0 {
				q = append(q, n)
			}
		}
	}
	if len(out) != len(adj) {
		return ""
	}
	return string(out)
}
```

- **Time:** O(C) — total chars in words  
- **Space:** O(1) — 26 letters

---

## 4. Walkthrough

"wrt","wrf": t→f. "wrf","er": w→e. "er","ett": r→t. "ett","rftt": e→r. Graph: w→e→r→t→f. Topo: wertf.

---

## 5. Edge Cases

1. **Single word:** Any order of its letters.  
2. **Prefix order invalid:** ["abc","ab"] → "".  
3. **Empty words:** "".  
4. **All same prefix:** Compare first differing char only.

---

## 6. Pitfalls

1. **Prefix check:** ["ab","a"] invalid (ab cannot come before a).  
2. **Duplicate edges:** Don't double-count in_deg.  
3. **Cycle:** len(out) < len(adj) means cycle.

---

## 7. Follow-ups

1. **Multiple valid orders?** Return any.  
2. **Partial order?** Include all chars that have no constraints.  
3. **Verify given order?** Check against constraints.

---

## 8. When to Use

- **Order from comparisons:** Build DAG, topological sort.  
- **Pattern:** Same as "Course Schedule" (topological sort).
