/*
 * @lc app=leetcode id=583 lang=golang
 *
 * [583] Delete Operation for Two Strings
 */

// @lc code=start

// Given two strings word1 and word2, return the minimum number of deletion operation required to make word1 and word2 the same. You can delete exactly one character in either string in a deletion operation. ("sea" -> "ea"; "eat" -> "ea", output: 2)

// esentially equivalent to 1143.longest-common-subsequence of two string. Number of deletion is equal to len(word1)-LCS(word1, word2)+len(word2)-LCS(word1, word2)

func minDistance(word1 string, word2 string) int {
	LCS := longestCommonSubsequence(word1, word2)
	return len(word1) + len(word2) - 2*LCS
}

func longestCommonSubsequence(text1 string, text2 string) int {
	// init DP table (initialize with 0) (adding one row and one column for padding, essentially define empty string)
	table := make([][]int, len(text1)+1)
	for i, _ := range table {
		table[i] = make([]int, len(text2)+1)
	}
	// iterative from last row to first row (then last column to first column)
	for i := len(text1) - 1; i >= 0; i-- {
		for j := len(text2) - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				table[i][j] = 1 + table[i+1][j+1]
			} else {
				table[i][j] = max(table[i+1][j], table[i][j+1])
			}
		}
	}
	return table[0][0]
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

// @lc code=end

