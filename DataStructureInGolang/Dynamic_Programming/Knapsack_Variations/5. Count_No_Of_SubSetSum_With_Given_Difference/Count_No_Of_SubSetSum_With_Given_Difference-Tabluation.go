package main

import "fmt"

func main() {
	n := 9
	arr := make([]int, n+1)
	//arr = []int{1, 2, 3, 1, 2} //diff=1 // n=5 // output=5
	arr = []int{0, 0, 0, 0, 0, 0, 0, 0, 1} // diff=1 // n=9 //output=256
	diff := 1
	fmt.Println("Count Of SubSets With Given Difference Is: ", CountSubSetsWithDiff(arr, n, diff))
}

func CountSubSetsWithDiff(arr []int, n, diff int) int {
	sumOfArray := 0
	for i := 0; i < n; i++ {
		sumOfArray += arr[i]
	}

	if (sumOfArray+diff)%2 != 0 {
		return 0
	} else {
		return countNoOfSubsets_Tabluation(arr, n, (sumOfArray+diff)/2)
	}
}

func countNoOfSubsets_Tabluation(arr []int, n, sum int) int {
	qb := make([][]int, n+1)

	for i := range qb {
		qb[i] = make([]int, sum+1)
	}

	for i := 0; i <= n; i++ {
		//Starting inner loop from zero as it handles this case {0,0,0,0,0,0,0,0,1}
		for j := 0; j <= sum; j++ {
			if i == 0 {
				qb[i][j] = 0
			}
			if j == 0 {
				qb[i][j] = 1
			}
		}
	}
	// i denotes size of the array
	// j denotes the target sum
	for i := 1; i <= n; i++ {
		for j := 0; j <= sum; j++ {
			if arr[i-1] <= j { //Including + exluding
				qb[i][j] = qb[i-1][j-arr[i-1]] + qb[i-1][j]
			} else { //Including
				qb[i][j] = qb[i-1][j]
			}
		}
	}

	return qb[n][sum]
}
