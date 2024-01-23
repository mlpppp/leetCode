/*
 * @lc app=leetcode id=461 lang=golang
 *
 * [461] Hamming Distance
 */

// @lc code=start

// The Hamming distance between two integers is the number of positions at which the corresponding bits are different. Given two integers x and y, return the Hamming distance between them.

// XOR operation (^) returns a sequence of bits where the different bits position is marked "1". Conduct XOR then count the bits in the result.

func hammingDistance(x int, y int) int {
	diffSeq := x ^ y
	count := 0
	for diffSeq != 0 {
		count += diffSeq & 1
		diffSeq >>= 1
	}
	return count
}

// @lc code=end

