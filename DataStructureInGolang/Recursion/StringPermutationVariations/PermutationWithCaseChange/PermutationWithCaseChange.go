package main

import (
	"fmt"
	"strings"
)

/*
Assuming input will always be in small case.
Input: ab
Output: AB Ab aB ab
*/
func main() {
	input := "ab"
	output := ""
	permutationWithCaseChange(input, output)
}
func permutationWithCaseChange(input, output string) {
	if len(input) == 0 {
		fmt.Print(" ", output)
		return
	}
	output1 := output
	output2 := output
	output1 += strings.ToUpper(string(input[0]))
	output2 += string(input[0])
	input = input[1:]
	permutationWithCaseChange(input, output1)
	permutationWithCaseChange(input, output2)

}
