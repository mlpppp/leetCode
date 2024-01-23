/*
 * @lc app=leetcode id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 */

// @lc code=start

// Given an array of positive integers nums and a positive integer target, return the minimal length of a subarray whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.

// slide window. Expand the window and add to the sum. Shrink (decrease the sum) while the sum is greater than the target. Record the subarray before shrinking.

func minSubArrayLen(target int, nums []int) int {
	minLength := math.MaxInt
	curSum := 0
	left, right := 0, 0
	for right < len(nums) {
		curSum += nums[right]
		right++
		for curSum >= target {
			minLength = min(minLength, right-left)
			curSum -= nums[left]
			left++
		}
	}
	if minLength == math.MaxInt { // fail to find any subarray
		return 0
	}
	return minLength
}

// @lc code=end

