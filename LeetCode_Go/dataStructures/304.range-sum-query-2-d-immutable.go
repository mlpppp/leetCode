/*
 * @lc app=leetcode id=304 lang=golang
 *
 * [304] Range Sum Query 2D - Immutable
 */

// @lc code=start

// Implement a data structure that store an static matrix `nums`, and able to query the  sum of the cells inside a submatrix defined by (r1, c1), and (r2, c2): the left-top and right-bottom cells of the submatrix in O(1)

// https://labuladong.github.io/algo/data-structure/prefix-sum/

// 2d prefixSum + trick: each cell (r, c) store the sum of cell inside a submatrix defined by (0, 0), and (r, c). Then for any given submatrix defined by (r1, c1), and (r2, c2), the sum is return: prefixSum(r2, c2) - prefixSum(r2, c1-1) - prefixSum(r1-1, c2) + prefixSum(r1-1, c1-1). Careful edgy cases: when r1 or c1, or both equals 0.

type NumMatrix struct {
	PrefixSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	prefixSum := make([][]int, len(matrix))
	for i := range matrix {
		prefixSum[i] = make([]int, len(matrix[0]))
	}

	for i := 0; i < len(matrix); i++ {
		rowSum := 0
		for j := 0; j < len(matrix[0]); j++ {
			rowSum += matrix[i][j]
			if i == 0 {
				prefixSum[i][j] = rowSum
			} else {
				prefixSum[i][j] = prefixSum[i-1][j] + rowSum
			}
		}
	}

	return NumMatrix{PrefixSum: prefixSum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row1 == 0 && col1 == 0 {
		return this.PrefixSum[row2][col2]
	}
	if row1 == 0 {
		return this.PrefixSum[row2][col2] -
			this.PrefixSum[row2][col1-1]
	}
	if col1 == 0 {
		return this.PrefixSum[row2][col2] -
			this.PrefixSum[row1-1][col2]
	}

	return this.PrefixSum[row2][col2] -
		this.PrefixSum[row2][col1-1] -
		this.PrefixSum[row1-1][col2] +
		this.PrefixSum[row1-1][col1-1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
// @lc code=end

