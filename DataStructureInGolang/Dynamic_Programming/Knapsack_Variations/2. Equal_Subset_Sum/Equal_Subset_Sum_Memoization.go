package main

import "fmt"

func main() {
	arr := []int{5, 5, 1, 11, 22}
	sum := 24
	n := len(arr)
	qb := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		qb[i] = make([]bool, sum+1)
	}
	for i := 0; i <= n; i++ {
		for j := 0; j <= sum; j++ {
			qb[i][j] = false
		}
	}
	fmt.Println("isSubSetSumMemoization: ", isSubSetSumMemoization(arr, sum, n, qb))
}

func isSubSetSumMemoization(arr []int, sum, n int, qb [][]bool) bool {
	if sum == 0 {
		return true
	}
	if n == 0 || sum%2 != 0 {
		return false
	}
	if qb[n][sum] != false {
		return qb[n][sum]
	}
	if arr[n-1] > sum {
		return isSubSetSumMemoization(arr, sum, n-1, qb)
	}
	exclusive := isSubSetSumMemoization(arr, sum, n-1, qb)
	inclusive := isSubSetSumMemoization(arr, sum-arr[n-1], n-1, qb)

	qb[n][sum] = exclusive || inclusive

	return qb[n][sum]
}
