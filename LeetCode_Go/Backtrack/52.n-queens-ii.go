/*
 * @lc app=leetcode id=52 lang=golang
 *
 * [52] N-Queens II
 */

// @lc code=start

// The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other. Given an integer n, return the number of distinct solutions to the n-queens puzzle.

// same problem as 51.n-queens

func totalNQueens(n int) int {
	ans := 0
	backtracking([]int{}, n, &ans)
	return ans
}

func backtracking(curSols []int, n int, ans *int) {
	// exit with a new solution
	if len(curSols) == n {
		*ans += 1
		return
	}

	// find available columns
	curRow := len(curSols) // the row we are current working with
	for curCol := 0; curCol < n; curCol++ {
		// skip unvaild cols (use col and used diagonals)
		isUnvaild := false
		for solRow, solCol := range curSols {
			if curCol == solCol || curCol == solCol+(curRow-solRow) || curCol == solCol-(curRow-solRow) {
				isUnvaild = true
				break
			}
		}
		if isUnvaild {
			continue
		}

		// do choice
		curSols = append(curSols, curCol)
		backtracking(curSols, n, ans)
		// undo choice
		curSols = curSols[:len(curSols)-1]
	}
}

// @lc code=end

