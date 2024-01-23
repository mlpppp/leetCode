/*
 * @lc app=leetcode id=476 lang=golang
 *
 * [476] Number Complement
 */

// @lc code=start
// The complement of an positive integer is the integer you get when you flip all the 0's to 1's and all the 1's to 0's in its binary representation (no leading zero/one bits). Given an integer num, return its complement.  (eg. 5 is "101",  complement is 2, "010")

// num is positive, suppose that excluding the leading 0s, the number has n non-leading bits (5, "101" has 3 bits), we want to XOR num with `000...0111`, with exact 3 '1s'. We first find out how many non-leading bits num has (>>), then produce such a bit sequence and XOR it with num.

func findComplement(num int) int {
	i := num
	count := 0
	for i != 0 {
		count += 1
		i >>= 1
	}
	seq := (1 << count) - 1
	return num ^ seq
}

// @lc code=end

