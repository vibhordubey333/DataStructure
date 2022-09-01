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
	outputString := string(inputString[0]) // OutputString = "a"
	inputString = inputString[1:]          //inputString = "bc"
	permutationWithSpacesInMiddle(inputString, outputString)
}

func permutationWithSpacesInMiddle(inputString, outputString string) {
	if len(inputString) == 0 {
		fmt.Println(" ", outputString)
		return
	}
	output_1 := outputString           //output_1 = "a"
	output_2 := outputString           //output_1 = "b"
	output_1 += "_"                    //output_1 = "a_"
	output_1 += string(inputString[0]) //output_1 = "a_b"
	output_2 += string(inputString[0]) //output_2 ="ab"
	inputString = inputString[1:]      //inputString = "c"
	permutationWithSpacesInMiddle(inputString, output_1)
	permutationWithSpacesInMiddle(inputString, output_2)
}
