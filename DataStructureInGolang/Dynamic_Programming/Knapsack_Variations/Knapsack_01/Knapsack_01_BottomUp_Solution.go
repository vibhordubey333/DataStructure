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
	K := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		K[i] = make([]int, W+1)
	}

	for i := 0; i <= n; i++ {
		for w := 0; w <= W; w++ {
			if i == 0 || w == 0 {
				K[i][w] = 0
			} else if wt[i-1] <= w {
				K[i][w] = int(math.Max(float64(val[i-1]+K[i-1][w-wt[i-1]]), float64(K[i-1][w])))
			} else {
				K[i][w] = K[i-1][w]
			}
		}
	}
	return K[n][W]
}
