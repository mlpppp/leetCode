/*
 * @lc app=leetcode id=516 lang=golang
 *
 * [516] Longest Palindromic Subsequence
 */

// @lc code=start

// Given a string s, find the longest palindromic subsequence's length in s.

// Subproblems: Two pointer 2D-DP. LPS(s[start:end]): LPS of a substring. Target: LPS(s[0:len(s)-1]). Start and the end will not out of range as they march from the edge to the middle.

// transition: LPS(s[start:end]) = {
// 	if s[start] == s[end], return 2+ LPS(s[start+1, end-1])
// 	if s[start] != s[end], return max(LPS(s[start+1, end]), LPS(s[start, end-1]))
// }

// base case:  start and the end are the same, return palindrome length 1. If start > end (empty string), return palindrome length 0

// recursive
// func longestPalindromeSubseq(s string) int {
// 	//  table [len(s)][len(s)] initalized with 0 (because 1 <= s.length, longest palindrome is at least 1)
// 	table := make([][]int, len(s))
// 	for i, _ := range table {
// 		table[i] = make([]int, len(s))
// 	}
// 	return LPS(s, 0, len(s)-1, table)
// }

// func LPS(s string, start, end int, table [][]int) int {
// 	//  base case
// 	if start > end { // empty string
// 		return 0
// 	}
// 	if start == end { // single character string
// 		table[start][end] = 1
// 		return table[start][end]
// 	}

// 	//  query DP table
// 	if table[start][end] > 0 {
// 		return table[start][end]
// 	}

// 	// transition
// 	if s[start] == s[end] {
// 		table[start][end] = 2 + LPS(s, start+1, end-1, table)
// 	} else {
// 		table[start][end] = max(
// 			LPS(s, start+1, end, table),
// 			LPS(s, start, end-1, table))
// 	}
// 	return table[start][end]

// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	} else {
// 		return b
// 	}
// }

// iteration implement: row: bottom-up, column: left to right
// test case: "bbbab"
// [0 0 0 0 4]
// [0 0 2 2 0]
// [0 0 1 1 0]	 		s[i,j-1] <- s[i,j]
// [0 0 0 1 0]   				  / |
// [0 0 0 0 0]   		s[i+1,j-1]	s[i+1,j]

func longestPalindromeSubseq(s string) int {
	//  table [len(s)][len(s)] initalized with 0 (because 1 <= s.length, longest palindrome is at least 1)
	table := make([][]int, len(s))
	for i, _ := range table {
		table[i] = make([]int, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := 0; j < len(s); j++ {
			// base case
			if i > j {
				continue
			}
			if i == j {
				table[i][j] = 1
				continue
			}
			// transition case
			if s[i] == s[j] {
				table[i][j] = 2 + table[i+1][j-1]
			} else {
				table[i][j] = max(table[i+1][j], table[i][j-1])
			}
		}
	}
	return table[0][len(s)-1]
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end

