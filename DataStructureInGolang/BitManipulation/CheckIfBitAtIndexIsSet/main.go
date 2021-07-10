package main

import "fmt"

func main() {

	// 13 Binary conversion is  1101
	// 244 Binary conversion is 11110100
	isBitSetAtIndex(13, 1)
	isBitSetAtIndex(244, 1)
	isBitSetAtIndex(13, 0)
	isBitSetAtIndex(244, 2)
}

func isBitSetAtIndex(n, i int) {
	f := 1
	f = f << i //Left shift with ith
	result := f & n
	if result == 0 {
		fmt.Println("False")
	} else {
		fmt.Println("True")
	}
}
