/*
 * @lc app=leetcode id=494 lang=golang
 *
 * [494] Target Sum
 */

// @lc code=start

// Given an integer array, creating expressions by adding '+' or '-' before each integer and concatenating them. Find the number of different expressions that evaluate to the a integer `target`.

// for each number, the choice is either "+" or "-", so brute force is O(2^n)

// subproblem: targetSum(nums[i:], target) int: given a subarray nums[i:], to choose from, find number of expressions that evaluate to target sum. target: targetSum(nums[:], target)

// transition: targetSum(nums[i:], target) = {
// 	targetSum(nums[i+1:], target + nums[i]), if put "-" to nums[i]
// 	+ targetSum(nums[i+1:], target - nums[i]), if put "+" to nums[i]
// }

// Use hashmap instead of 2-D matrix as DP table {key:[i, target], value: targetSum(nums[i:], target)}. It is not efficient to use an 2d matrix as DP tables this time because the second dimension `target's` size is 2*sum(abs(nums[i])).

// base case: target is meanless until decided all numbers symbols. so base case is when i == len(nums){
// 	if target == 0: return 1, found a target sum expression
// 	else: return 0, expression failed the target sum
// }

func findTargetSumWays(nums []int, target int) int {
	table := make(map[[2]int]int)
	return targetSum(nums, 0, target, table)
}

func targetSum(nums []int, i, target int, table map[[2]int]int) int {
	// base case
	if i == len(nums) {
		if target == 0 {
			return 1
		} else {
			return 0
		}
	}
	// query DP table
	if val, ok := table[[2]int{i, target}]; ok {
		return val
	}
	// recursion
	result := targetSum(nums, i+1, target+nums[i], table) +
		targetSum(nums, i+1, target-nums[i], table)
	table[[2]int{i, target}] = result
	return result
}

// @lc code=end

