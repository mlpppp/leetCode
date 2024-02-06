/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */

// @lc code=start

// You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position. Return true if you can reach the last index, or false otherwise.

// DP solution: O(n^2): canJump(nums[i:]), goal: canJump(nums[:])
// transition: canJump(nums[i:]) = {
// 	for jump in range (0, nums[i]) {
// 		if dp[i+jump] return true
// 	}
// 	return false
// }

// dp: 1d len(nums), dp[-1] = true

// func canJump(nums []int) bool {
// 	dp := make([]bool, len(nums))
// 	dp[len(nums)-1] = true
// 	for i := len(nums) - 2; i >= 0; i-- {
// 		for jump := nums[i]; jump >= 0; jump-- {
// 			if i+jump >= len(nums)-1 || dp[i+jump] {
// 				dp[i] = true
// 				break
// 			}
// 		}
// 	}
// 	return dp[0]
// }

// Greedy O(n): one scan from right to left, move the goal from the last position to the leftmost position that can reaches the previous goal. If we can move the goal to nums[0], then return true

func canJump(nums []int) bool {
	curGoal := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i]+i >= curGoal {
			curGoal = i
		}
	}
	return curGoal == 0
}

// @lc code=end

