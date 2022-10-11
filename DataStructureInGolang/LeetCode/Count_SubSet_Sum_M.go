//https://leetcode.com/problems/number-of-subsequences-that-satisfy-the-given-sum-condition/
package main

import (
	"sort"
)

/*
You are given an array of integers nums and an integer target.

Return the number of non-empty subsequences of nums such that the sum of the minimum and maximum element on it is less or equal to target.
Since the answer may be too large, return it modulo 109 + 7

Example 1:

Input: nums = [3,5,6,7], target = 9
Output: 4
Explanation: There are 4 subsequences that satisfy the condition.
[3] -> Min value + max value <= target (3 + 3 <= 9)
[3,5] -> (3 + 5 <= 9)
[3,5,6] -> (3 + 6 <= 9)
[3,6] -> (3 + 6 <= 9)

Example 2:

Input: nums = [3,3,6,8], target = 10
Output: 6
Explanation: There are 6 subsequences that satisfy the condition. (nums can have repeated numbers).
[3] , [3] , [3,3], [3,6] , [3,6] , [3,3,6]
*/

/*Runtime: 257 ms, faster than 57.14% of Go online submissions for Number of Subsequences That Satisfy the Given Sum Condition.
Memory Usage: 9.3 MB, less than 14.29% of Go online submissions for Number of Subsequences That Satisfy the Given Sum Condition.
TC: O(nlog(n))
SC: O(n)
*/
func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	const MOD = 1000000007 //1e9 + 7
	n := len(nums)
	exponents := make([]int, n+1)
	exponents[1] = 1
	//Pre-computing power of 2 as it can result into TLE.
	//We will return answer % MOD as answer can be large.
	for i := 2; i < n+1; i++ {
		exponents[i] = (2 * exponents[i-1]) % MOD //Or exponents[i-1] << 1%MOD
	} //Output of exponents: [0 1 2 4 8 16]
	left := 0
	right := n - 1
	count := 0

	for left <= right {
		if nums[left]+nums[right] > target {
			right--
		} else {
			count = (count + exponents[right-left+1]) % MOD
			left++
		}
	}
	return count
}
