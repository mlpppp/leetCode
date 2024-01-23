/*
 * @lc app=leetcode id=1838 lang=golang
 *
 * [1838] Frequency of the Most Frequent Element
 */

// @lc code=start

// You are given an integer array nums and an integer k. In one operation, you can choose an element and increment it by 1. Return the maximum possible frequency of a number after performing at most k operations.

// Sort the array before the algorithm. justify why: To maximize the frequency of a number consider all numbers smaller than it (we can only increament a number). And the closer a number is to the number the more priorities. A sorted array naturally has this property.

// brute force: for each unique number in the sorted array, we try to maximize its frequency by increment numbers smaller than the number. And find the largest possible frequency. O(nlogn + n^2) = O(n^2)

// sliding window: a window is vaild when every number in the window, after applying the budget `k`, can become the same number as window[right].
// keep track of `the sum of the window`. A window is vaild when sum(window) + k >= len(window)*window[right]
// When the window is vaild, expand the window. When the window is not vaild, shrink the window. Log result when the window is vaild

func maxFrequency(nums []int, k int) int {
	slices.Sort(nums)
	maxFrequency := 0
	windowSum := 0
	left, right := 0, 0
	for right < len(nums) {
		// expand window, update window state
		windowSum += nums[right]
		right++
		// check vaildity, if not vaild shrink the window
		for !(windowSum+k >= (right-left)*nums[right-1]) {
			// shrink window
			windowSum -= nums[left]
			left++
		}
		maxFrequency = max(maxFrequency, right-left)
	}
	return maxFrequency
}

// @lc code=end

