/*
 * @lc app=leetcode id=338 lang=golang
 *
 * [338] Counting Bits
 */

// @lc code=start

// Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i. (0 <= n <= 10^5)

// use right shift to iterate over the number's bit representation until it becomes 0

func countBits(n int) []int {
	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		count := 0
		num := i
		for num != 0 {
			count += num & 1
			num >>= 1
		}
		ans[i] = count
	}
	return ans
}

// @lc code=end

