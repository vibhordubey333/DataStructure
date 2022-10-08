//Problem: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
/*
Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.
Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.
The tests are generated such that there is exactly one solution. You may not use the same element twice.
Your solution must use only constant extra space.

Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].

Input: numbers = [2,3,4], target = 6
Output: [1,3]
Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].

Input: numbers = [-1,0], target = -1
Output: [1,2]
Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].
*/
package main

func main() {}

/*
Runtime: 17 ms, faster than 68.83% of Go online submissions for Two Sum II - Input Array Is Sorted.
Memory Usage: 6.1 MB, less than 5.45% of Go online submissions for Two Sum II - Input Array Is Sorted.
SC: O(n)
TC: O(n)
*/
func twoSum_HashMapApproach(numbers []int, target int) []int {
	n := len(numbers)
	mapObject := make(map[int]int)
	result := make([]int, n)

	for i := 0; i < n; i++ {
		complement := target - numbers[i]

		val, ok := mapObject[complement]
		if ok {
			result[0] = val
			result[1] = i + 1
			return result
		}
		mapObject[numbers[i]] = i + 1
	}
	return result
}

/*
Runtime: 25 ms, faster than 43.27% of Go online submissions for Two Sum II - Input Array Is Sorted.
Memory Usage: 5.4 MB, less than 56.41% of Go online submissions for Two Sum II - Input Array Is Sorted.
TC: O(nlog(n)))
SP: O(1)
*/
func twoSum_TwoPointerApproach(numbers []int, target int) []int {
	frontPtr := 0
	secondPtr := len(numbers) - 1
	result := make([]int, 2)
	for frontPtr < secondPtr {
		totalSum := numbers[frontPtr] + numbers[secondPtr]

		if totalSum == target {
			break
		} else if totalSum > target {
			secondPtr--
		} else {
			frontPtr++
		}
	}
	result[0], result[1] = frontPtr+1, secondPtr+1
	return result
}

/*
Runtime: 554 ms, faster than 7.77% of Go online submissions for Two Sum II - Input Array Is Sorted.
Memory Usage: 5.4 MB, less than 56.41% of Go online submissions for Two Sum II - Input Array Is Sorted.

TC: O(n^2)
SC: O(n)
*/
func twoSum_BruteForce(numbers []int, target int) []int {

	length := len(numbers)
	result := make([]int, 2)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if numbers[j] == target-numbers[i] {
				result[0] = i + 1
				result[1] = j + 1
				return result
			}
		}
	}
	return result
}
