package main

import (
	"fmt"
)

func main() {
	input := "abc"
	output := ""
	permuteStringWithSpaceAtEndAndBeginning(input, output)
}

func permuteStringWithSpaceAtEndAndBeginning(input, output string) {
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
	permuteStringWithSpaceAtEndAndBeginning(input, output1)
	permuteStringWithSpaceAtEndAndBeginning(input, output2)
	return

}
