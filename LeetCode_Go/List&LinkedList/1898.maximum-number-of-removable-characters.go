/*
 * @lc app=leetcode id=1898 lang=golang
 *
 * [1898] Maximum Number of Removable Characters
 */

// @lc code=start

// Given two strings `s` and `p` where p is a subsequence of s, and an integer array `removable` containing a subset of indices of s. You want to find a maximum number `k` such that after removing first k indices in removable from s, p is still a subsequence of s.

// brute force: O(nm) + O(m). for each number in removable, scan s once to check if q is still a subsequence of s (skip the number that has been removed, use hashmap)

// binary search O(nlogm)+ O(m): binary search in the removable list. Because we know that the larger the index in removable the less likely p being a subsequence of s, the relation is strictly monotomic. If p fails to be a subsequence of s, then we move left, otherwise we move right. When the serach terminate we expect it to be subsequence (right bound)

func maximumRemovals(s string, p string, removable []int) int {
	left, right := 0, len(removable)-1
	result := -1
	for left <= right {
		mid := (left + right) / 2
		removed := make(map[int]int)
		for i := 0; i <= mid; i++ {
			removed[removable[i]] = 1
		}
		if isSubsequence(s, p, removed) { // expand removed
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if result == -1 { // result has not changed, mid being push to the left bound
		return 0
	}
	return result + 1
}

func isSubsequence(s string, p string, removed map[int]int) bool {
	sIdx, pIdx := 0, 0
	matched := 0
	for sIdx < len(s) && pIdx < len(p) {
		// index in removed, skip
		if _, exists := removed[sIdx]; exists {
			sIdx++
			continue
		}
		// same char, matched++, pIdx++
		if s[sIdx] == p[pIdx] {
			matched++
			pIdx++
		}
		sIdx++
	}
	return matched == len(p)
}

// @lc code=end

