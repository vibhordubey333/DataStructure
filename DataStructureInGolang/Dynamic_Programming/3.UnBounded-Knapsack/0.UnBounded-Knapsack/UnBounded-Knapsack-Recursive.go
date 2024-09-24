package main

import (
	"fmt"
	"math"
)

func main() {
	val := make([]int, 0)
	wt := make([]int, 0)
	w := 100
	val = []int{1, 30}
	wt = []int{1, 50}
	n := 2
	fmt.Println("Profit:", unBoundedKnapsackRecursive(wt, val, w, n)) //Expected Output: Profit: 100
}

func unBoundedKnapsackRecursive(wt, val []int, w, n int) int {
	if w == 0 || n == 0 {
		return 0
	}
	if wt[n-1] <= w {
		return int(math.Max(float64(val[n-1]+unBoundedKnapsackRecursive(wt, val, w-wt[n], n-1)), float64(unBoundedKnapsackRecursive(wt, val, w, n))))
	} else if wt[n-1] > w {
		return unBoundedKnapsackRecursive(wt, val, w, n-1)
	}
	return 1
}
