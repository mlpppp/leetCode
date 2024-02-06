/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 */

// @lc code=start

// Given two integers n and k, return combinations of k numbers chosen from the range [1, n].

// all possible combinations of k numbers chosen from the range [1, n], essentially all subsets of [1, n] of size k (ref. 78.subsets). backtracking(start:) where we only branch from number whose is greater than start. choiceSet: all number greater than the current number

func combine(n int, k int) [][]int {
	ans := [][]int{}
	backtracking(1, n, k, []int{}, &ans)
	return ans
}

func backtracking(start, n, k int, comb []int, ans *[][]int) {
	// exit condition: when size is k, save to ans
	if len(comb) == k {
		temp := make([]int, k)
		copy(temp, comb)
		*ans = append(*ans, temp)
		return
	}

	for i := start; i <= n; i++ {
		// do choice
		comb = append(comb, i)
		backtracking(i+1, n, k, comb, ans)

		// undo choice
		comb = comb[:len(comb)-1]
	}
}

// @lc code=end

