/*
 * @lc app=leetcode id=51 lang=golang
 *
 * [51] N-Queens
 */

// @lc code=start

// The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other. Queen attacks all direction: row,column and diagonal (north, northeast, east, southeast, south, southwest, west, northwest). Given n, return all distinct solutions to the n-queens puzzle (sample solution: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]).

// each row can only have 1 queen. Build solution by adding the queen to a row. curSols: []int store queen's col in each row (index is solution's row, value is solution's col). goal: len(curSol) = n.

// choiceSet: all possible cols in curRow, excludes unavailable col: used columns and diagonals (for all (solRow, solCol) in curSols, [][solCol+i] and [][solCol-i] where i is the computed for the current row, i = diff(solRow, curRow))

func solveNQueens(n int) [][]string {
	ans := [][]int{}
	backtracking([]int{}, n, &ans)
	// format answers
	ansString := make([][]string, len(ans))
	for i, sol := range ans {
		ansString[i] = make([]string, n)
		for j, solRow := range sol {
			dots := []byte(strings.Repeat(".", n))
			dots[solRow] = 'Q'
			ansString[i][j] = string(dots)
		}
	}
	return ansString
}

func backtracking(curSols []int, n int, ans *[][]int) {
	// exit with a new solution
	if len(curSols) == n {
		temp := make([]int, n)
		copy(temp, curSols)
		*ans = append(*ans, temp)
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

