package main

import (
	"fmt"
)

func main() {
	val := make([]int, 0)
	wt := make([]int, 0)
	w := 100
	val = []int{25, 30, 15}
	wt = []int{15, 5, 10}
	n := 3
	fmt.Println("Unbounded Knapsack: ", unboundedKnapsackRecursive(wt, val, w, n))
}

func unboundedKnapsackRecursive(wt, val []int, w, n int) int {

	if n == 0 || w == 0 {
		return 0
	}

	// If value is greater than the given weight don't process
	if wt[n-1] > w {
		return unboundedKnapsackRecursive(wt, val, w, n-1)
	} else {
		//1. nth item will be included.
		//2. nth item will not be included.
		return max(val[n-1]+unboundedKnapsackRecursive(wt, val, w-wt[n-1], n), unboundedKnapsackRecursive(wt, val, w, n-1))
	}

}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

//WorkingCode
func unBoundedKnapsackRecursive_2(wt, val []int, w, n int) int {
	if n == 0 || w == 0 {
		return 0
	}
	excluded := unBoundedKnapsackRecursive_2(wt, val, w, n-1)
	if wt[n-1] > w {
		return excluded
	} else {
		included := val[n-1] + unBoundedKnapsackRecursive_2(wt, val, w-wt[n-1], n)
		return max(included, excluded)
	}
}
