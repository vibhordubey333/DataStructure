package main

import "fmt"

func main() {
	userInput := 7
	limit := 10
	offset := 1
	printTable(userInput, limit, offset)
}

func printTable(userInput, limit, offset int) {
	if offset == limit+1 { // Base
		return
	}
	fmt.Println(userInput, "*", offset, "=", userInput*offset) //Induction
	printTable(userInput, limit, offset+1)                     //Hypothesis
}
