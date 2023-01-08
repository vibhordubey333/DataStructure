package main

import (
	"fmt"
)

func main() {
	arr := []int{89, 12, 2, 23, 1, 2, 334, 2, 1, 2, 99}
	findFrequencyOfElements(arr)
}
func findFrequencyOfElements(arr []int) {
	frequencyCounter := make(map[int]int, 0)
	for _, v := range arr {
		frequencyCounter[v] = frequencyCounter[v] + 1
	}
	fmt.Println("Frequency Count Of All : ", frequencyCounter)
	for k, v := range frequencyCounter {
		if v == 1 {
			fmt.Println("UniqueElement :", k)
		}
	}
}
