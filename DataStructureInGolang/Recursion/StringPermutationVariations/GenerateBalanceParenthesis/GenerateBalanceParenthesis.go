package main

import (
	"fmt"
)

/*
((())) (()()) (())() ()(()) ()()()
*/

func main() {
	userInput := 3
	open := userInput
	close := userInput
	output := ""
	generateAllBalanceParenthesis(open, close, output)
}

func generateAllBalanceParenthesis(open, close int, output string) {
	if open == 0 && close == 0 {
		fmt.Print(output, " ")
		return
	}
	if open != 0 {
		output1 := output
		output1 += "("
		generateAllBalanceParenthesis(open-1, close, output1)
	}
	if close > open {
		output2 := output
		output2 += ")"
		generateAllBalanceParenthesis(open, close-1, output2)
	}
}
