/*
 * @lc app=leetcode id=59 lang=golang
 *
 * [59] Spiral Matrix II
 */

// @lc code=start

// Given a positive integer n, generate an n x n matrix filled with elements from 1 to n^2 in spiral order.

// ref 54.spiral-matrix ==> Maintain bound to current row/col: (iMin, iMax), (jMin, jMax). Shrink the bound for every col/row traverse combination according to the rule: left to right: iMin+1 -> top to down: jMax-1 -> right to left: iMax-1 -> down to top: jMin+1. Exit condition: when either iMin > iMax or jMin > jMax.

func generateMatrix(n int) [][]int {
	// init matrix
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	// spiral traverse
	val := 1
	iMin, iMax := 0, n-1
	jMin, jMax := 0, n-1
	// terminate when we enter the last pass
	for iMin <= iMax && jMin <= jMax {
		// left to right pass
		for j := jMin; j <= jMax; j++ {
			matrix[iMin][j] = val
			val++
		}
		iMin++ // shrink bound
		if iMin > iMax {
			break
		}

		// top to down pass
		for i := iMin; i <= iMax; i++ {
			matrix[i][jMax] = val
			val++
		}
		jMax-- // shrink bound
		if jMin > jMax {
			break
		}

		// right to left pass
		for j := jMax; j >= jMin; j-- {
			matrix[iMax][j] = val
			val++
		}
		iMax--
		if iMin > iMax {
			break
		}

		// down to top pass
		for i := iMax; i >= iMin; i-- {
			matrix[i][jMin] = val
			val++
		}
		jMin++
	}
	return matrix
}

// @lc code=end

