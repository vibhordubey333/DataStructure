package main

import (
	"fmt"
)

//Complexity: O(N)
//Auxillary Space: O(1)

func main() {

	n, k := 5, 2
	fmt.Println(Josephus(n, k))
}

func Josephus(n, k int) int {
	sum := 0
	for i := 2; i <= n; i++ {
		sum = (sum + k) % i
	}
	return sum + 1
}
