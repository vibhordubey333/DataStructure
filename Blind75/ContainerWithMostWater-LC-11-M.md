# Container With Most Water — LeetCode 11 (M)

## 1. What is the problem?

**LeetCode 11 — Container With Most Water**

You are given an integer array `height` of length n; there are n vertical lines such that the two endpoints of the i-th line are (i, 0) and (i, height[i]). Find two lines that form a container with the x-axis such that the container holds the most water (area = width × min(height[i], height[j])). Return the maximum area.

**Example 1**

- **Input:** `height = [1, 8, 6, 2, 5, 4, 8, 3, 7]`
- **Output:** `49`
- **Explanation:** Lines at index 1 and 8: width = 7, min height = 7 → 7×7 = 49.

**Example 2**

- **Input:** `height = [1, 1]`
- **Output:** `1`

---

## 2. Brute Force Solution

Try every pair (i, j) with i < j; area = (j - i) * min(height[i], height[j]). Return the maximum. Correct by exhaustion.

### Python (brute force)

```python
def max_area_brute(height: list[int]) -> int:
    n = len(height)
    best = 0
    for i in range(n):
        for j in range(i + 1, n):
            area = (j - i) * min(height[i], height[j])
            best = max(best, area)
    return best
```

### Golang (brute force)

```go
func maxAreaBrute(height []int) int {
	best := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			w := j - i
			h := height[i]
			if height[j] < h {
				h = height[j]
			}
			area := w * h
			if area > best {
				best = area
			}
		}
	}
	return best
}
```

- **Time:** O(n²)  
- **Space:** O(1)

---

## 3. Optimized Solution (Two Pointers)

**Idea:** Start with maximum width: pointers at left=0 and right=n-1. Area = (right-left) * min(height[left], height[right]). To try to get a larger area, we must move one pointer; moving the pointer at the **taller** line cannot increase the area (width decreases, height limited by the shorter). So always move the pointer at the **shorter** line. Repeat until left >= right.

**Steps:**

1. left=0, right=len(height)-1, best=0.
2. While left < right:  
   - area = (right - left) * min(height[left], height[right]); best = max(best, area).  
   - If height[left] <= height[right]: left++. Else: right--.  
3. Return best.

### Python (optimized)

```python
def max_area(height: list[int]) -> int:
    left, right = 0, len(height) - 1
    best = 0
    while left < right:
        area = (right - left) * min(height[left], height[right])
        best = max(best, area)
        if height[left] <= height[right]:
            left += 1
        else:
            right -= 1
    return best
```

### Golang (optimized)

```go
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	best := 0
	for left < right {
		w := right - left
		h := height[left]
		if height[right] < h {
			h = height[right]
		}
		area := w * h
		if area > best {
			best = area
		}
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return best
}
```

- **Time:** O(n)  
- **Space:** O(1)

---

## 4. Walkthrough

**Example 1:** `height = [1, 8, 6, 2, 5, 4, 8, 3, 7]`

- left=0, right=8: area = 8*1 = 8; height[0]<height[8] → left=1.
- left=1, right=8: area = 7*7 = 49; best=49; height[1]<height[8] → left=2. (Continue; no area will exceed 49.)
- Eventually best remains 49. Output: 49.

---

## 5. Edge Cases

1. **Two lines:** area = 1 * min(a,b).  
2. **All same height:** area = (j-i)*h; max when j-i max, so left=0, right=n-1.  
3. **Descending then ascending:** two pointers still consider all “good” candidates.  
4. **One tall, rest short:** max might be tall with far end.  
5. **n=2:** one pair only.

---

## 6. Pitfalls

1. **Moving the taller pointer:** That can only reduce width; height is still capped by the shorter, so area cannot increase.  
2. **Starting from middle:** Starting from max width (two ends) and shrinking is the key.  
3. **Equality:** When heights equal, moving either pointer is fine (we move the left in the code above).

---

## 7. Follow-ups

1. **Three lines (two containers)?** Different problem; might use DP or enumerate split.  
2. **Return the two indices?** Track left/right when updating best.  
3. **Trapping rain water (LC 42)?** Different: water above each bar; use two pointers or stack.

---

## 8. When to Use

- **Max area with two boundaries / two pointers from ends:** Shrink the side that limits the value (here the shorter line).  
- **Pattern:** Two pointers at both ends; move the “limiting” pointer to try to improve.
