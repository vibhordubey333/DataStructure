package main

import (
	"fmt"
)

func main() {
	arr := make([]int, 0)
	arr = []int{3, 34, 4, 12, 5, 2}
	sum := 90
	n := len(arr)
	qb := make([][]bool, n+1)
	for i := range qb {
		qb[i] = make([]bool, sum+1)
	}
	fmt.Println("Output", subsetsumBottomUp(arr, sum, n, qb))
}

func subsetsumBottomUp(arr []int, sum, n int, qb [][]bool) bool {

	for i := 0; i <= n; i++ {
		qb[i][0] = true
	}
	for i := 1; i <= sum; i++ {
		qb[0][i] = false
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			//Ignoring the element
			if arr[i-1] > j {
				qb[i][j] = qb[i-1][j]
			} else {
				exclusion := qb[i-1][j]
				inclusion := qb[i-1][j-arr[i-1]]
				qb[i][j] = exclusion || inclusion
			}
		}
	}
	// Printing the table
	for i := 0; i <= n; i++ {
		for j := 0; j <= sum; j++ {
			fmt.Print(" ", qb[i][j])
		}
		fmt.Println()
	}

	return qb[n][sum]
}
