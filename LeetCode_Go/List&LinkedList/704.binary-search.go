/*
 * @lc app=leetcode id=704 lang=golang
 *
 * [704] Binary Search
 */

// @lc code=start

// Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.

// both-closed interval implementation
// func search(nums []int, target int) int {
// 	left, right := 0, len(nums)-1
// 	for left <= right {
// 		mid := left + (right-left)/2
// 		if nums[mid] == target {
// 			return mid
// 		} else if nums[mid] < target {
// 			left = mid + 1
// 		} else if nums[mid] > target {
// 			right = mid - 1
// 		}
// 	}
// 	return -1
// }

// right-open interval implementation
func search(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		}
	}
	return -1
}

// @lc code=end

