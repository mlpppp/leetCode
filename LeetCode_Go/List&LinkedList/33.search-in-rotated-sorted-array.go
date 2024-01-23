/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */

// @lc code=start

// Given a sorted array `nums` with distinct values. It might be rotated at any position ([0,1,2,4,5,6,7] -> [4,5,6,7,0,1,2] when rotated at position 3). Search for the index of an integer value `target` in the rotated array, or return -1 if it does not exist in O(logn) time.

// There are two pivot points in a rotated array: nums[0] or nums[mid] (we can use either of them, in this case choose nums[0]). When target > nums[0], the target must be in the first part, when target < nums[0], the target must be in the second part. If current mid is in the false half a target belongs to, goes the the direction of the opposite part. If current mid is already in the correct part, choose the direction by comparing nums[mid] and target's value.

func search(nums []int, target int) int {
	// identify the part target might belong to
	var isLeftpart bool
	if target > nums[0] {
		isLeftpart = true
	} else if target < nums[0] {
		isLeftpart = false
	} else {
		return 0
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if isLeftpart && nums[mid] < nums[0] { // move mid to leftPart
			right = mid - 1
		} else if !isLeftpart && nums[mid] >= nums[0] { // move mid to rightPart
			left = mid + 1
		} else { // already in the correct part
			if target < nums[mid] {
				right = mid - 1
			} else if target > nums[mid] {
				left = mid + 1
			} else {
				return mid
			}
		}
	}
	return -1
}

// @lc code=end

