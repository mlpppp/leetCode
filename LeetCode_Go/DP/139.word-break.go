/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */

// @lc code=start

// Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words. Note that the same word in the dictionary may be reused multiple times in the segmentation.

// subproblem: wordBreak(s[i:], wordDict) bool: whether substring s[i:] can be can be segmented by word in wordDict. Goal: wordBreak(s[:], wordDict)

// transition: wordBreak(s[i:], wordDict) = {
// 	for word in wordDict {
// 		if s[i:i+len(word)] == word && wordBreak(s[i+len(word):], wordDict), return true
// 	}
// 	return false
// }

// overlapping subproblem: there may be multiple way to reach a position i: "applepenapple", wordDict = ["apple","pen", "applepen"]

// basecase: a word reach the end of string

// complexity: O(n*n*m) the recursion fill DPTable of at most n size, in order to solve each of the problem, it had to compare m words to a substring of size at most n.

// DFS + memo solution
// func wordBreak(s string, wordDict []string) bool {
// 	DPTable := make([]int, len(s)) // 0: uninit, 1: true, -1: false
// 	return wordBreakHelper(s, 0, wordDict, DPTable)
// }

// func wordBreakHelper(s string, i int, wordDict []string, DPTable []int) bool {
// 	// base case: the postfix is a word
// 	if i == len(s) {
// 		return true
// 	}
// 	// query DP table
// 	dp := DPTable[i]
// 	if dp != 0 {
// 		if dp == 1 {
// 			return true
// 		}
// 		if dp == -1 {
// 			return false
// 		}
// 	}
// 	// check every word in the dictionary
// 	for _, word := range wordDict {
// 		if i+len(word) <= len(s) && s[i:i+len(word)] == word { // prefix matches a word
// 			if wordBreakHelper(s, i+len(word), wordDict, DPTable) {
// 				DPTable[i] = 1
// 				return true
// 			} else {
// 				DPTable[i+len(word)] = -1
// 			}
// 		}
// 	}
// 	// no luck for all words
// 	DPTable[i] = -1
// 	return false
// }

// iteration solution: dp table initialize with falsem, one extra padding append to dp table, true), solve right to left

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[len(s)] = true // base case
	for i := len(s) - 1; i >= 0; i-- {
		for _, word := range wordDict {
			if i+len(word) <= len(s) &&
				s[i:i+len(word)] == word && // prefix matches a word
				dp[i+len(word)] { // the remain can be segmented

				dp[i] = true
				break

			}
		}
	}
	return dp[0]
}

// @lc code=end

