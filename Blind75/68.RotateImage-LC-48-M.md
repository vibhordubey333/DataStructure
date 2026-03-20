# Rotate Image — LeetCode 48 (M)

## 1. What is the problem?

**LeetCode 48 — Rotate Image**

You are given an n×n 2D matrix. Rotate it 90 degrees clockwise in-place.

**Example 1**

- **Input:** matrix = [[1,2,3],[4,5,6],[7,8,9]]
- **Output:** [[7,4,1],[8,5,2],[9,6,3]]

**Example 2**

- **Input:** matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
- **Output:** [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

---

## 2. Brute Force Solution

Create new matrix. new[j][n-1-i] = matrix[i][j]. O(n²) time, O(n²) space. Not in-place.

### Python (optimized — in-place)

**Idea:** Swap in groups of 4: (i,j) → (j,n-1-i) → (n-1-i,n-1-j) → (n-1-j,i) → (i,j). Process each layer; inner layers shrink.

### Python (optimized)

```python
def rotate(matrix: list[list[int]]) -> None:
    n = len(matrix)
    for i in range(n // 2):
        for j in range(i, n - 1 - i):
            top = matrix[i][j]
            matrix[i][j] = matrix[n - 1 - j][i]
            matrix[n - 1 - j][i] = matrix[n - 1 - i][n - 1 - j]
            matrix[n - 1 - i][n - 1 - j] = matrix[j][n - 1 - i]
            matrix[j][n - 1 - i] = top
```

### Golang (optimized)

```go
func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-1-i; j++ {
			top := matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			matrix[j][n-1-i] = top
		}
	}
}
```

- **Time:** O(n²)  
- **Space:** O(1)

---

## 4. Walkthrough

4-way swap: (0,0)→(0,3)→(3,3)→(3,0)→(0,0). (0,1)→(1,3)→(3,2)→(2,0)→(0,1). etc.

---

## 5. Edge Cases

1. **1×1:** no change.  
2. **2×2:** one 4-way swap.  
3. **Odd n:** center unchanged; rest rotate.

---

## 6. Pitfalls

1. **Layer bounds:** j from i to n-2-i (exclusive of n-1-i).  
2. **Swap order:** Must do all 4 in one cycle; order matters.  
3. **Transpose + reverse:** Alternative: transpose and reverse each row. Also O(n²).

---

## 7. Follow-ups

1. **Rotate 180?** Reverse each row and column.  
2. **Rotate 90 counter-clockwise?** Different swap order.  
3. **Transpose only?** transpose: m[i][j], m[j][i] swap for i<j.

---

## 8. When to Use

- **Rotate matrix 90°:** 4-way swap per cell in each layer, or transpose + reverse.  
- **Pattern:** In-place matrix; use symmetry.
