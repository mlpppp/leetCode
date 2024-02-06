/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start

// Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order. The same number may be chosen from candidates an unlimited number of times.

// backtracking(startIdx:), after a choice is made, we can still make the same choice unlimit times, so we call backtracking(startIdx:) instead of backtracking(startIdx+1:). Exit condition: sum >= target.

func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	backtracking(0, 0, target, []int{}, candidates, &ans)
	return ans
}

func backtracking(start, sum, target int, choices, candidates []int, ans *[][]int) {
	if sum >= target {
		if sum == target {
			// save to ans
			temp := make([]int, len(choices))
			copy(temp, choices)
			*ans = append(*ans, temp)
		}
		return
	}

	for i := start; i < len(candidates); i++ {
		// do choice
		choices = append(choices, candidates[i])
		sum += candidates[i]
		backtracking(i, sum, target, choices, candidates, ans)
		// undo choice
		choices = choices[:len(choices)-1]
		sum -= candidates[i]
	}
}

// @lc code=end

