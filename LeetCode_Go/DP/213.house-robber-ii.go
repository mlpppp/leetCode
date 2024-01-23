/*
 * @lc app=leetcode id=213 lang=golang
 *
 * [213] House Robber II
 */

// @lc code=start

// reference [198] House Robber

// Given solution of 198 rob(nums []int) int, this problem is equivalent to max(rob(nums[0:len(nums)-1]), rob(nums[1:len(nums)])). ie either rob the last house or rob the first house.

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(rob1(nums[1:]), rob1(nums[:len(nums)-1]))
}

// solution of 198 House Robber
func rob1(nums []int) int {
	// DP table, filled with 0
	var left, right int = 0, 0

	// iteration
	for i := len(nums) - 1; i >= 0; i-- {
		newLeft := max(nums[i]+right, left)
		left, right = newLeft, left
	}
	return left
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end

