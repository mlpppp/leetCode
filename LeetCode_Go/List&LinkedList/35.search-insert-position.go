/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start

// Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order, in O(log n) time.

// close to the vanilla binary search, understand the termination position. If a target does not exists, the smallest element that is larger than the target occupies the position target should have occupied.
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := len(nums)
	for left <= right {
		mid := (left + right) / 2
		if target < nums[mid] {
			result = mid // might be the position of target
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		} else if target == nums[mid] {
			return mid
		}
	}
	return result
}

// @lc code=end

