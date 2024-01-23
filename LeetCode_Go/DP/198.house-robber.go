/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */

// @lc code=start

// Given an array of integers `nums` (houses), find a subset(house you rob) that maximize the sum of the subset, with a constraint that no adjacent elements in `nums` can be put into the subset

// brute force: each element have to choose rob->to i+2, or not rob->to i+1. O(2^n).

// exists overlapping subproblem. in nums = [1,2,3,1]. There are multiple ways to reach nums[2], either through nums[0] or nums[1]

// subproblem: robFrom(nums[i:]) int. maximum amount of money you can rob, start from house i to the right. Target: robFrom(nums[0:])

// transition: rob current house, and skip the next. OR dont rob current house, and go to the next house.

// robFrom(nums[i:]) = {
// 	max(nums[i] + robFrom(nums[i+2:]), robFrom(nums[i+1:]))
// }

// base case: i out of range, no house to rob, return 0

// func rob(nums []int) int {
// 	// DP table, init with -1
// 	table := make([]int, len(nums))
// 	for i, _ := range table {
// 		table[i] = -1
// 	}
// 	return robFrom(nums, table, 0)
// }

// func robFrom(nums, table []int, i int) int {
// 	// base case: i out of range
// 	if i >= len(nums) {
// 		return 0
// 	}
// 	// query DP table
// 	if table[i] != -1 {
// 		return table[i]
// 	}
// 	// recursion
// 	table[i] = max(
// 		robFrom(nums, table, i+1),
// 		nums[i]+robFrom(nums, table, i+2))

// 	return table[i]
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	} else {
// 		return b
// 	}
// }

// iterative: right to left. base case: solution of last two houses
func rob(nums []int) int {
	// DP table, filled with 0
	table := make([]int, len(nums))
	// base case
	table[len(nums)-1] = nums[len(nums)-1]
	if len(nums) == 1 { // early return for edgy input
		return table[0]
	}
	table[len(nums)-2] = max(nums[len(nums)-1], nums[len(nums)-2])
	if len(nums) == 2 { // early return for edgy input
		return table[0]
	}

	// iteration
	for i := len(nums) - 3; i >= 0; i-- {
		table[i] = max(nums[i]+table[i+2], table[i+1])
	}
	return table[0]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end

