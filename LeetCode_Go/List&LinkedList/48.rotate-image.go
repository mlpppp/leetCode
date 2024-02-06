/*
 * @lc app=leetcode id=48 lang=golang
 *
 * [48] Rotate Image
 */

// @lc code=start

// You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise). You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.

// (trick) Rotate symmetrically through the diagonal axis (left-top to bottom-right), the do symmetrically rotate for each row. (if ask to rotate 90 counterclockwise, do symmetrically rotate for each column)

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if i > j {
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		l, r := 0, len(matrix[0])-1
		for l < r {
			matrix[i][l], matrix[i][r] = matrix[i][r], matrix[i][l]
			l++
			r--
		}
	}
}

// @lc code=end

