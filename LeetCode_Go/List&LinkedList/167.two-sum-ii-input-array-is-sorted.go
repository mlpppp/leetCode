/*
 * @lc app=leetcode id=167 lang=golang
 *
 * [167] Two Sum II - Input Array Is Sorted
 */

// @lc code=start

// Given an sorted array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to target. Assume that each input would have exactly one solution. Elements in the array may not be unique.

// two pointers from both end, if sum larger then target, shrink the right, else shrink the left.

func twoSum(numbers []int, target int) []int {
	result := make([]int, 2)
	pLeft, pRight := 0, len(numbers)-1
	for pLeft <= pRight {
		if numbers[pLeft]+numbers[pRight] < target {
			pLeft++
		} else if numbers[pLeft]+numbers[pRight] > target {
			pRight--
		} else {
			result[0] = pLeft + 1
			result[1] = pRight + 1
			break
		}
	}
	return result
}

// @lc code=end

