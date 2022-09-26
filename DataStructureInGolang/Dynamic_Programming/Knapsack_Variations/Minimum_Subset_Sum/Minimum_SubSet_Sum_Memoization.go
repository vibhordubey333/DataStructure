package main

/*
import (
	"fmt"
	"math"
)

func main() {
	arr := make([]int, 0)
	arr = []int{1, 2, 7}
	n := len(arr)
	findSum(arr, n)
}

func findSum(arr []int, n int) {
	totalSum := 0
	calculatedSum := 0
	qb := make([][]int, n+1)
	for _, v := range arr {
		totalSum += v
	}

	if totalSum == 0 {
		fmt.Print("Minimum SubSetSum is : ", totalSum)
		return
	}

	for i := range qb {
		qb[i] = make([]int, totalSum+1)
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= totalSum; j++ {
			qb[i][j] = -1
		}
	}


	fmt.Println("Minimum SubSetSum is :", findMinSubSetSum(arr, n, totalSum, calculatedSum, qb))
}

func findMinSubSetSum(arr []int, n, totalSum, calculatedSum int, qb [][]int) int {
	//If we've reached last element. Sum of one subset is calculatedSum and sum of other subset is [totalSum-calculatedSum]. Then return absolute difference of two sums.
	if n == 0 {
		qb[n][totalSum] = int(math.Abs(float64((totalSum - calculatedSum) - calculatedSum)))
		return qb[n][totalSum]
	}
	if qb[n][totalSum] != -1 {
		return qb[n][totalSum]
	}

	inclusive := findMinSubSetSum(arr, n-1, totalSum, calculatedSum+arr[n-1], qb)
	exclusive := findMinSubSetSum(arr, n-1, totalSum, calculatedSum, qb)
	qb[n][totalSum] = int(math.Min(float64(inclusive), float64(exclusive)))

	return qb[n][totalSum]
}
*/
