package main

import (
	"fmt"
)

/*
*********************OUTPUT***********************
a_b_c
a_bc
ab_c
abc
*/

func main() {
	inputString := "abc"
	output := ""
	output += string(inputString[0])
	inputString = inputString[1:]
	permutationWithSpace(inputString, output)
}

func permutationWithSpace(input, output string) {
	if len(input) == 0 {
		fmt.Print(" ", output)
		return
	}
	output1 := output
	output2 := output
	output1 += "_"
	output1 += string(input[0])
	output2 += string(input[0])
	input = input[1:]
	permutationWithSpace(input, output1)
	permutationWithSpace(input, output2)

}
