/*
 * @lc app=leetcode id=191 lang=golang
 *
 * [191] Number of 1 Bits
 */

// @lc code=start

// Write a function that takes the binary representation of an unsigned integer and returns the number of '1' bits it has (also known as the Hamming weight).

// same as 338.counting-bits. just beware the difference between uint and int.

func hammingWeight(num uint32) int {
	var count uint32
	for num != 0 {
		count += num & uint32(1)
		num >>= 1
	}
	return int(count)
}

// @lc code=end

