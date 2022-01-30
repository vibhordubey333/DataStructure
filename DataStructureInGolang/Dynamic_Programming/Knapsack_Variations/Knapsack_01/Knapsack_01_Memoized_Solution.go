package main

import (
	"fmt"
	"math"
)

func main() {
	value := make([]int, 0)
	weight := make([]int, 0)
	N := 3
	qb := make([][]int, 1002)
	W := 50
	value = []int{60, 100, 120}
	weight = []int{10, 20, 30}

	for i := range qb {
		qb[i] = make([]int, 1002)
	}
	for i := 0; i < 1002; i++ {
		for j := 0; j < 1002; j++ {
			qb[i][j] = -1
		}
	}

	fmt.Println("Profit: ", knapsackMemoized(weight, value, W, N, qb))
}

func knapsackMemoized(weight, value []int, W, N int, qb [][]int) int {
	if W == 0 || N == 0 {
		return 0
	}
	if qb[W][N] != -1 {
		return qb[W][N]
	}
	if weight[N-1] <= W {
		return int(math.Max(float64(value[N-1]+knapsackMemoized(weight, value, W-weight[N-1], N-1, qb)), float64(knapsackMemoized(weight, value, W, N-1, qb))))
	} else if weight[N-1] > W {
		knapsackMemoized(weight, value, W, N-1, qb)
	}
	return 1
}
