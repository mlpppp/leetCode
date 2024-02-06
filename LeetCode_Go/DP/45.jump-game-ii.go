/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */

// @lc code=start

// You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position. Return the minimum number of jumps to reach last index nums[n - 1]. The test cases are generated such that you can reach nums[n - 1].

// DP solution: O(n^2): jump(nums[i:]), goal: jump(nums[:])
// transition: jump(nums[i:]) = {
// 	return 1 + min(dp[i+jump] for jump in range (0, nums[i]))
// }

// dp: 1d len(nums), dp[-1] = 0

// func jump(nums []int) int {
// 	dp := make([]int, len(nums))
// 	dp[len(nums)-1] = 0
// 	for i := len(nums) - 2; i >= 0; i-- {
// 		minStep := math.MaxInt - 1
// 		for jump := nums[i]; jump > 0; jump-- {
// 			// i can go to nums[-1] in one jump
// 			if i+jump >= len(nums)-1 {
// 				minStep = 0
// 				break
// 			}
// 			minStep = min(minStep, dp[i+jump])
// 		}
// 		dp[i] = minStep + 1
// 	}
// 	return dp[0]
// }

// Greedy solution O(n): solve in one pass. Delay the decision of choosing which jump to take until we reach the last possible jump (nums[i]), before that, we update the the maximum possible jump for the next step.
// https://labuladong.github.io/algo/di-er-zhan-a01c6/tan-xin-le-9bedf/ru-he-yun--48a7c/#jump-game-ii

func jump(nums []int) int {
	curBoundary, nextBoundary, steps := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		// find the maximum boundary next step can take to
		nextBoundary = max(nextBoundary, nums[i]+i)
		// if we reach the boundary of the current step, take one step further
		if i == curBoundary {
			steps++
			curBoundary = nextBoundary
		}
	}
	return steps
}

// @lc code=end

