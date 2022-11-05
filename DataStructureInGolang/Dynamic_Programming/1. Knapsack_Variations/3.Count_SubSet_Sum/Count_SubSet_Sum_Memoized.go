package main

import (
	"fmt"
)

func main() {
	arr := make([]int, 4)
	arr = []int{1, 2, 2, 3} //{1, 1, 1, 3, 3}
	sum := 3
	n := len(arr)
	qb := make([][]int, n+1)
	for i := range qb {
		qb[i] = make([]int, sum+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < sum; j++ {
			qb[i][j] = -1
		}
	}

	fmt.Println(subSetSumCount_Memoized(arr, sum, n, qb))
}

func subSetSumCount_Memoized(arr []int, sum, n int, qb [][]int) int {

	if sum == 0 {
		return 1
	}

	if n == 0 {
		return 0
	}

	if arr[n-1] > sum {
		qb[n][sum] = subSetSumCount_Memoized(arr, sum, n-1, qb)
		return qb[n][sum]
	}

	exclusive := subSetSumCount_Memoized(arr, sum, n-1, qb)
	inclusive := subSetSumCount_Memoized(arr, sum-arr[n-1], n-1, qb)

	qb[n][sum] = exclusive + inclusive
	return qb[n][sum]
}
