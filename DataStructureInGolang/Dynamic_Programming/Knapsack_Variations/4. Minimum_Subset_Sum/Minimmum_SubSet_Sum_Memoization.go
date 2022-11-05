package main

import "fmt"

func main() {
	arr := make([]int, 0)
	
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

	for i := range qb {
		qb[i] = make([]int, totalSum+1)
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= totalSum; j++ {
			qb[i][j] = -1
		}
	}
	if totalSum == 0 {
		fmt.Print("Minimum SubSetSum is : ", totalSum)
		return
	}

	fmt.Println("Minimum SubSetSum is :", findMinSubSetSum(arr, n, totalSum, calculatedSum, qb))
}

func findMinSubSetSum(arr []int, n, totalSum, calculatedSum int, qb [][]int) int {
	if n == 0 {

	}
}
