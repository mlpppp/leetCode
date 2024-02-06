/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */

// @lc code=start

// Given an m x n matrix, return all elements of the matrix in spiral order.

// Maintain bound to current row/col: (iMin, iMax), (jMin, jMax). Shrink the bound for every col/row traverse according to the rule: left to right: iMin+1 -> top to down: jMax-1 -> right to left: iMax-1 -> down to top: jMin+1. Exit condition: when either iMin > iMax or jMin > jMax.

func spiralOrder(matrix [][]int) []int {
	iMin, iMax := 0, len(matrix)-1
	jMin, jMax := 0, len(matrix[0])-1
	ans := make([]int, len(matrix)*len(matrix[0]))
	ansIdx := 0
	// terminate when we enter the last pass
	for iMin <= iMax && jMin <= jMax {
		// left to right pass
		for j := jMin; j <= jMax; j++ {
			ans[ansIdx] = matrix[iMin][j]
			ansIdx++
		}
		iMin++ // shrink bound
		if iMin > iMax {
			break
		}

		// top to down pass
		for i := iMin; i <= iMax; i++ {
			ans[ansIdx] = matrix[i][jMax]
			ansIdx++
		}
		jMax-- // shrink bound
		if jMin > jMax {
			break
		}

		// right to left pass
		for j := jMax; j >= jMin; j-- {
			ans[ansIdx] = matrix[iMax][j]
			ansIdx++
		}
		iMax--
		if iMin > iMax {
			break
		}

		// down to top pass
		for i := iMax; i >= iMin; i-- {
			ans[ansIdx] = matrix[i][jMin]
			ansIdx++
		}
		jMin++
	}

	return ans
}

// @lc code=end

