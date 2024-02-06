/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */

// @lc code=start

// Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

// use hashset, in one pass O(n) + O(n)
func containsDuplicate(nums []int) bool {
	dup := make(map[int]bool)
	for _, num := range nums {
		if _, exist := dup[num]; exist {
			return true
		} else {
			dup[num] = false
		}
	}
	return false
}

// @lc code=end

