// O/P should be
/*
c b bc a ac ab abc
*/
package main

import (
	"fmt"
)

func main() {
	inputString := "abc"
	output := ""
	printSubsets(inputString, output)
}

func printSubsets(inputString, output string) {
	if len(inputString) == 0 {
		fmt.Print(" " + output)
		return
	}

	output1 := output
	output2 := output
	output2 += string(inputString[0])

	inputString = inputString[1:] //Converting into smaller string

	printSubsets(inputString, output1)
	printSubsets(inputString, output2)

}
