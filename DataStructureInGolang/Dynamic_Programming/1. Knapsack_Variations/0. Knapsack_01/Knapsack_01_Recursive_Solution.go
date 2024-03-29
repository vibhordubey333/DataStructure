package main

import (
	"fmt"
	"math"
)

/*
	Time Complexity: O(2^n) -> Because there are redundant sub-problems
	Space Complexity: O(1)  -> As no data-structures is used for storing values.
	Max Profit: 220
*/

func main() {
	value := make([]int, 0)
	weight := make([]int, 0)
	W := 50
	N := 3
	value = []int{60, 100, 120}
	weight = []int{10, 20, 30}
	fmt.Println("Profit: ", knapsack_recursive(weight, value, W, N))
}

func knapsack_recursive(weight, value []int, W, N int) int {
	/*
		Lowest valid input for N is 0 as there can be zero elements left to put in a bag.
		W is 0, say in input is given that knapsack can carry 0 amount of weight.
	*/
	if N == 0 || W == 0 {
		return 0
	}
	//Happy case.
	if weight[N-1] <= W {
		return int(math.Max(float64(value[N-1]+knapsack_recursive(weight, value, W-weight[N-1], N-1)), float64(knapsack_recursive(weight, value, W, N-1))))
	} else if weight[N-1] > W { // Non-happy scenario.
		return knapsack_recursive(weight, value, W, N-1)
	}
	return 1
}
