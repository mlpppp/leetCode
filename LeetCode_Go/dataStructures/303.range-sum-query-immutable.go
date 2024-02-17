/*
 * @lc app=leetcode id=303 lang=golang
 *
 * [303] Range Sum Query - Immutable
 */

// @lc code=start

// Implement a data structure that store an static array `nums`, and able to query the  sum of the elements of nums between indices `left` and `right` inclusive in O(1)

// NumArray(int[] nums) Initializes the object with the integer array nums.
// int sumRange(int left, int right) Returns the sum of the elements of nums between indices left and right inclusive (i.e. nums[left] + nums[left + 1] + ... + nums[right]).

// PrefixSum: O(1), use prefixSum helper array of size n+1, one 0 padding to the left. prefixSum[i] = sum(nums[:i+1]). SumRange(left, right) returns prefixSum[right] - prefixSum[left-1]

type NumArray struct {
	PreSum []int
}

func Constructor(nums []int) NumArray {
	preSum := make([]int, len(nums))
	sum := 0
	for i, num := range nums {
		sum += num
		preSum[i] = sum
	}
	return NumArray{
		PreSum: preSum,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.PreSum[right]
	} else {
		return this.PreSum[right] - this.PreSum[left-1]
	}
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
// @lc code=end

