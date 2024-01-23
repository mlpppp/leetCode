/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start
// You are climbing a staircase of n steps. Each time you either climb 1 or 2 steps. How many distinct ways can you climb?

// brute force: each time we take a 1/2 choice, and add it to the sum, until the sum == n. Backtracking: when sum == n, stop and add numPaths by 1

// overlapping subproblems: for n==5, there are 2 ways to climb the first two steps, and they both want climbStairs(3)

// subproblems: climbStairs(n) int: number of ways to climb n steps, target: climbStairs(n)

// base case: {
// 	climbStairs(1) = 1
// 	climbStairs(2) = 2
// }

// transition: climbStairs(n) = {
// 	climbStairs(n-1), choose to climb 1
// 	+ climbStairs(n-2), choose to climb 2
// }, similar to fibonacci sequence

// func climbStairs(n int) int {
// 	table := make([]int, n+1) // use one padding to make for the 1 based n
// 	return climbSubStairs(n, table)
// }
// func climbSubStairs(n int, table []int) int {
// 	// base case:
// 	if n == 1 {
// 		return 1
// 	}
// 	if n == 2 {
// 		return 2
// 	}
// 	// query DP table
// 	if table[n] != 0 {
// 		return table[n]
// 	}
// 	// recursion
// 	table[n] = climbSubStairs(n-1, table) + climbSubStairs(n-2, table)
// 	return table[n]
// }

// iterative: left to right
func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	table := make([]int, n+1) // use one padding to make for the 1 based n
	table[1] = 1
	table[2] = 2
	for i := 3; i <= n; i++ {
		table[i] = table[i-1] + table[i-2]
	}
	return table[n]
}

// @lc code=end

