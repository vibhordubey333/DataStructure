package main

import "fmt"

func main() {
	arr := make([]int, 0)
	arr = []int{1, 1, 2, 4} // OP: 1
	n := len(arr)
	sum := 0
	count := 0
	i := 0
	fmt.Println("SubSets Present Are: ", countSubSetWithGivenDiff(arr, n, sum, count, i))
}

func countSubSetWithGivenDiff(arr []int, n, sum, count, i int) int {
	//Here, i acts as a variable to check
	if n == i {
		if sum == 0 {
			count++
		}
		return count
	}

	count = countSubSetWithGivenDiff(arr, n, sum-arr[i], count, i+1)
	count = countSubSetWithGivenDiff(arr, n, sum, count, i+1)

	return count
}
