package main

import (
	"fmt"
)

func main() {
	arr := make([]int, 0)
	arr = []int{5, 5, 1, 11, 22}
	sum := 22
	n := len(arr)
	fmt.Println(isEqualSumPartitionExist(arr, n, sum))
}

func isEqualSumPartitionExist(arr []int, n, sum int) bool {
	if sum == 0 {
		return true
	}
	//Performing sum%2 as we've to check if they can be divided into equal subset.
	if n == 0 || sum%2 != 0 || sum < 0 { // Return false if n and sum is not even or is sum is less than zero.
		return false
	}
	if arr[n-1] > sum {
		return isEqualSumPartitionExist(arr, n-1, sum)
	}

	inclusive := isEqualSumPartitionExist(arr, n-1, sum-arr[n-1])
	exclusive := isEqualSumPartitionExist(arr, n-1, sum)
	return inclusive || exclusive

}
