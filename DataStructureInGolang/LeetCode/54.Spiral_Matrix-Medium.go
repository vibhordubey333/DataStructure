package main

import "fmt"

/*
Input:
1  2  3  4
5  6  7  8
9  10 11 12
13 14 15 16

Output:

1 2 3 4 8 12 16 15 14 13 9 5 6 7 11 10
*/


func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	spiralMatrix(matrix)
}

func spiralMatrix(matrix [][]int) {
	rowCount := len(matrix) - 1
	colCount := len(matrix[0]) - 1

	top, bottom, left, right := 0, rowCount, 0, colCount
	direction := 0
	resultSpiralMatrix := make([]int, 0)

	for top <= bottom && left <= right {
		//First row
		if direction == 0 {
			for i := left; i <= right; i++ {
				resultSpiralMatrix = append(resultSpiralMatrix, matrix[top][i])
			}
			top++
		} else if direction == 1 { //Nth column
			for i := top; i <= bottom; i++ {
				resultSpiralMatrix = append(resultSpiralMatrix, matrix[i][right])
			}
			right--
		} else if direction == 2 { //Nth row
			for i := right; i >= left; i-- {
				resultSpiralMatrix = append(resultSpiralMatrix, matrix[bottom][i])
			}
			bottom--
		} else if direction == 3 { //First Column
			for i := bottom; i >= top; i-- {
				resultSpiralMatrix = append(resultSpiralMatrix, matrix[i][left])

			}
			left++
		}
		direction = (direction + 1) % 4
	}
	return resultSpiralMatrix
}
