/*
https://leetcode.com/problems/target-sum/

*/

package main

import (
	"fmt"
)

var findTargetSum_Memoized = findTargetSumWays

func main() {
	n := 5
	//	n := 10
	nums := make([]int, n)
	nums = []int{1, 1, 1, 1, 1}
	//nums = []int{7, 9, 3, 8, 0, 2, 4, 8, 3, 9}
	targetSum := 3
	fmt.Println("TargetSum_Memoized: ", findTargetSum_Memoized(nums, targetSum))
}

/*
Runtime: 10 ms, faster than 84.57% of Go online submissions for Target Sum.
Memory Usage: 8.6 MB, less than 6.79% of Go online submissions for Target Sum.
*/
func findTargetSumWays(nums []int, targetSum int) int {
	numsLength := len(nums)
	totalNumsSum := 0

	//Compute totalsum of elements.
	for _, i := range nums {
		totalNumsSum += i
	}

	qb := make([][]int, numsLength+1)

	for rows := range qb {
		qb[rows] = make([]int, 2*totalNumsSum+1)

		for cols := 0; cols <= 2*totalNumsSum; cols++ {
			qb[rows][cols] = -1
		}
	}

	var computeTargetSum func(int, int) int

	computeTargetSum = func(startIndex, sum int) int {
		if startIndex == numsLength {
			if sum == targetSum {
				return 1
			} else {
				return 0
			}
		}

		if qb[startIndex][totalNumsSum+sum] != -1 {
			return qb[startIndex][totalNumsSum+sum]
		}

		addOp := computeTargetSum(startIndex+1, sum+nums[startIndex])
		subOp := computeTargetSum(startIndex+1, sum-nums[startIndex])

		qb[startIndex][totalNumsSum+sum] = addOp + subOp

		return qb[startIndex][totalNumsSum+sum]
	}

	//Calling function.
	return computeTargetSum(0, 0)

}
