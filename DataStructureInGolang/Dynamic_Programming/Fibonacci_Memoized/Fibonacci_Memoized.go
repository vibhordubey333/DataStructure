package main

import (
	"fmt"
)

var (
	N = 10 // Output should be 55
	//N  = 0				// Output should be 0
	qb = make([]int, N+1) // qb is our question bank from where we'll fetch the problem output which we've already solved.
	// Using N+1 to store the Output
)

func main() {
	for i := range qb {
		qb[i] = -1
	}
	fmt.Println("Fibonacci", Fibonacci(N))
}

func Fibonacci(N int) int {
	if N <= 0 {
		return N
	}
	if N <= 2 {
		return 1
	}

	if qb[N] != -1 {
		return qb[N]
	}
	qb[N] = Fibonacci(N-1) + Fibonacci(N-2)
	return qb[N]

}
