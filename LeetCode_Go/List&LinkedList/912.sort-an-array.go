/*
 * @lc app=leetcode id=912 lang=golang
 *
 * [912] Sort an Array
 */

// @lc code=start
func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, i, j int) {
	// base case: empty or single element array
	if i == j || i > j {
		return
	}

	// pre recursion
	target := nums[i]
	pSmall := i + 1
	pEdit := i + 1
	for pSmall <= j {
		// if find one smaller than target, sawp pEdit and pSmall, and move pEdit.
		if nums[pSmall] < target {
			if pSmall != pEdit {
				nums[pSmall], nums[pEdit] = nums[pEdit], nums[pSmall]
			}
			pEdit++
		}
		pSmall++
	}

	// swap pEdit - 1 and target
	nums[i], nums[pEdit-1] = nums[pEdit-1], nums[i]

	// recursion
	quickSort(nums, i, pEdit-2)
	quickSort(nums, pEdit, j)
}

// @lc code=end

