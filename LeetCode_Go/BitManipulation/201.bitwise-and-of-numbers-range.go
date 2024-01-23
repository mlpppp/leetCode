/*
 * @lc app=leetcode id=201 lang=golang
 *
 * [201] Bitwise AND of Numbers Range
 */

// @lc code=start

// Given two integers left and right that represent the range [left, right], return the bitwise AND of all numbers in this range, inclusive. (0 <= left <= right <= 2^31 - 1)

// eg. [5, 7], 5 & 6 & 7
// 5	0101
// 6	0110
// 7	0111
// &    0100 = 4

// Brute force, bitwise & one by one. Beware the right might be the maxInt which may create infinite loops

// func rangeBitwiseAnd(left int, right int) int {
// 	ans := left
// 	for i := left + 1; i <= right && i > 0; i++ {
// 		ans &= i
// 		if ans == 0 {
// 			return ans
// 		}
// 	}
// 	return ans
// }

// Count number of bits `zeroCount` that makes left and right different. In the result, Keep the left the same as left or right, and fill the rest to the right with zeroCount number if 0s

// 26 to 30
// 11010
// 11011
// 11100
// 11101
// 11110
// we want to cut the three last bits between 11010 and 11110 and make them 0, hence it becomes 11000

func rangeBitwiseAnd(left int, right int) int {
	zeroCount := 0 // the last zeroCount bits in the result are all 0s
	for left != right {
		zeroCount++
		left >>= 1
		right >>= 1
	}
	left <<= zeroCount
	return left
}

// @lc code=end

