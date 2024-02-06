/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */

// @lc code=start

// Given two integer `m` and `n` that represent a m*n matrix, a robot at matrix[0][0] tries to go to matrix[m-1][n-1]. The robot can only move either down or right at any point in time. Return  number of possible unique paths that the robot can take.

// Subproblem: uniquePathsAt(i, j). Goal: uniquePathsAt(0, 0).
// transition: uniquePathsAt(i, j) = {
	uniquePathsAt(i+1, j), go down +
	uniquePathsAt(i, j+1), go right
}

// dp table: 2d table with padding 1 column padding to the right, 1 row padding at the bottom, all fills with 0. make sure dp[m-1][n-1] is 1. Iterate from right to left, bottom to top.

// iterative solution
func uniquePaths(m int, n int) int {
	// dp table with padding filled with 0
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[m-1][n] = 1 //make sure dp[m-1][n-1] is 1.

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			dp[i][j] = dp[i+1][j] + dp[i][j+1]
		}
	}
	return dp[0][0]
}

// @lc code=end

