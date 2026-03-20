# Sum of Two Integers — LeetCode 371 (M)

## 1. What is the problem?

**LeetCode 371 — Sum of Two Integers**

Given two integers a and b, return the sum of the two integers without using the operators + and -.

**Example 1**

- **Input:** a = 1, b = 2
- **Output:** 3

**Example 2**

- **Input:** a = 2, b = 3
- **Output:** 5

---

## 2. Brute Force Solution

Without + and -, we can use bit manipulation. “Brute” approach: use a loop that adds 1 or -1 repeatedly (using increment/decrement if allowed), or use sum([a, b]) in Python (often disallowed). The intended approach is bits from the start.

**Alternative “brute” (simulate with bits):** For each bit position, compute sum bit and carry; repeat until no carry. Correct by binary addition definition.

### Python (bit simulation)

```python
def get_sum_brute(a: int, b: int) -> int:
    # Mask to 32 bits for LeetCode (avoid infinite loop with negatives)
    MASK = 0xFFFFFFFF
    a, b = a & MASK, b & MASK
    while b:
        carry = (a & b) << 1
        a = a ^ b
        b = carry & MASK
    return a if a < 0x80000000 else a - (1 << 32)
```

### Golang (optimized / standard approach)

```go
func getSum(a int, b int) int {
	for b != 0 {
		carry := (a & b) << 1
		a ^= b
		b = carry
	}
	return a
}
```

---

## 3. Optimized Solution

**Idea:** Binary addition: sum without carry = a ^ b; carry = (a & b) << 1. So: a, b = a^b, (a&b)<<1; repeat until b (carry) is 0. Then a is the sum.

**Steps:**

1. While b != 0: carry = (a & b) << 1; a = a ^ b; b = carry.  
2. Return a.

For negative numbers in languages with two’s complement (e.g. Go, Java), this works. In Python, integers are arbitrary precision so we mask to 32 bits and then convert back to signed if needed.

### Python (optimized, 32-bit)

```python
def get_sum(a: int, b: int) -> int:
    MASK = 0xFFFFFFFF
    a, b = a & MASK, b & MASK
    while b:
        carry = (a & b) << 1
        a = (a ^ b) & MASK
        b = carry & MASK
    return a if a < 0x80000000 else ~(a ^ MASK)
```

Simpler (if no overflow required): same as Go logic; Python handles big integers.

```python
def get_sum(a: int, b: int) -> int:
    while b:
        a, b = (a ^ b), (a & b) << 1
    return a
```

### Golang (optimized)

```go
func getSum(a int, b int) int {
	for b != 0 {
		carry := (a & b) << 1
		a ^= b
		b = carry
	}
	return a
}
```

- **Time:** O(1) for fixed-width (e.g. 32 bits); number of iterations is at most bit width.  
- **Space:** O(1)

---

## 4. Walkthrough

a = 1 (01), b = 2 (10). a^b = 3, a&b = 0 → carry = 0, b = 0. Return 3.

a = 3, b = 5: 011 ^ 101 = 110 (6), carry = 001<<1 = 010 (2). a=6, b=2. 110^010=100 (4), carry=010<<1=100 (4). a=4, b=4. 100^100=0, carry=100<<1=1000 (8). a=0, b=8. 0^8=8, carry=0. Return 8.

---

## 5. Edge Cases

1. **Both zero:** return 0.  
2. **One zero:** return the other.  
3. **Negative (two’s complement):** Same bit logic works; Python may need masking.  
4. **Overflow:** In 32-bit context, mask to 32 bits.

---

## 6. Pitfalls

1. **Python infinite loop:** Unbounded integers make carry never zero if we don’t mask; use 32-bit mask for LeetCode-style.  
2. **Using + or -:** Even in loop (e.g. a+1) is disallowed.  
3. **Confusing XOR and OR:** Sum bit is XOR; carry is AND.

---

## 7. Follow-ups

1. **Subtract without -?** getSum(a, getSum(~b, 1)) (two’s complement negate b).  
2. **Multiply without *?** Repeated add (slow) or shift-and-add.  
3. **Divide without /?** Repeated subtract or binary long division.

---

## 8. When to Use

- **Add/subtract with bits:** XOR for sum, AND<<1 for carry.  
- **Pattern:** Adder logic; also used in “single number” style problems (XOR cancels pairs).
