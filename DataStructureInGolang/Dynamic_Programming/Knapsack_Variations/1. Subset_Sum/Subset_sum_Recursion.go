package main

import (
	"fmt"
)

/*
Time Complexity: O(2^N)
This problem is NP-Complete(There is no known polynoial time solution for this problem)
Space Complexity: O(1)
As no data structure is used to store the values
*/
func main() {
	arr := make([]int, 0)
	arr = []int{3, 34, 4, 12, 5, 2}
	n := len(arr)
	sum := 9
	fmt.Println("Output: ", isSubSetSum_Recursive(arr, n, sum))
}

func isSubSetSum_Recursive(arr []int, n, sum int) bool {
	//Base Cases
	if sum == 0 {
		return true
	}
	if n == 0 { //IF N is 0 it means it will be able to make 0 subsets so returning as false.
		return false
	}
	//If last element is greater than sum than ignore it.
	if arr[n-1] > sum {
		isSubSetSum_Recursive(arr, n-1, sum)
	}

	/*
		Checking if sum can be obtained by either
		A. Including the last element.
		B. Excluding the last element i.e sum-arr[n-1]
	*/

	return isSubSetSum_Recursive(arr, n-1, sum) || isSubSetSum_Recursive(arr, n-1, sum-arr[n-1])
}
