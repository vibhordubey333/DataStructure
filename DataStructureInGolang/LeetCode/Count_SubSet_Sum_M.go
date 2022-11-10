//https://leetcode.com/problems/number-of-subsequences-that-satisfy-the-given-sum-condition/
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

/*Runtime: 134 ms, faster than 93.33% of Go online submissions for Number of Subsequences That Satisfy the Given Sum Condition.
Memory Usage: 8.2 MB, less than 100.00% of Go online submissions for Number of Subsequences That Satisfy the Given Sum Condition.
TC: O(nlog(n))
SC: O(n)
*/
package main

import (
	"sort"
)

func numSubseq(nums []int, target int) int {
	//On LHS there will be small no. and on RHS large no.
	sort.Ints(nums)
	const MOD = 1000000007 // Avoid using 1e9 + 7 as it takes more time to execute
	n := len(nums)
	exponents := make([]int, n+1)
	//Setting index 1 value as 1 other values like exponents[0] will already be populated by zero itself.
	exponents[1] = 1
	//Pre-computing power of 2 as it can result into TLE.
	//We will return answer % MOD as answer can be large.
	//Other values are already set initializing loop from i = 2
	for i := 2; i <= n; i++ {
		exponents[i] = (2 * exponents[i-1]) % MOD //Or exponents[i-1] << 1%MOD
	} //Output of exponents: [0 1 2 4 8 16]
	left := 0
	right := n - 1
	count := 0

	for left <= right {
		if nums[left]+nums[right] > target {
			right--
		} else { //Found our sum so increasing count.
			count = (count + exponents[right-left+1]) % MOD
			left++
		}
	}
	return count
}

func main() {}
