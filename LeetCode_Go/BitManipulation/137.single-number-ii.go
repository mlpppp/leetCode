/*
 * @lc app=leetcode id=137 lang=golang
 *
 * [137] Single Number II
 */

// @lc code=start

// Given an integer array nums where every element appears three times except for one, which appears exactly once. Find the single element and return it. Solve in O(n) time and O(1) space.

// bit manipulation (O(32n) or O(64n)): compare all number in bit representation (similar to 169.majority-element). For each bit position, count number of 1 and 0. if either of 1 or 0 is not a multiplier of 3 (should be count = 3*k + 1, 1 is the single number), it is the answer's bit at the position.

func singleNumber(nums []int) int {
	ans := 0
	mask := 1
	for i := 0; i < 64; i++ {
		// count number of 1s in the position
		count1 := 0
		for _, num := range nums {
			if num&mask == mask {
				count1++
			}
		}
		// if count1 is not multiplier of 3, set the position to 1, otherwise set 0
		if count1%3 != 0 {
			ans |= mask
		}
		// move to the next position
		mask <<= 1
	}
	return ans
}

// @lc code=end

