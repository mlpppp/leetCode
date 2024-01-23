/*
 * @lc app=leetcode id=371 lang=golang
 *
 * [371] Sum of Two Integers
 */

// @lc code=start

// Given two integers a and b (-1000 <= a, b <= 1000), return the sum of the two integers without using the operators + and -.

// iterate and add each bit position

func getSum(a int, b int) int {
	ans := 0
	carry := 0
	mask := 1
	for i := 0; i < 64; i++ {
		if mask&a == mask && mask&b == mask { // both set (1)
			if carry == 1 {
				ans |= 1 << i
			}
			carry = 1
			mask <<= 1
		} else if mask&a == 0 && mask&b == 0 { // both unset (0)
			if carry == 1 {
				ans |= 1 << i
			}
			carry = 0
			mask <<= 1
		} else { // single set (1) and single unset (0)
			if carry == 1 {
				carry = 1
			} else {
				ans |= 1 << i
				carry = 0
			}
			mask <<= 1
		}
	}
	return ans
}

// @lc code=end

