/*
 * @lc app=leetcode id=190 lang=golang
 *
 * [190] Reverse Bits
 */

// @lc code=start

// Reverse bits of a given 32 bits unsigned integer. (n = 00000010100101000001111010011100; Output: 964176192 (00111001011110000010100101000000))

// https://www.youtube.com/watch?v=UcoN6UjAI64

// scan num with rightshift(>>) for 1 bit, while count the position (nth, 0 based). When a 1 bit is found, set the (31-n)th position in answer with 1.
func reverseBits(num uint32) uint32 {
	var ans uint32 = 0
	pos := 0
	for num != 0 {
		// found 1, set the ans' bit
		if num&1 == 1 {
			ans |= 1 << (31 - pos)
		}
		// shift one position
		pos++
		num >>= 1
	}
	return ans
}

// @lc code=end

