/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */

// @lc code=start

// combination is subset in this problem. backtracking(start, sum) to iterate all subsets of candidates and record answers when sum reaches the target. Since there are duplicated numbers, we need to prune like (90.subsets-ii).

func combinationSum2(candidates []int, target int) [][]int {
	slices.Sort(candidates)
	ans := [][]int{}
	backtracking(0, 0, target, []int{}, candidates, &ans)
	return ans
}

func backtracking(start, sum, target int, subset, candidates []int, ans *[][]int) {
	// exit condition; when the sum is greater or equal to the target
	if sum >= target {
		if sum == target { // save to ans
			temp := make([]int, len(subset))
			copy(temp, subset)
			*ans = append(*ans, temp)
		}
		return
	}

	for i := start; i < len(candidates); i++ {
		// prune, skip duplicated numbers
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		// make choice
		subset = append(subset, candidates[i])
		sum += candidates[i]
		backtracking(i+1, sum, target, subset, candidates, ans)
		// undo choice
		subset = subset[:len(subset)-1]
		sum -= candidates[i]
	}
}

// @lc code=end

