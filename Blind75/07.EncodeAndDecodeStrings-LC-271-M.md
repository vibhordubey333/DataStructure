# Encode and Decode Strings — LeetCode 271 (M, Premium)

## 1. What is the problem?

**LeetCode 271 — Encode and Decode Strings**

Design an algorithm to encode a list of strings to a string and decode the encoded string back to the list. The encoded string must be decodable to the original list.

**Example 1**

- **Input:** ["Hello","World"]
- **Output:** ["Hello","World"]
- **Explanation:** encode → "5#Hello5#World", decode → ["Hello","World"]

**Example 2**

- **Input:** [""]
- **Output:** [""]

---

## 2. Brute Force / Standard Solution

**Length-prefix:** Encode each string as `len#content`. Decode: read len, read len chars, repeat. Handles empty strings and chars like '#'.

### Python

```python
class Codec:
    def encode(self, strs: list[str]) -> str:
        return "".join(f"{len(s)}#{s}" for s in strs)

    def decode(self, s: str) -> list[str]:
        out = []
        i = 0
        while i < len(s):
            j = s.index("#", i)
            length = int(s[i:j])
            out.append(s[j + 1 : j + 1 + length])
            i = j + 1 + length
        return out
```

### Golang

```go
import (
	"strconv"
	"strings"
)

type Codec struct{}

func (c *Codec) Encode(strs []string) string {
	var b strings.Builder
	for _, s := range strs {
		b.WriteString(strconv.Itoa(len(s)))
		b.WriteByte('#')
		b.WriteString(s)
	}
	return b.String()
}

func (c *Codec) Decode(s string) []string {
	var out []string
	i := 0
	for i < len(s) {
		j := i
		for s[j] != '#' {
			j++
		}
		length, _ := strconv.Atoi(s[i:j])
		out = append(out, s[j+1:j+1+length])
		i = j + 1 + length
	}
	return out
}
```

- **Time:** O(n) encode and decode  
- **Space:** O(1) extra (excluding output)

---

## 4. Walkthrough

Encode ["Hello","World"]: "5#Hello5#World". Decode: len=5, read "Hello"; len=5, read "World".

---

## 5. Edge Cases

1. **Empty list:** "" → [""] or [] per spec.  
2. **Empty string in list:** "0#" for "".  
3. **String contains "#":** Length-prefix handles it.  
4. **Single char:** "1#a".

---

## 6. Pitfalls

1. **Delimiter:** Use length-prefix so '#' in content is OK.  
2. **Multi-digit length:** Parse until '#'.  
3. **Chunked transfer:** Same format works for streaming.

---

## 7. Follow-ups

1. **Compression?** Add compression step.  
2. **Binary strings?** Use base64 or length-prefix with byte count.  
3. **Large strings?** Stream encode/decode.

---

## 8. When to Use

- **Encode list of strings:** Length-prefix (len#content).  
- **Pattern:** Same for "Serialize/Deserialize" style problems.
