/*
 * @lc app=leetcode id=342 lang=golang
 *
 * [342] Power of Four
 */

// @lc code=start

// Given an integer n, return true if it is a power of four. Otherwise, return false. An integer n is a power of four, if there exists an integer x such that n == 4^x.

// an int is a power of 4 if: 1. it is positive, 2. there are only one '1' in the binary representation. 3. there '1' bit must be in the [2,4,6...] position (0 based)

// func isPowerOfFour(n int) bool {
// 	if n <= 0 {
// 		return false
// 	}
// 	pos := 0
// 	for n != 0 {
// 		if n&1 == 1 {
// 			if pos%2 == 0 { // 3. correct position
// 				if n&n-1 == 0 { // 2. only one '1' bit
// 					return true
// 				} else {
// 					return false
// 				}
// 			} else {
// 				return false
// 			}
// 		}
// 		pos += 1
// 		n >>= 1
// 	}
// 	return false
// }

// one line sols: (n-1)%3 == 0 if true only for power of four, not power of two
func isPowerOfFour(n int) bool {
	return n > 0 && (n&(n-1)) == 0 && (n-1)%3 == 0
}

// @lc code=end

