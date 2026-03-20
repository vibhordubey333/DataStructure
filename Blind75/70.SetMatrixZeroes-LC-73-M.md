# Set Matrix Zeroes — LeetCode 73 (M)

## 1. What is the problem?

**LeetCode 73 — Set Matrix Zeroes**

Given an m×n matrix, if an element is 0, set its entire row and column to 0. Do it in-place.

**Example 1**

- **Input:** matrix = [[1,1,1],[1,0,1],[1,1,1]]
- **Output:** [[1,0,1],[0,0,0],[1,0,1]]

**Example 2**

- **Input:** matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
- **Output:** [[0,0,0,0],[0,4,5,0],[0,3,1,0]]

---

## 2. Brute Force Solution

First pass: record all (r,c) where matrix[r][c]==0. Second pass: for each (r,c), set row r and col c to 0. But that would overwrite and create new zeros. So we need to mark first, then set. Use O(m+n) extra space for row_zero[], col_zero[].

### Python (O(m+n) space)

```python
def set_zeroes_brute(matrix: list[list[int]]) -> None:
    rows, cols = len(matrix), len(matrix[0])
    zero_r, zero_c = set(), set()
    for r in range(rows):
        for c in range(cols):
            if matrix[r][c] == 0:
                zero_r.add(r)
                zero_c.add(c)
    for r in zero_r:
        for c in range(cols):
            matrix[r][c] = 0
    for c in zero_c:
        for r in range(rows):
            matrix[r][c] = 0
```

### Golang (O(m+n))

```go
func setZeroes(matrix [][]int) {
	rows, cols := len(matrix), len(matrix[0])
	zeroR, zeroC := make(map[int]bool), make(map[int]bool)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if matrix[r][c] == 0 {
				zeroR[r], zeroC[c] = true, true
			}
		}
	}
	for r := range zeroR {
		for c := 0; c < cols; c++ {
			matrix[r][c] = 0
		}
	}
	for c := range zeroC {
		for r := 0; r < rows; r++ {
			matrix[r][c] = 0
		}
	}
}
```

- **Time:** O(m * n)  
- **Space:** O(m + n)

---

## 3. Optimized Solution (O(1) space)

**Idea:** Use first row and first column as markers. matrix[0][c]=0 means col c has zero; matrix[r][0]=0 means row r has zero. Handle first row/col separately (use two booleans for row0 and col0).

### Python (O(1) space)

```python
def set_zeroes(matrix: list[list[int]]) -> None:
    rows, cols = len(matrix), len(matrix[0])
    row0_zero = any(matrix[0][c] == 0 for c in range(cols))
    col0_zero = any(matrix[r][0] == 0 for r in range(rows))
    for r in range(1, rows):
        for c in range(1, cols):
            if matrix[r][c] == 0:
                matrix[r][0] = matrix[0][c] = 0
    for r in range(1, rows):
        for c in range(1, cols):
            if matrix[r][0] == 0 or matrix[0][c] == 0:
                matrix[r][c] = 0
    if row0_zero:
        for c in range(cols):
            matrix[0][c] = 0
    if col0_zero:
        for r in range(rows):
            matrix[r][0] = 0
```

### Golang (O(1))

```go
func setZeroes(matrix [][]int) {
	rows, cols := len(matrix), len(matrix[0])
	row0Zero, col0Zero := false, false
	for c := 0; c < cols; c++ {
		if matrix[0][c] == 0 {
			row0Zero = true
			break
		}
	}
	for r := 0; r < rows; r++ {
		if matrix[r][0] == 0 {
			col0Zero = true
			break
		}
	}
	for r := 1; r < rows; r++ {
		for c := 1; c < cols; c++ {
			if matrix[r][c] == 0 {
				matrix[r][0], matrix[0][c] = 0, 0
			}
		}
	}
	for r := 1; r < rows; r++ {
		for c := 1; c < cols; c++ {
			if matrix[r][0] == 0 || matrix[0][c] == 0 {
				matrix[r][c] = 0
			}
		}
	}
	if row0Zero {
		for c := 0; c < cols; c++ {
			matrix[0][c] = 0
		}
	}
	if col0Zero {
		for r := 0; r < rows; r++ {
			matrix[r][0] = 0
		}
	}
}
```

- **Time:** O(m * n)  
- **Space:** O(1)

---

## 4. Walkthrough

Mark zeros in matrix[1:][1:]. Use row0,col0 as flags. Set cells where row or col marked. Finally set row0/col0 if needed.

---

## 5. Edge Cases

1. **Single element 0:** entire matrix 0.  
2. **No zeros:** no change.  
3. **First row/col has zero:** Use booleans for row0, col0.  
4. **1×n or m×1:** Handle.

---

## 6. Pitfalls

1. **Don't use row0/col0 for marking before reading:** First scan 1..rows, 1..cols; set markers. Then apply. Finally handle row0, col0.  
2. **Order:** Mark first, then zero; otherwise markers get overwritten.  
3. **In-place:** Modify matrix directly.

---

## 7. Follow-ups

1. **Set to 1 instead of 0?** Same logic.  
2. **Only one row/col?** Simpler.  
3. **Sparse matrix?** Store only (r,c) pairs; same O(m*n) time for dense.

---

## 8. When to Use

- **Propagate to row/col:** Use first row/col as flags for O(1) space.  
- **Pattern:** In-place matrix; use existing structure as storage.
