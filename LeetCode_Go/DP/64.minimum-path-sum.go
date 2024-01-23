/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 */

// @lc code=start

// Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right, which minimizes the sum of all numbers along its path. You can only move either down or right.

// brute force: essentially in every position You have to face two choices, so O(2^(mn)). Since there are multiple ways to reach a position, there are overlapping subproblems.

// subproblem: minPathSumFrom(grid[i][j]) int: minimum path sum starting from position grid[i][j]. Use 2-D DP table for i and j. Goal: grid[0][0]

// transition: minPathSumFrom(grid[i][j]) = {
// 	grid[i][j] + min(minPathSumFrom(grid[i+1][j]), minPathSumFrom(grid[i][j+1])), cost of current position + min cost of either of the decisions
// }

// base case: reaching the bottom right. / i or j out of bound {
// 	if reaching the bottom right: return grid[m-1][n-1], value of the bottom right
// 	if i or j out of bound: return maxInt, indicates it is a offlimit position
// }

// func minPathSum(grid [][]int) int {
// 	// DP table: [m][n]int initalize with -1 (sum must not be negative as 0 <= grid[i][j] )
// 	table := make([][]int, len(grid))
// 	for i, _ := range table {
// 		table[i] = make([]int, len(grid[0]))
// 		for j, _ := range table[i] {
// 			table[i][j] = -1
// 		}
// 	}
// 	return minPathSumFrom(grid, 0, 0, table)
// }

// func minPathSumFrom(grid [][]int, i, j int, table [][]int) int {
// 	// base case: out of range/ reaching goal
// 	if i >= len(grid) || j >= len(grid[0]) {
// 		return math.MaxInt
// 	}
// 	if i == len(grid)-1 && j == len(grid[0])-1 {
// 		table[i][j] = grid[i][j]
// 		return table[i][j]
// 	}
// 	// query DP table
// 	if table[i][j] != -1 {
// 		return table[i][j]
// 	}
// 	// recursion
// 	table[i][j] = grid[i][j] + min(
// 		minPathSumFrom(grid, i+1, j, table),
// 		minPathSumFrom(grid, i, j+1, table))
// 	return table[i][j]
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	} else {
// 		return b
// 	}
// }

// iterative solution: [m+1][n+1]int DP table, rightmost column and bottom row used for padding(fill with maxInt). iterate from bottom to top, right to left

// table[i][j] -> table[i][j+1]
// |
// table[i+1][j]

func minPathSum(grid [][]int) int {
	// DP table: [m+1][n+1]int
	table := make([][]int, len(grid)+1)
	for i, _ := range table {
		table[i] = make([]int, len(grid[0])+1)
	}
	// init table
	for i, _ := range table[len(grid)] { // padding last row
		table[len(grid)][i] = math.MaxInt
	}
	for i := 0; i <= len(grid); i++ { // padding last column
		table[i][len(grid[0])] = math.MaxInt
	}

	// iteration
	for i := len(grid) - 1; i >= 0; i-- { // bottom to top
		for j := len(grid[0]) - 1; j >= 0; j-- { // right to left
			// first solution (base case)
			if i == len(grid)-1 && j == len(grid[0])-1 {
				table[i][j] = grid[i][j]
			} else {
				table[i][j] = grid[i][j] + min(table[i+1][j], table[i][j+1])
			}
		}
	}
	return table[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end

