/*
 * @lc app=leetcode id=260 lang=golang
 *
 * [260] Single Number III
 */

// @lc code=start

// Given an integer array nums, in which exactly two elements appear only once and all the other elements appear exactly twice. Find the two elements that appear only once. You can return the answer in any order. Solve in O(n) time and O(1) spce.

// Once again, we need to use XOR to solve this problem. But this time, we need to do it in two passes:
// In the first pass, we XOR all elements in the array, and get the XOR of the two numbers we need to find. Note that since the two numbers are distinct, so there must be a set bit (that is, the bit with value '1') in the XOR result. Find out an arbitrary set bit (for example, the rightmost set bit).
// In the second pass, we divide all numbers into two groups, one with the aforementioned bit set, another with the aforementinoed bit unset. Two different numbers we need to find must fall into thte two distrinct groups. XOR numbers in each group, we can find a number in either group. ref
func singleNumber(nums []int) []int {
	// first pass: XOR all numbers and get x1 ^ x2
	xor := 0
	for _, n := range nums {
		xor ^= n
	}

	// create a mask from xor to separate nums into two groups each contains one of x1 and x2
	pos := 0
	for xor&1 != 1 {
		xor >>= 1
		pos++
	}
	mask := 1 << pos

	// second pass: find the two numbers
	x1, x2 := 0, 0
	for _, n := range nums {
		if n&mask == 0 { // zero group
			x1 ^= n
		} else if n&mask != 0 { // one group
			x2 ^= n
		}
	}
	return []int{x1, x2}
}

// @lc code=end

