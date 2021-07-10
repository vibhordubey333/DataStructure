package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 5; i++ {
		output := LeftShift(5, i)
		fmt.Println("LeftBitShift : ", output)
	}
}
func LeftShift(n, i int) int {
	fmt.Println("N:", n, " i:", i)
	output := n << i
	return output
}
