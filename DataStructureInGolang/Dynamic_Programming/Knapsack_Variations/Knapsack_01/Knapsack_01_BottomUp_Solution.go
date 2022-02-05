package main

import (
	"fmt"
	"math"
)

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
				qb[i][j] = int(math.Max(float64(val[i-1]+qb[i-1][j-wt[i-1]]), float64(qb[i-1][j])))
			} else {
				qb[i][j] = qb[i-1][j]
			}
		}
	}
	return qb[n][W]
}
