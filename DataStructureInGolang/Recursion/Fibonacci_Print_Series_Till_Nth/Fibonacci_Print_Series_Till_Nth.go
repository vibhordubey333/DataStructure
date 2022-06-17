package main

import (
	"fmt"
)

/*Notes:
1. If while calling printFibonacci we do n-2 then pattern for N=5 will be [0,1,1,2,3]. Comment this line below fmt.Printf(" 0 || 1 ||")
2. And if we pass N=5 as it is then output will be [1,2,3,5,8]
3. Sample Fibonacci till 15
	0 || 1 ||1 || 2 || 3 || 5 || 8 || 13 || 21 || 34 || 55 || 89 || 144 || 233 || 377 || 610 || 987
*/

func main() {
	firstNum := 0
	secondNum := 1
	n := 5
	fmt.Printf(" 0 || 1 ||")
	printFibonacci(firstNum, secondNum, n)
}

func printFibonacci(firstNum, secondNum, n int) {
	sum := 0
	if n > 0 {
		sum = firstNum + secondNum
		fmt.Printf("%d || ", sum)
		firstNum, secondNum = secondNum, sum
		printFibonacci(firstNum, secondNum, n-1)
	}
}
