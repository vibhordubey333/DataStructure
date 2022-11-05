package main

import (
	"fmt"
	"math"
)

/*
	Topdown uses memoization technique.
	Time Complexity:	O(N*W)	As redundant calculations of states are avoided.
						Memoization and Bottomup have same complexity but upper hand is
						of bottom-up as there is no stackoverflow error.

	Space Complexity:	o(N*W)
*/

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
	/*for i := 0; i < 1002; i++ {
		for j := 0; j < 1002; j++ {
			qb[i][j] = -1
		}
	}*/
	for i := range qb {
		for j := range qb {
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
		qb[W][N] = int(math.Max(float64(value[N-1]+knapsackMemoized(weight, value, W-weight[N-1], N-1, qb)), float64(knapsackMemoized(weight, value, W, N-1, qb))))
		return qb[W][N]
	} else if weight[N-1] > W {
		qb[W][N] = knapsackMemoized(weight, value, W, N-1, qb)
		return qb[W][N]
	}
	return 1
}
