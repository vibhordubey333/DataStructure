package main

import (
	"fmt"
)

func main() {
	arr := make([]int, 0)
	arr = []int{3, 34, 4, 12, 5, 2}
	n := len(arr)
	//sum := 9999999
	sum := 12
	qb := make([][]bool, n+1)
	for i := range qb {
		qb[i] = make([]bool, sum+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < sum; j++ {
			qb[i][j] = false
		}
	}
	fmt.Println("Output: ", isSubSetSum(arr, n, sum, qb))
}

func isSubSetSum(arr []int, n, sum int, qb [][]bool) bool {
	if sum == 0 {
		return true
	}
	if n <= 0 {
		return false
	}
	if qb[n][sum] != false {
		return qb[n][sum]
	}
	// Checking if last element is greater than sum .
	if arr[n-1] > sum {
		qb[n-1][sum] = isSubSetSum(arr, n-1, sum, qb)
		return qb[n-1][sum]
	}
	exclusion := isSubSetSum(arr, n-1, sum, qb)
	inclusion := isSubSetSum(arr, n-1, sum-arr[n-1], qb)
	qb[n-1][sum] = exclusion || inclusion

	return qb[n-1][sum]
}
