/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 */

// @lc code=start

// given a `sorted matrix`: every row is sorted; all elements of a row are greater than any element of the previous row. Given an integer target, return true if target is in matrix or false otherwise, in O(log(m * n)).

// Run two binary search. first search the first column (first element smaller than) to identify the row (logm), then search the row (existance) to identify the column (logn)

func searchMatrix(matrix [][]int, target int) bool {
	left, right := 0, len(matrix)-1

	// binary search for 'first element smaller than' in rows
	targetRow := -1
	for left <= right {
		mid := (left + right) / 2
		if target < matrix[mid][0] {
			// target must be above, search to the above column
			right = mid - 1
		}
		if target >= matrix[mid][0] {
			// could be the same column or below column
			targetRow = mid
			left = mid + 1
		}
	}

	// edge case: target is smaller than smallest matrix[0][0]
	if targetRow == -1 {
		return false
	}

	// binary search for existance in row
	left, right = 0, len(matrix[0])-1
	for left <= right {
		mid := (left + right) / 2
		if target < matrix[targetRow][mid] {
			right = mid - 1
		} else if target > matrix[targetRow][mid] {
			left = mid + 1
		} else if target == matrix[targetRow][mid] {
			return true
		}
	}
	return false
}

// @lc code=end

