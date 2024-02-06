/*
 * @lc app=leetcode id=73 lang=golang
 *
 * [73] Set Matrix Zeroes
 */

// @lc code=start

// Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's. Could you devise a constant space solution?

// we need to know all the position of 0 (1 scan to matrix) before filling 0s, otherwise some 0's col/row will be missed.
// O(mn): use a copy of the input as reference
// O(m+n) space: save cols/rows that have 0 occurs, and need to be fill 0 (array or hashmap): zeroCol, zeroRow.
// O(1) space: when we find the first 0, says (i, j), use its column j as zeroCol, and its row i as zeroRow (they will be filled with 0 at last). No need to reinitialize, use 0 to mark the col/row need to be filled 0.

func setZeroes(matrix [][]int) {
	zeroRow, zeroCol := -1, -1
	hasZero := false
	// mark row/col that need to be filled 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				// the first 0, mark it row/col as the recording row/col
				if zeroRow < 0 {
					zeroRow = i
					zeroCol = j
					hasZero = true
					continue
				} else {
					// mark current row i
					matrix[i][zeroCol] = 0
					// mark current col j
					matrix[zeroRow][j] = 0
				}
			}
		}
	}
	// if there is no 0 in matrix, return
	if !hasZero {
		return
	}
	// fill marked rows with 0s
	for i := 0; i < len(matrix); i++ {
		if matrix[i][zeroCol] == 0 {
			// skip the marking row
			if i == zeroRow {
				continue
			}
			// fill marked row with 0
			for j := 0; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}
	// fill marked cols with 0s
	for j := 0; j < len(matrix[0]); j++ {
		if matrix[zeroRow][j] == 0 {
			// skip the marking row
			if j == zeroCol {
				continue
			}
			// fill marked row with 0
			for i := 0; i < len(matrix); i++ {
				matrix[i][j] = 0
			}
		}
	}
	// fill marking row/col with 0s
	for i := 0; i < len(matrix); i++ {
		matrix[i][zeroCol] = 0
	}
	for j := 0; j < len(matrix[0]); j++ {
		matrix[zeroRow][j] = 0
	}
}

// @lc code=end

