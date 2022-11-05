package main

import (
	"fmt"
)

/*BottomUp or tabluation method.
Time Complexity: O(sum*n)
Space Complexity: O(sum*n)
https://www.geeksforgeeks.org/count-of-subsets-with-sum-equal-to-x/
*/

func main() {
	n := 3
	sum := 3
	arr := make([]int, n)
	arr = []int{0, 2, 3, 5, 6, 8, 10, 7}
	//arr = []int{0, 1, 1, 1}
	qb := make([][]int, n+1)
	for i := range qb {
		qb[i] = make([]int, sum+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < sum; j++ {
			qb[i][j] = 1
		}
	}
	fmt.Println("Output=>", subSetSum_Bottomup(qb, arr, sum, n))
}

func subSetSum_Bottomup(qb [][]int, arr []int, sum, n int) int {
	//1 stands for true. Initializing first column.
	for i := 0; i <= n; i++ {
		qb[i][0] = 1
	}

	//0 stands for false. Initializing first row.
	for i := 1; i <= sum; i++ {
		qb[0][i] = 0
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < sum+1; j++ {
			if arr[i-1] > j {
				qb[i][j] = qb[i-1][j]
			} else {
				excluding := qb[i-1][j]
				including := qb[i-1][j-arr[i-1]]
				//Solution is similar to subsetsum problem difference is in datatype of qb. qb is of type int here.
				//And another difference is instead of performing OR operation below we're storing addition of it.
				qb[i][j] = excluding + including
			}

		}
	}
	//TODO: If 0 is at the start then output will be 0. Not all cases but edge cases like, for arr=[]int{0, 1, 1, 1}.  Sorting in reverse order.
	return qb[n][sum]
}
