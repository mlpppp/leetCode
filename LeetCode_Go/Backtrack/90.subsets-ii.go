/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 */

// @lc code=start

// Given an integer array nums that may contain duplicates, return all possible subsets (the power set).

// subset (78.subsets) with further prune: when iterating the choiceSet, if a distinct number has been seen, skip it. We can sort nums first, and check if nums[i] == nums[i-1] (only if 'i-1' falls in [startIdx:]), if so, skip to next.

func subsetsWithDup(nums []int) [][]int {
	slices.Sort(nums)
	ans := [][]int{}
	backtracking(0, []int{}, nums, &ans)
	return ans
}

func backtracking(startIdx int, subset, nums []int, ans *[][]int) {
	// save to ans
	temp := make([]int, len(subset))
	copy(temp, subset)
	*ans = append(*ans, temp)

	// iterate choiceSet
	for i := startIdx; i < len(nums); i++ {
		// prune: if the number is duplicated, skip it
		if i > startIdx && nums[i] == nums[i-1] {
			continue
		}

		// make choice
		subset = append(subset, nums[i])
		// fmt.Printf("%v is calling backtracking(%v)\n", subset, i+1)
		backtracking(i+1, subset, nums, ans)

		// undo choice
		subset = subset[:len(subset)-1]
	}
}

// @lc code=end

