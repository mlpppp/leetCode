/*
 * @lc app=leetcode id=698 lang=golang
 *
 * [698] Partition to K Equal Sum Subsets
 */

// @lc code=start

// Given an integer array nums and an integer k, return true if it is possible to divide this array into k non-empty subsets whose sums are all equal.

// https://www.youtube.com/watch?v=mBk4I0X46oI

// O(k*2^n), for each subset, we iterate through the nums array, and for each position we decide whether or not putting it into the current subset. If a subset is satisfied(sum(subset) == sums(nums)/k), we go to the next subset. Return true if all subsets are satisfied. As a subset problem, do backtracking with 'start', and use a 'used' array to mark numbers that has been used by other subset, backtracking(nums[start:], k, sum(subset), used), goal: backtracking(k=0)

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// some basic cases to speed up the search
	if sum%k != 0 {
		return false
	}
	slices.Sort(nums)
	slices.Reverse(nums)

	// start backtracking
	targetSum := sum / k
	used := make([]bool, len(nums))
	return backtracking(0, k, 0, targetSum, nums, used)
}
func backtracking(start, k, sum, targetSum int, nums []int, used []bool) bool {
	if k == 0 {
		return true
	}
	if sum == targetSum {
		return backtracking(0, k-1, 0, targetSum, nums, used)
	}

	for i := start; i < len(nums); i++ {
		if nums[i] > targetSum {
			return false
		}
		// prune: skip used and impossible numbers
		if used[i] || sum+nums[i] > targetSum {
			continue
		}
		// choose current number
		used[i] = true
		if backtracking(i+1, k, sum+nums[i], targetSum, nums, used) {
			return true
		}
		// undo choice
		used[i] = false
	}
	return false
}

// @lc code=end

