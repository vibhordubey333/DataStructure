/*
1) Use two index variables i and j, initial values i = 0, j = 0
2) If arr1[i] is smaller than arr2[j] then increment i.
3) If arr1[i] is greater than arr2[j] then increment j.
4) If both are same then print any of them and increment both i and j.
Complexity:  O(i+j)
Space Complexity :  O(1)

*/
package main

import "log"

func main() {
	listA := make([]int, 0)

	listB := make([]int, 0)

	listA = []int{1, 2, 3, 4, 5}

	listB = []int{1, 2, 3, 5, 6, 7, 8, 5, 4, 5}

	lenA := len(listA)

	lenB := len(listB)

	intersectionResolver(listA, lenA, listB, lenB)
}

func intersectionResolver(listA []int, lenA int, listB []int, lenB int) {
	var i int = 0
	var j int = 0
	/*
		listA = []int{1, 2, 3, 4, 5}
		listB = []int{1, 2, 3, 5, 6, 7, 8, 5, 4, 5}
	*/
	for i < lenA && j < lenB {
		if listA[i] < listB[j] {
			i++
		} else if listA[i] > listB[j] {

			j++
		} else {
			log.Println(listA[i]) // Print any of the two list here doesn't matter as equalent element is found now.
			i++
			j++
		}
	}
}
