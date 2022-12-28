package main

import (
	"fmt"
	"math"
)

func main() {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4} //{4, 5, -3}
	maxSum, elementsList := findMaxSubArraySum(input)
	fmt.Println("MaxSum: ", maxSum, "\nElementsList:", elementsList)
}

func findMaxSubArraySum(input []int) (int, []int) {
	currentMax := 0
	max := math.MinInt
	elementList := make([]int, 0)
	for i := 0; i < len(input); i++ {
		currentMax = currentMax + input[i]
		if currentMax > max {
			elementList = append(elementList, input[i])
			max = currentMax
		}
		if currentMax < 0 {
			currentMax = 0
		}
	}
	return max, elementList
}
/*
MaxSum:  6 
ElementsList: [-2 1 4 2 1]
*/
