package main
import(
	"fmt"
)
//Problem Statement: https://leetcode.com/problems/target-sum/

/*
Summary: You are given an integer array nums and an integer target.

You want to build an expression out of nums by adding one of the symbols '+' and '-' before each integer in nums and then concatenate all the integers.

For example, if nums = [2, 1], you can add a '+' before 2 and a '-' before 1 and concatenate them to build the expression "+2-1".
Return the number of different expressions that you can build, which evaluates to target.
*/


func main(){
	arr := make([]int,0)
	arr = []int{1,1,1,1,1} // Output: 5
	target := 3

	fmt.Println(findTargetSumWays(arr,target))

}

//Leetcode starts.
func findTargetSumWays(nums []int, target int) int {
	return computeTargetSum(nums,target,len(nums))
}

func computeTargetSum(nums []int,target,n int) int{

	if target == 0 && n == 0{
		return 1
	}

	if n == 0{
		return 0
	}
	addOp := computeTargetSum(nums,target+nums[n-1],n-1)
	subOp := computeTargetSum(nums,target-nums[n-1],n-1)
	return addOp + subOp

}

