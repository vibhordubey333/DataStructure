package main

import (
	"fmt"
)

/*
	Time Complexity: O(N^2)
	Space Complexity: O(N^2)
*/

func main() {
	arr := make([]int, 0)
	arr = []int{1,2,2,3}//{2, 3, 5, 6, 8, 10}//{1, 2, 3, 4, 5}
	sum := 3
	n := len(arr)
	fmt.Println("Recursive Approach Result: ", countSubSetsRecursive(arr, sum, n, 0, 0))
	//Expected Output : 3
}

func countSubSetsRecursive(arr []int, sum, n, count, i int) int {

	if sum == 0{
		count = count + 1
		return count
	}
	if n == 0{
		return count
	}

	// If sum is greater than sum then ignore.
	if arr[n-1] > sum {
		return countSubSetsRecursive(arr, sum, n-1, count, i)
	}
	count = countSubSetsRecursive(arr, sum, n-1, count, i)
	count = countSubSetsRecursive(arr, sum-arr[n-1], n-1, count, i)
	return count
}
