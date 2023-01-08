package main

import (
	"fmt"
	"sort"
)

func main() {
	listA := []int{5, 1, 3}
	listB := []int{4, 2, 6}
	lenA := len(listA)
	lenB := len(listB)

	response := mergeUnSortedList_ManualAppend(listA, listB, lenA, lenB)
	fmt.Println("Sorted Merge List Manual Append:", response)

	response2 := mergeUnSortedList_InBuiltAppend(listA, listB)
	fmt.Println("Sorted Merge List InBuild Append:", response2)
}
/*
  1. Sort list A & B.
  2. Merge them manuallly.
*/
func mergeUnSortedList_ManualAppend(listA, listB []int, lenA, lenB int) []int {
	//Sorting listA and listB
	sort.Ints(listA)
	sort.Ints(listB)

	newList := make([]int, lenA+lenB)
	listAIndex, listBIndex, newListIndex := 0, 0, 0

	for listAIndex < lenA && listBIndex < lenB {
		if listA[listAIndex] <= listB[listBIndex] {
			newList[newListIndex] = listA[listAIndex]
			listAIndex++
			newListIndex++
		} else {
			newList[newListIndex] = listB[listBIndex]
			listBIndex++
			newListIndex++
		}
	}
	//Append Remaining Items From ListA to NewList
	for listAIndex < lenA {
		newList[newListIndex] = listA[listAIndex]
		listAIndex++
		newListIndex++
	}

	//Append Remaining Items From ListA to NewList
	for listBIndex < lenB {
		newList[newListIndex] = listB[listBIndex]
		listBIndex++
		newListIndex++
	}
	return newList
}
/*
1. Append list A and B to newList.
2. Sort new list.
*/
func mergeUnSortedList_InBuiltAppend(listA, listB []int) []int {

	newList := append([]int{}, append(listA, listB...)...)
	sort.Ints(newList)

	return newList
}

/*
----------Output----------
Sorted Merge List Manual Append: [1 2 3 4 5 6]
Sorted Merge List InBuild Append: [1 2 3 4 5 6]
*/
