package main

import (
	"fmt"
)

func main() {
	a := make([]int, 0)
	b := make([]int, 0)
	a = []int{1, 2, 3, 4, 5}
	b = []int{11, 12, 13, 14, 15, 195}
	response := merge(a, b, len(a), len(b))
	fmt.Println("Merged Output: ", response)
}
func merge(a []int, b []int, lenA, lenB int) []int {
	defer panicHandler()
	i, j := 0, 0
	tempArr := make([]int, 0)

	for i < lenA && j < lenB {
		if a[i] < b[j] {
			tempArr = append(tempArr, a[i])
			i++
		} else {
			tempArr = append(tempArr, b[j])
			j++
		}
	}
	//Handling scenarios when extra elements are left in either of the list.
	for ; i <= lenA-1; i++ {
		tempArr = append(tempArr, a[i])
	}
	for ; j <= lenB-1; j++ {
		tempArr = append(tempArr, b[j])
	}
	return tempArr
}
func panicHandler() {
	if r := recover(); r != nil {
		fmt.Println("Recovered: ", r)
	}
}
