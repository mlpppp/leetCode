/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */

// @lc code=start

// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i]. The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer. You must write an algorithm that runs in O(n) time, in O(1) extra space except the output array, and without using the division operation.

// the result in output[i] is prod(nums[:i])*prod(nums[i+1:]). Each prod at all positions can be computed in one traverse in O(n). In the first pass (right to left), store prod(nums[:i]) for all i to the output array, then the second pass (left to right), multiple output[i] by the prod(nums[i+1:]) for each i.

func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))
	// right to left pass (prod to the right of nums[i])
	prod := 1
	output[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		prod = prod * nums[i+1]
		output[i] = prod
	}

	// left to right pass (prod to the left of nums[i])
	prod = 1
	for i := 1; i < len(nums); i++ {
		prod = prod * nums[i-1]
		output[i] = prod * output[i]
	}
	return output
}

// @lc code=end

