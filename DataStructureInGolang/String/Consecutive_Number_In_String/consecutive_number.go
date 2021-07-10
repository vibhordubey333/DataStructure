package main

import (
	"fmt"
	"strconv"
)

func main() {
	//input := "99100101102"
	//input := "0199100101102"
	input := "099100101102"
	if output := checkConsecutive(input); output != -1{
		fmt.Println("String is consecutive.",output)
	}else{
		fmt.Println("String is not consecutive.")
	}
}

func checkConsecutive(input string) int{

	length := len(input)
	start := 0
	for i:= 0 ; i < length/2 ; i++{

		//New string containing the starting
		newStr := input[0:i+1]
		//Converting string to integer
		num , _:= strconv.Atoi(newStr)
		start = num
		for len(newStr) < length{
			num++
			//Converting integer to string to concatenate
			newStr = newStr + strconv.Itoa(num)
		}
		if (newStr == input){
			return start
		}
	}
	return -1
}