package main

import (
	"fmt"
)

var (
//N = 10 // Output should be 55
//N  = 0				// Output should be 0
// qb is our question bank from where we'll fetch the problem output which we've already solved.
)

func main() {
	N := 10 //Output : 55
	qb := make(map[int]int)
	qb[0] = 0
	qb[1] = 1
	fmt.Println("Fibonacci", CalculateFibonacci(N, qb))
}

func CalculateFibonacci(n int, qb map[int]int) int {

	value, ok := qb[n]
	if ok {
		return value
	} else {
		qb[n] = CalculateFibonacci(n-1, qb) + CalculateFibonacci(n-2, qb)
		return qb[n]
	}
}
