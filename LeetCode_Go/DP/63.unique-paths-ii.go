/*
 * @lc app=leetcode id=63 lang=golang
 *
 * [63] Unique Paths II
 */

// @lc code=start

// Given m*n matrix where obstacles and space are marked as 1 or 0 respectively, a robot at matrix[0][0] tries to go to matrix[m-1][n-1]. The robot can only move either down or right at any point in time, and a path that the robot takes cannot include any obstacle. Return number of possible unique paths that the robot can take.

// Subproblem: uniquePathsAt(i, j). Goal: uniquePathsAt(0, 0).
// transition: uniquePathsAt(i, j) = {
// 	if obstacles[i][j] == 1, 0
// 	else:
// 		uniquePathsAt(i+1, j) +
// 		uniquePathsAt(i, j+1)
// }

// dp table: 2d table with padding 1 column padding to the right, 1 row padding at the bottom, all fills with 0. make sure dp[m-1][n-1] is 1. Iterate from right to left, bottom to top.

// iterative solution
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// dp table with padding filled with 0
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[m-1][n] = 1 //make sure dp[m-1][n-1] is 1.

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			dp[i][j] = dp[i+1][j] + dp[i][j+1]
		}
	}
	return dp[0][0]
}

// @lc code=end

