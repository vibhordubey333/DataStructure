package main

import (
	"fmt"
	"math"
)

/*
	Bottomup uses the tabluation technique.

	Time Complexity : O(N*W) -> Where N is the number of weight element and W is capacity. As for every weight element we traverse through
								all weight capacities 1<=j<=W
	Space Complexity: O(N*W) -> The use of 2-D array of size N*W.

*/

func main() {
	val := make([]int, 0)
	wt := make([]int, 0)
	W := 50
	val = []int{60, 100, 120}
	wt = []int{10, 20, 30}
	n := 3
	fmt.Println("Profit:", knapsack(wt, val, W, n))
}

func knapsack(wt, val []int, W, n int) int {
	qb := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		qb[i] = make([]int, W+1)
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= W; j++ {
			if i == 0 || j == 0 {
				qb[i][j] = 0
			} else if wt[i-1] <= j {
				qb[i][j] = int(math.Max(float64(val[i-1]+qb[i-1][j-wt[i-1]]), float64(qb[i-1][j]))) // Including the profit as weight is acceptable.
			} else {
				qb[i][j] = qb[i-1][j] // Ignoring the profit as weight is exceeding.
			}
		}
	}
	fmt.Println("Total Profit: ", qb[n][W])
	//Printing Items which have been included in the knapsack
	i, j := n, W
	fmt.Println("Items which are included in knapsack")
	for i > 0 && j > 0 {
		if qb[i][j] == qb[i-1][j] {
			fmt.Println(i, "= 0")
			i--
		} else {
			fmt.Println(i, " = 1")
			i--
			j = j - wt[i]
		}
	}
	//Returning total profit.
	return qb[n][W]
}

/*	========= Output ==========
	Items which are included in knapsack
	3  = 1
	2  = 1
	Profit: 220
*/
