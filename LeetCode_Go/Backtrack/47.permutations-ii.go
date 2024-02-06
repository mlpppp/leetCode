/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */

// @lc code=start

// Given a collection of numbers, nums, that might contain duplicates, return all possible unique permutations in any order.

// two factors prune:  1. (ref.46.permutations) use usedNums, []bool as the alternative to choiceSet, skip numbers that has been used. 2. sort the input and skip current choice num[i] if same number has been visited in current call stack iteration.

func permuteUnique(nums []int) [][]int {
	slices.Sort(nums)
	ans := [][]int{}
	used := make([]bool, len(nums))
	backtracking(nums, []int{}, used, &ans)
	return ans
}

func backtracking(nums, perm []int, used []bool, ans *[][]int) {
	// exit: when the length is correct
	if len(perm) == len(nums) {
		// save to ans
		temp := make([]int, len(perm))
		copy(temp, perm)
		*ans = append(*ans, temp)
		return
	}

	// reinitialize lastAdd in every call stack, to skip branches that have already been visited
	lastVisited := -11 // init value must not appear in nums
	for i, num := range nums {
		// prune 1: used, prune 2: deal with duplicate
		if used[i] || (nums[i] == lastVisited) {
			continue
		}
		// make choice
		used[i] = true
		perm = append(perm, num)
		lastVisited = num
		backtracking(nums, perm, used, ans)
		// undo choice
		used[i] = false
		perm = perm[:len(perm)-1]
	}
}

// @lc code=end

