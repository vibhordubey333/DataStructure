# Reverse Bits — LeetCode 190 (E)

## 1. What is the problem?

**LeetCode 190 — Reverse Bits**

Reverse the bits of a given 32-bit unsigned integer. (e.g. 00000010100101000001111010011100 → 00111001011110000010100101000000)

**Example 1**

- **Input:** n = 00000010100101000001111010011100 (decimal 43261596)
- **Output:** 964176192 (00111001011110000010100101000000)

**Example 2**

- **Input:** n = 11111111111111111111111111111101
- **Output:** 3221225471

---

## 2. Brute Force Solution

Extract each bit (n >> i) & 1, place it in the reversed position (31-i), build the result. Correct by definition.

### Python (brute force)

```python
def reverse_bits_brute(n: int) -> int:
    out = 0
    for i in range(32):
        bit = (n >> i) & 1
        out |= bit << (31 - i)
    return out
```

### Golang (brute force)

```go
func reverseBitsBrute(n uint32) uint32 {
	var out uint32
	for i := 0; i < 32; i++ {
		bit := (n >> i) & 1
		out |= bit << (31 - i)
	}
	return out
}
```

- **Time:** O(32) = O(1)  
- **Space:** O(1)

---

## 3. Optimized Solution

**Idea:** Same as brute force but expressed as a single loop: for each of 32 bits, take the i-th bit from n and set the (31-i)-th bit in the result. No asymptotically better for 32-bit; can use divide-and-conquer bit swap for same O(1) with fewer operations.

**Steps (simple O(1) approach):**

1. result = 0.
2. For i in 0..31: result |= ((n >> i) & 1) << (31 - i).
3. Return result.

### Python (optimized)

```python
def reverse_bits(n: int) -> int:
    out = 0
    for i in range(32):
        out |= ((n >> i) & 1) << (31 - i)
    return out
```

### Golang (optimized)

```go
func reverseBits(n uint32) uint32 {
	var out uint32
	for i := 0; i < 32; i++ {
		out |= (n >> i & 1) << (31 - i)
	}
	return out
}
```

- **Time:** O(1)  
- **Space:** O(1)

---

## 4. Walkthrough

n = 43261596 (binary ...10100101000001111010011100). Bit 0 is 0 → out bit 31 = 0; bit 2 is 1 → out bit 29 = 1; ... After 32 steps, out = 964176192.

---

## 5. Edge Cases

1. **n = 0:** out = 0.  
2. **n = 0xFFFFFFFF:** out = 0xFFFFFFFF.  
3. **n = 1:** out = 2^31.  
4. **LSB and MSB:** Ensure first and last bits swap correctly.

---

## 6. Pitfalls

1. **Using 31 instead of 32 bits:** Problem says 32-bit.  
2. **Signed integer:** Use unsigned (uint32) in Go; in Python treat as 32-bit (mask if needed).  
3. **Assuming leading zeros:** Reverse all 32 bits.

---

## 7. Follow-ups

1. **Reverse bytes instead of bits?** Swap byte positions.  
2. **64-bit?** Same loop with 64 iterations.  
3. **Reverse bits of a string of '0'/'1'?** Build integer, reverse, then format back.

---

## 8. When to Use

- **Bit reversal:** Loop over bit positions and place in mirror position.  
- **Pattern:** Bit manipulation; same idea in “reverse binary” problems.
