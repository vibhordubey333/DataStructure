# Spiral Matrix — LeetCode 54 (M)

## 1. What is the problem?

**LeetCode 54 — Spiral Matrix**

Given an m×n matrix, return all elements in spiral order (right, down, left, up, repeat).

**Example 1**

- **Input:** matrix = [[1,2,3],[4,5,6],[7,8,9]]
- **Output:** [1,2,3,6,9,8,7,4,5]

**Example 2**

- **Input:** matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
- **Output:** [1,2,3,4,8,12,11,10,9,5,6,7]

---

## 2. Brute Force / Standard Solution

Maintain boundaries: top, bottom, left, right. Traverse: right → down → left → up. Shrink boundaries. Repeat until done.

### Python

```python
def spiral_order(matrix: list[list[int]]) -> list[int]:
    if not matrix or not matrix[0]:
        return []
    top, bottom = 0, len(matrix) - 1
    left, right = 0, len(matrix[0]) - 1
    out = []
    while top <= bottom and left <= right:
        for c in range(left, right + 1):
            out.append(matrix[top][c])
        top += 1
        for r in range(top, bottom + 1):
            out.append(matrix[r][right])
        right -= 1
        if top <= bottom:
            for c in range(right, left - 1, -1):
                out.append(matrix[bottom][c])
            bottom -= 1
        if left <= right:
            for r in range(bottom, top - 1, -1):
                out.append(matrix[r][left])
            left += 1
    return out
```

### Golang

```go
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1
	var out []int
	for top <= bottom && left <= right {
		for c := left; c <= right; c++ {
			out = append(out, matrix[top][c])
		}
		top++
		for r := top; r <= bottom; r++ {
			out = append(out, matrix[r][right])
		}
		right--
		if top <= bottom {
			for c := right; c >= left; c-- {
				out = append(out, matrix[bottom][c])
			}
			bottom--
		}
		if left <= right {
			for r := bottom; r >= top; r-- {
				out = append(out, matrix[r][left])
			}
			left++
		}
	}
	return out
}
```

- **Time:** O(m * n)  
- **Space:** O(1) excluding output

---

## 4. Walkthrough

3×3: top=0, right→[1,2,3], top=1. right=2, down→[6,9], right=1. bottom=2, left→[8,7], bottom=1. left=0, up→[4], left=1. top=1,bottom=1,left=1,right=1: right→[5]. Done.

---

## 5. Edge Cases

1. **1×1:** single element.  
2. **1×n:** one row.  
3. **m×1:** one column.  
4. **Empty:** return [].

---

## 6. Pitfalls

1. **Check top<=bottom before left traverse:** Avoid duplicate row.  
2. **Check left<=right before up traverse:** Avoid duplicate col.  
3. **Spiral Matrix II (generate):** Same boundaries; fill instead of read.

---

## 7. Follow-ups

1. **Spiral Matrix II (generate n×n)?** Same logic; fill with 1..n².  
2. **Spiral Matrix III (start at (r,c))?** Different; expand outward.  
3. **Diagonal order?** Different traversal.

---

## 8. When to Use

- **Spiral/layer traversal:** Boundaries (top, bottom, left, right); four directions.  
- **Pattern:** Matrix traversal; same in "rotate image" (layer by layer).
