/*
 * @lc app=leetcode id=37 lang=golang
 *
 * [37] Sudoku Solver
 */

// @lc code=start

// Write a program to solve a Sudoku puzzle. A sudoku solution must satisfy all of the following rules: 1. Each of the digits 1-9 must occur exactly once in each row. 2. Each of the digits 1-9 must occur exactly once in each column. 3. Each of the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.

// backtracking(i, j) solve sudoku at the location (i, j). Goal: reach backtracking(8, 8). Solve from left to right, top to bottom. O(9^m) where m is the number of cells to be filled
// cases: 1. if is filled, go to the next. 2. solve from left to right by default, if reached the end of row, go to the next row.
// choiceSet: only fill vaild number from '1' to '9', use 3 hashmap of hashmap to check vaildity (ref. 36.valid-sudoku)
// early termination: terminate all recursion when the first solution is found (we need only one solution). We can return true when a solution is found, and use "if backtracking() return true" instead of just call backtracking() when making choice

func solveSudoku(board [][]byte) {
	// init vaildity maps
	rowRule := make(map[int]map[byte]bool)
	colRule := make(map[int]map[byte]bool)
	boxRule := make(map[[2]int]map[byte]bool)

	for i := 0; i < 9; i++ {
		rowRule[i] = make(map[byte]bool)
		colRule[i] = make(map[byte]bool)
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			boxRule[[2]int{i, j}] = make(map[byte]bool)
		}
	}
	// set pre-filled rules
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]
			if num != '.' {
				rowRule[i][num] = true
				colRule[j][num] = true
				boxRule[[2]int{i / 3, j / 3}][num] = true
			}
		}
	}

	// start DFS
	backtracking(board, 0, 0, rowRule, colRule, boxRule)
}

func backtracking(board [][]byte, i, j int, rowRule, colRule map[int]map[byte]bool, boxRule map[[2]int]map[byte]bool) bool {
	// exit case: when finished the last cell (both col and row are out of bounds)
	if i == 9 {
		return true
	}

	// if reached the end of row, go to the next row
	if j == 9 {
		return backtracking(board, i+1, 0, rowRule, colRule, boxRule)
	}

	// if filled, go to the next cell
	if board[i][j] != '.' {
		return backtracking(board, i, j+1, rowRule, colRule, boxRule)
	}

	// choice set: filter out invaild numbers
	for num := byte('1'); num <= byte('9'); num++ {
		// prune: skip invaild number
		_, rowExists := rowRule[i][num]
		_, colExists := colRule[j][num]
		_, boxExists := boxRule[[2]int{i / 3, j / 3}][num]
		if rowExists || colExists || boxExists {
			continue
		}

		// do choice: update board and rules
		board[i][j] = num
		rowRule[i][num] = true
		colRule[j][num] = true
		boxRule[[2]int{i / 3, j / 3}][num] = true

		// if a solution is found, cancel any other choices and terminate without undo
		if backtracking(board, i, j+1, rowRule, colRule, boxRule) {
			return true
		}

		// undo choice: update board and rules
		board[i][j] = '.'
		delete(rowRule[i], num)
		delete(colRule[j], num)
		delete(boxRule[[2]int{i / 3, j / 3}], num)
	}
	// no luck this solution
	return false
}

// @lc code=end

