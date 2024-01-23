/*
 * @lc app=leetcode id=153 lang=golang
 *
 * [153] Find Minimum in Rotated Sorted Array
 */

// @lc code=start

// Define `rotate` of array: pop array to the end and append the result to the head ([0,1,2,4,5,6,7] -> [7,0,1,2,4,5,6]). Suppose an sorted unique-element-array of length n is rotated between [1, n] times, return the minimum element of this array in O(logn) time.

// Binary search: when the nums is unaltered(ie. rotated n times), there is one sorted array, and nums[i] < num[len(num)-1] for every i. otherwise there are two sorted subarrays, and every element in the first subarray is greater than the element in the second subarray. `We want to move right if we are at the first subarray, and want to move left to find the left bound if we are at the second array`. Binary search with criteria: if the element is greater than num[len(num)-1], we search to the right, if the element is smaller or equals to num[len(num)-1], we search to the left for the left bound.

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	result := nums[len(nums)-1]
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > nums[len(nums)-1] {
			left = mid + 1
		} else if nums[mid] <= nums[len(nums)-1] {
			result = nums[mid]
			right = mid - 1
		}
	}
	return result
}

// @lc code=end

