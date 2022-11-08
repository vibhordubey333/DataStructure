package main

import "fmt"
import "math"

func main() {
	val := make([]int, 0)
	wt := make([]int, 0)
	w := 50
	val = []int{60, 100, 120}
	wt = []int{10, 20, 30}
	n := 3
	fmt.Println("Profit:", unBoundedKnapsackBottomsUp(wt, val, w, n))
}

func unBoundedKnapsackBottomsUp(wt, val []int, w, n int) int {
	qb := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		qb[i] = make([]int, w+1)
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= w; j++ {
			//Initializing first row and column
			if i == 0 || j == 0 {
				qb[i][j] = 0
			} else if wt[i-1] <= j {
				qb[i][j] = int(math.Max(float64(val[i-1]+qb[i-1][j-wt[i-1]]), float64(qb[i-1][j])))
			} else {
				//Ignoring the profit when weight is exceeding
				qb[i][j] = qb[i-1][j]
			}
		}
	}
	return qb[n][w]
}
