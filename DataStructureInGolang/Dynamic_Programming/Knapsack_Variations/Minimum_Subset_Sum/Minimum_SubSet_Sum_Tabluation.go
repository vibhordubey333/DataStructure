package main

import (
	"fmt"
	"math"
	_ "math"
)

func main() {
	arr := make([]int, 0)
	arr = []int{1, 2, 7} //{1, 2, 3} //
	n := len(arr)
	findSum_Tabluation(arr, n)
}

func findSum_Tabluation(arr []int, n int) {
	totalSum := 0

	for _, v := range arr {
		totalSum += v
	}
	if totalSum == 0 {
		fmt.Println("Minimum SubSet is:", totalSum)
		return
	}
	elementsSlice := findSubSetSum_Tabluation(arr, n, totalSum)
	findMinimumElement(elementsSlice, totalSum)

}

func findMinimumElement(arr []int, totalSum int) {
	minimum := math.MaxInt
	for i := range arr {
		minimum = int(math.Min(float64(minimum), math.Abs(float64(totalSum-2*arr[i]))))
	}
	fmt.Println("Minimum SubSet Sum: ", minimum)
}

func findSubSetSum_Tabluation(arr []int, n, sum int) []int {
	qb := make([][]bool, n+1)

	for i := range qb {
		qb[i] = make([]bool, sum+1)
	}

	//Initialization
	for i := 0; i <= n; i++ {
		for j := 0; j <= sum; j++ {
			if i == 0 {
				qb[i][j] = false
			}
			if j == 0 {
				qb[i][j] = true
			}
		}
	}

	// Taking from the 2nd row and 2nd column
	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			if arr[i-1] <= j {
				qb[i][j] = qb[i-1][j-arr[i-1]] || qb[i-1][j]
			} else {
				//Excluding
				qb[i][j] = qb[i-1][j]
			}
		}
	}

	elementsSlice := make([]int, 0)

	for j := 0; j <= sum; j++ {
		if qb[n][j] == true {
			elementsSlice = append(elementsSlice, j)
		}
	}

	return elementsSlice
}
