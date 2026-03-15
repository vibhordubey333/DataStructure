# Number of 1 Bits — LeetCode 191 (E)

## 1. What is the problem?

**LeetCode 191 — Number of 1 Bits (Hamming Weight)**

Write a function that takes the binary representation of an unsigned 32-bit integer and returns the number of '1' bits (popcount).

**Example 1**

- **Input:** n = 11 (binary 1011)
- **Output:** 3

**Example 2**

- **Input:** n = 128 (binary 10000000)
- **Output:** 1

---

## 2. Brute Force Solution

Shift and count: while n > 0, add n & 1, then n >>= 1. Correct by definition.

### Python (brute force)

```python
def hamming_weight_brute(n: int) -> int:
    c = 0
    while n:
        c += n & 1
        n >>= 1
    return c
```

### Golang (brute force)

```go
func hammingWeightBrute(n uint32) int {
	c := 0
	for n > 0 {
		c += int(n & 1)
		n >>= 1
	}
	return c
}
```

- **Time:** O(32) = O(1)  
- **Space:** O(1)

---

## 3. Optimized Solution

**Idea:** n & (n-1) clears the lowest set bit. So repeat: count++, n = n & (n-1) until n == 0. Number of iterations = number of 1 bits.

### Python (optimized)

```python
def hamming_weight(n: int) -> int:
    c = 0
    while n:
        n &= n - 1
        c += 1
    return c
```

### Golang (optimized)

```go
func hammingWeight(n uint32) int {
	c := 0
	for n != 0 {
		n &= n - 1
		c++
	}
	return c
}
```

- **Time:** O(number of 1 bits)  
- **Space:** O(1)

---

## 4. Walkthrough

n = 11 (1011). n & (n-1) = 1011 & 1010 = 1010; then 1010 & 1001 = 1000; then 1000 & 0111 = 0. Count = 3.

---

## 5. Edge Cases

1. **n = 0:** return 0.  
2. **n = 0xFFFFFFFF:** return 32.  
3. **Power of two:** one iteration per 1 bit (e.g. 128 → 1).

---

## 6. Pitfalls

1. **Signed right shift:** Use unsigned (uint32) in Go so no sign extension.  
2. **In Python:** Input may be given as signed; mask to 32 bits if needed: n &= 0xFFFFFFFF.  
3. **Infinite loop:** n & (n-1) always removes one 1, so loop terminates.

---

## 7. Follow-ups

1. **Count 0 bits?** 32 - hamming_weight(n) for 32-bit.  
2. **Parity?** Same loop, return count % 2; or XOR fold.  
3. **Next number with same popcount?** Bit manipulation (next permutation of bits).

---

## 8. When to Use

- **Popcount / clear lowest set bit:** n & (n-1) is a standard trick.  
- **Pattern:** Used in many bit problems (e.g. power of two: n > 0 && (n & (n-1)) == 0).
