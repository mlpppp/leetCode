/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start

// Given an integer array nums, find the subarray with the largest sum, and return its sum.

// subproblem: maxSumStartWith(int i) int: find the subarray with the largest sum, AND must starting with nums[i]. Solve the problem from the i = end to i = 0.

// DP table: nums = [-2,1,-3,4,-1,2,1,-5,4], table = [2,4,3,6,2,3,1,-1,4]

// transition: maxSumStartWith(i) = {
// 	if maxSumStartWith(i+1) > 0, return nums[i] + maxSumStartWith(i+1)
// 	if maxSumStartWith(i+1) <= 0, return nums[i]
// }

// pseudocode: start
// nums = [-2,1,-3,4,-1,2,1,-5,4]
// table = [0] of length nums
// table[len(nums)-1] = nums[len(nums)-1]
// for i = len(table) - 2; i >= 0; i-- {
// 	if table[i+1] > 0 {
// 		table[i] =  nums[i] + table[i+1]
// 	} else {
// 		table[i] =  nums[i]
// 	}
// }
// return max(table)

func maxSubArray(nums []int) int {
	table := make([]int, len(nums))
	table[len(nums)-1] = nums[len(nums)-1]
	for i := len(table) - 2; i >= 0; i-- {
		if table[i+1] > 0 {
			table[i] = nums[i] + table[i+1]
		} else {
			table[i] = nums[i]
		}
	}
	// fmt.Printf("%v\n", table)
	maxSum := math.MinInt64
	for i := 0; i < len(table); i++ {
		if table[i] > maxSum {
			maxSum = table[i]
		}
	}
	return maxSum
}

// @lc code=end

