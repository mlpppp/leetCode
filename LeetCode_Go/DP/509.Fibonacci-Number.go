/*
 * @lc app=leetcode id=509 lang=golang
 *
 * [509] Fibonacci Number
 */

// @lc code=start

// Given n, calculate Fibonacci numbers F(n).
// F(0) = 0, F(1) = 1
// F(n) = F(n - 1) + F(n - 2), for n > 1.

// O(n) time, O(n) space

// recursion

// func fib(n int) int {
// 	table := make([]int, n+1)
// 	return fibHelper(n, table)
// }

// func fibHelper(n int, table []int) int {
// 	if n <= 1 {
// 		return n
// 	}
// 	if table[n] != 0 {
// 		return table[n]
// 	}
// 	fibn := fibHelper(n-1, table) + fibHelper(n-2, table)
// 	table[n] = fibn
// 	return fibn
// }

// iteration
func fib(n int) int {
	if n <= 1 {
		return n
	}
	table := make([]int, n+1)
	table[0], table[1] = 0, 1
	for i := 2; i < len(table); i++ {
		table[i] = table[i-1] + table[i-2]
	}
	return table[n]
}

// @lc code=end

