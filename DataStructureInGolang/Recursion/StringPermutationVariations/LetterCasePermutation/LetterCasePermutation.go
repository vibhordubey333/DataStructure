package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "a1B2"
	output := ""
	LetterCasePermutation(input, output)
}

func LetterCasePermutation(input, output string) {
	if len(input) == 0 {
		fmt.Println(" ", output)
		return
	}
	if _, err := strconv.Atoi(string(input[0])); err != nil {
		output1 := output
		output2 := output
		output1 += strings.ToUpper(string(input[0]))
		output2 += strings.ToLower(string(input[0]))
		input = input[1:]
		LetterCasePermutation(input, output1)
		LetterCasePermutation(input, output2)
	} else { // In case if the character is digit then only we've one choice only.
		output1 := output
		output1 += string(input[0])
		input = input[1:]
		LetterCasePermutation(input, output1)
	}

}
