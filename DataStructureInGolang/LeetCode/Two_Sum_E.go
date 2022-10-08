package main

import (
	"fmt"
)

//Problem: https://leetcode.com/problems/two-sum/
/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.
*/

/*
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Input: nums = [3,2,4], target = 6
Output: [1,2]
*/

/*
Runtime: 9 ms, faster than 75.38% of Go online submissions for Two Sum.
Memory Usage: 5.7 MB, less than 6.40% of Go online submissions for Two Sum.
*/
func twoSum(nums []int, target int) []int {
	// For storing output.
	result := make([]int, 2)
	mapObject := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		// If we have seen the current element before
		// It means we have already encountered the other number of the pair
		value, ok := mapObject[nums[i]]
		if ok {
			// Index of the current element
			result[0] = i
			// Index of the other element of the pair
			result[1] = value

			// Else part if we have not seen the current before
			// It means we have not yet encountered any number of the pair
		} else {
			//Save the difference of the target and the current element
			// with the index of the current element
			mapObject[target-nums[i]] = i
		}
	}
	return result
}

func main() {
	arr := make([]int, 4)
	arr = []int{2, 7, 11, 15}
	target := 9
	fmt.Println("Two-Sum-E: ", twoSum(arr, target))
}
