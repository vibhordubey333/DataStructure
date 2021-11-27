package main

import "fmt"

func main(){
	n := 7
	Print_N_1(n)
}
func Print_N_1(n int){
	if n == 0{ //Base condition, lowest valid input.
		return
	}
	fmt.Println(n) // Induction step.
	Print_N_1(n-1) // Hypothesis i.e how our function will behave.:wq!
}