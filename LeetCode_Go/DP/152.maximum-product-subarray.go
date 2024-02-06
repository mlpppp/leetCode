/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 */

// @lc code=start
// Given an integer array nums, find a subarray that has the largest product, and return the product.The test cases are generated so that the answer will fit in a 32-bit integer.

// subproblem: positiveMaxProd(nums[i:]), find the maximum prod subarray that start with nums[i]. Goal: positiveMaxProd(nums[:]). Use two DP table: positiveMaxProd, negativeMaxProd, return max(positiveMaxProd)

// transition: positiveMaxProd(nums[i:]) = {
// 	if nums[i] < 0 {
// 		negativeMaxProd[i+1]* nums[i]
// 		update: {
// 			negativeMaxProd[i] = positiveMaxProd[i+1]* nums[i]
// 			positiveMaxProd[i] = max(1, negativeMaxProd[i+1]* nums[i])
// 		}
// 	}
// 	if nums[i] > 0 {
// 		positiveMaxProd[i+1]* nums[i]
// 		update {
// 			negativeMaxProd[i] = negativeMaxProd[i+1]* nums[i]
// 			positiveMaxProd[i] = positiveMaxProd[i+1]* nums[i]
// 		}
// 	}
// 	if nums[i] = 0 {
// 		0,
// 		update: {
// 			negativeMaxProd[i] = 1
// 			positiveMaxProd[i] = 1
// 		}
// 	}
// }

// func maxProduct(nums []int) int {
// 	posProd, negProd, ans := make([]int, len(nums)), make([]int, len(nums)), make([]int, len(nums))
// 	// base cases
// 	last := len(nums) - 1
// 	ans[last] = nums[last]
// 	if nums[last] > 0 {
// 		posProd[last] = nums[last]
// 		negProd[last] = 1
// 	} else if nums[last] < 0 {
// 		posProd[last] = 1
// 		negProd[last] = nums[last]
// 	} else {
// 		posProd[last] = 1
// 		negProd[last] = 1
// 	}

// 	for i := len(nums) - 2; i >= 0; i-- {
// 		if nums[i] > 0 {
// 			ans[i] = posProd[i+1] * nums[i]
// 			negProd[i] = negProd[i+1] * nums[i]
// 			posProd[i] = posProd[i+1] * nums[i]
// 		} else if nums[i] < 0 {
// 			ans[i] = negProd[i+1] * nums[i]
// 			negProd[i] = posProd[i+1] * nums[i]
// 			posProd[i] = max(1, ans[i])
// 		} else if nums[i] == 0 {
// 			ans[i] = 0
// 			negProd[i] = 1
// 			posProd[i] = 1
// 		}
// 	}
// 	maxAns := -11
// 	for _, num := range ans {
// 		maxAns = max(maxAns, num)
// 	}
// 	return maxAns
// }

// space optimized
func maxProduct(nums []int) int {
	var posProd, negProd, ans int
	// base cases
	last := len(nums) - 1
	ans = nums[last]
	if nums[last] > 0 {
		posProd = nums[last]
		negProd = 1
	} else if nums[last] < 0 {
		posProd = 1
		negProd = nums[last]
	} else {
		posProd = 1
		negProd = 1
	}

	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] > 0 {
			ans = max(ans, posProd*nums[i])
			negProd = negProd * nums[i]
			posProd = posProd * nums[i]
		} else if nums[i] < 0 {
			temp := negProd * nums[i]
			negProd = posProd * nums[i]
			posProd = max(1, temp)
			ans = max(ans, temp)
		} else if nums[i] == 0 {
			ans = max(ans, 0)
			negProd = 1
			posProd = 1
		}
	}
	return ans
}

// @lc code=end

