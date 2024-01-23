/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start

// Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

// Same as binary search, either right-open or both-closed implementation. Only difference is that when nums[mid] == target, it is possible that mid is the boundary or not. We save the mid position can continue binary search to the left(for left boundary), or to the right(for right boundary). When binary search in the next iteration terminates, the currently saved mid position is the boundary.

func searchRange(nums []int, target int) []int {
	return []int{
		binarySearch(nums, target, true),
		binarySearch(nums, target, false),
	}
}

// right-open implementation
// func binarySearch(nums []int, target int, isSearchLeft bool) int {
// 	left, right := 0, len(nums)
// 	found := -1
// 	for left < right {
// 		mid := left + (right-left)/2
// 		if nums[mid] < target {
// 			left = mid + 1 // search in right
// 		} else if nums[mid] > target {
// 			right = mid // search in left
// 		} else if nums[mid] == target {
// 			found = mid
// 			if isSearchLeft {
// 				right = mid // search in left
// 			} else {
// 				left = mid + 1
// 			}
// 		}
// 	}
// 	return found
// }

// both-closed implementation
func binarySearch(nums []int, target int, isSearchLeft bool) int {
	left, right := 0, len(nums)-1
	found := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1 // search in right
		} else if nums[mid] > target {
			right = mid - 1 // search in left
		} else if nums[mid] == target {
			found = mid
			if isSearchLeft {
				right = mid - 1 // search in left
			} else {
				left = mid + 1
			}
		}
	}
	return found
}

// @lc code=end

