/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */

// @lc code=start

// Determine if a 9 x 9 Sudoku board is valid. A vaild Sudoku is: 1. Each row must contain the digits 1-9 without repetition. 2. Each column must contain the digits 1-9 without repetition. 3. Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition. (A Sudoku board (partially filled) could be valid but is not necessarily solvable.)

// Iterate the matrix from top to bottom, left to right. Use a hashmap of hashmap for each rule. (row & col (map[int]map[byte]bool), box (map[[2]int]map[byte]bool))

func isValidSudoku(board [][]byte) bool {
	rowRule := make(map[int]map[byte]bool)
	colRule := make(map[int]map[byte]bool)
	boxRule := make(map[[2]int]map[byte]bool)

	// init maps
	for i := 0; i < len(board); i++ {
		rowRule[i] = make(map[byte]bool)
		colRule[i] = make(map[byte]bool)
	}
	for i := 0; i < len(board)/3; i++ {
		for j := 0; j < len(board)/3; j++ {
			boxRule[[2]int{i, j}] = make(map[byte]bool)
		}
	}

	// scan matrix
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			// skip unfilled
			if board[i][j] == '.' {
				continue
			}
			curNum := board[i][j]
			// check row rule
			if _, exists := rowRule[i][curNum]; exists {
				return false
			} else {
				rowRule[i][curNum] = true
			}
			// check col rule
			if _, exists := colRule[j][curNum]; exists {
				return false
			} else {
				colRule[j][curNum] = true
			}
			// check box rule
			if _, exists := boxRule[[2]int{i / 3, j / 3}][curNum]; exists {
				return false
			} else {
				boxRule[[2]int{i / 3, j / 3}][curNum] = true
			}
		}
	}
	return true
}

// @lc code=end

