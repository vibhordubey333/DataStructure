# Decode Ways — LeetCode 91 (M)

## 1. What is the problem?

**LeetCode 91 — Decode Ways**

A message containing letters A–Z can be encoded as '1'→A, '2'→B, … '26'→Z. Given a string `s` of digits, return the number of ways to decode it.

**Example 1**

- **Input:** s = "12"
- **Output:** 2 ("1"+"2" or "12")

**Example 2**

- **Input:** s = "226"
- **Output:** 3 ("2"+"2"+"6", "2"+"26", "22"+"6")

---

## 2. Brute Force Solution

Recurse: try decode 1 char (if valid) and recurse on rest; try 2 chars (if valid) and recurse. Sum both. Exponential.

### Python (brute force)

```python
def num_decodings_brute(s: str) -> int:
    def f(i: int) -> int:
        if i >= len(s):
            return 1
        if s[i] == '0':
            return 0
        ways = f(i + 1)
        if i + 1 < len(s) and 10 <= int(s[i:i+2]) <= 26:
            ways += f(i + 2)
        return ways
    return f(0) if s else 0
```

### Golang (brute force)

```go
func numDecodingsBrute(s string) int {
	var f func(int) int
	f = func(i int) int {
		if i >= len(s) {
			return 1
		}
		if s[i] == '0' {
			return 0
		}
		ways := f(i + 1)
		if i+1 < len(s) {
			n := int(s[i]-'0')*10 + int(s[i+1]-'0')
			if n >= 10 && n <= 26 {
				ways += f(i + 2)
			}
		}
		return ways
	}
	if len(s) == 0 {
		return 0
	}
	return f(0)
}
```

- **Time:** O(2^n)  
- **Space:** O(n) stack

---

## 3. Optimized Solution (DP)

**Idea:** dp[i] = ways to decode s[i:]. dp[n]=1. dp[i] = 0 if s[i]=='0'; else dp[i+1] + (dp[i+2] if valid two-char). Use two variables for O(1) space.

### Python (optimized)

```python
def num_decodings(s: str) -> int:
    if not s:
        return 0
    n = len(s)
    prev2, prev1 = 1, 1 if s[-1] != '0' else 0
    for i in range(n - 2, -1, -1):
        if s[i] == '0':
            cur = 0
        else:
            cur = prev1
            if 10 <= int(s[i:i+2]) <= 26:
                cur += prev2
        prev2, prev1 = prev1, cur
    return prev1
```

### Golang (optimized)

```go
func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	n := len(s)
	prev2, prev1 := 1, 0
	if s[n-1] != '0' {
		prev1 = 1
	}
	for i := n - 2; i >= 0; i-- {
		cur := 0
		if s[i] != '0' {
			cur = prev1
			d := int(s[i]-'0')*10 + int(s[i+1]-'0')
			if d >= 10 && d <= 26 {
				cur += prev2
			}
		}
		prev2, prev1 = prev1, cur
	}
	return prev1
}
```

- **Time:** O(n)  
- **Space:** O(1)

---

## 4. Walkthrough

s="226". Start from end: prev2=1, prev1=1 (s[2]='6'≠'0'). i=1: s[1]='2', two-char 26 valid → cur=1+1=2. i=0: s[0]='2', two-char 22 valid → cur=2+1=3. Return 3.

---

## 5. Edge Cases

1. **Leading zero:** "01" → 0.  
2. **Single "0":** 0.  
3. **Single digit 1–9:** 1.  
4. **"10", "20":** Valid two-char only.  
5. **"27":** Only one-char decode.

---

## 6. Pitfalls

1. **'0' alone:** Cannot decode; return 0.  
2. **Two-char "07":** Invalid (07 not 1–26).  
3. **Iteration direction:** Forward or backward both work; backward often simpler.

---

## 7. Follow-ups

1. **Decode Ways II (with '*')?** Handle '*' as 1–9; more cases.  
2. **Print all decodings?** Backtrack and build strings.  
3. **Mod 10^9+7?** Add mod to avoid overflow.

---

## 8. When to Use

- **Count ways to segment/parse string:** DP over suffixes or prefixes.  
- **Pattern:** Similar to "climbing stairs", "word break" count.
