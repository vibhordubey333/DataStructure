package main

import (
	"fmt"
	"math"
)

func main() {
	arr := make([]int, 0)
	arr = []int{1, 2, 7}
	//arr = []int{}
	n := len(arr)
	findSum(arr, n)
}

func findSum(arr []int, n int) {

	totalSum := 0
	calculatedSum := 0
	for _, v := range arr {
		totalSum += v
	}

	if totalSum == 0 {
		fmt.Println("Minimum SubSetSum: ", totalSum)
		return
	}

	fmt.Println("Minimum SubSetSum", findMinSubSetSum(arr, n, totalSum, calculatedSum))
}

func findMinSubSetSum(arr []int, n, totalSum, calculatedSum int) int {

	//If we've reached last element. Sum of one subset is calculatedSum and sum of other subset is [totalSum-calculatedSum]. Then return absolute difference of two sums.
	if n == 0 {
		return int(math.Abs(float64((totalSum - calculatedSum) - calculatedSum)))
	}

	return int(math.Min(float64(findMinSubSetSum(arr, n-1, totalSum, calculatedSum+arr[n-1])), float64(findMinSubSetSum(arr, n-1, totalSum, calculatedSum))))
}
