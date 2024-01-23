/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 */

// @lc code=start

// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one. You must implement a solution with O(n) time and O(1) space

// XOR^: XOR all numbers in `nums` together, the result is the single number. Note that for any given number `num`, `num^0=num`; `num^num=0`. And that XOR has Commutativity and Associativity, ie where there are multiple numbers x, y, z... being XOR, the order of operations does not matter. Consider [4 ^ 1 ^ 2 ^ 1 ^ 2] = [4 ^ (1 ^ 1) ^ (2 ^ 2)] = [4 ^ 0 ^ 0] = 4

func singleNumber(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans ^= num
	}
	return ans
}

// @lc code=end

