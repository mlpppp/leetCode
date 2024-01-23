/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start

// Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.

// permutation: any possible arrange (n!). Implementation: 1. use usedNums, []bool as the alternative to choiceSet, when compute the current choice, iterate nums and skip numbers that has been used (in the usedNums) 2. make sure that create a deep copy of the curSol

func permute(nums []int) [][]int {
	// init values: ans, curSol and usedNums
	ans := [][]int{}
	curSol := []int{}
	useNums := make([]bool, len(nums))
	// start backtracking
	backtracking(curSol, nums, useNums, &ans)
	return ans
}

func backtracking(curSol []int, nums []int, usedNums []bool, ans *[][]int) {
	if len(curSol) == len(nums) {
		// careful: make a copy
		newSol := make([]int, len(curSol))
		copy(newSol, curSol)
		*ans = append(*ans, newSol)
		return
	}
	for i, num := range nums {
		// skip used numbers
		if usedNums[i] {
			continue
		}

		// do choice
		curSol = append(curSol, num)
		usedNums[i] = true

		backtracking(curSol, nums, usedNums, ans)

		// undo choices
		curSol = curSol[:len(curSol)-1]
		usedNums[i] = false
	}
}

// @lc code=end

