package main

import (
	_ "encoding/json"
	"fmt"
	_ "reflect"
)

func main() {
	input := 2
	input_1 := 5
	output := BrianKernighanAlgo(input)
	fmt.Println("Occurrence Counter:", output)
	output = BrianKernighanAlgo(input_1)
	fmt.Println("Occurrence Counter:", output)
}

func BrianKernighanAlgo(input int) []int {
	storeCounter := make([]int, 0)
	i := 0
	//Iterating
	for i <= input {
		count := 0
		temp := i
		for temp > 0 {
			count++
			temp = temp & (temp - 1)
		}
		storeCounter = append(storeCounter, count)
		i++
	}
	return storeCounter
}
