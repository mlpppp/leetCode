/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start

// Given an integer array nums of unique elements, return all possible subsets (the power set). Return the solution in any order.

// draw the DFS tree, prune branches by only add branch whose index is greater than current number. backtracking(i:) where i is the first index that will be add to the DFS tree.

func subsets(nums []int) [][]int {
	ans := [][]int{}
	backtracking(0, nums, []int{}, &ans)
	return ans
}

func backtracking(start int, nums, subset []int, ans *[][]int) {

	// add subset to ans
	temp := make([]int, len(subset))
	copy(temp, subset)
	*ans = append(*ans, temp)

	// choice set controled by state: start
	for i := start; i < len(nums); i++ {
		// make choice
		subset = append(subset, nums[i])
		fmt.Printf("%v is calling backtracking(%v)\n", subset, start+1)
		backtracking(i+1, nums, subset, ans)
		// undo choice
		subset = subset[:len(subset)-1]
	}
}

// @lc code=end

