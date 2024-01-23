/*
 * @lc app=leetcode id=231 lang=golang
 *
 * [231] Power of Two
 */

// @lc code=start

// Given an integer n, return true if it is a power of two. Otherwise, return false.

// if n <= 0, it must not be a power of two. Otherwise, for a positive number, if it is a power of two, there must be only one '1' in the binary number. Either: count number of 1 in it, Or: use n & (n - 1) trick to remove the last '1' and see if n becomes 0.
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return (n & (n - 1)) == 0
}

// @lc code=end

