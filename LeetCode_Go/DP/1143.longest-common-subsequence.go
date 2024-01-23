/*
 * @lc app=leetcode id=1143 lang=golang
 *
 * [1143] Longest Common Subsequence
 */

// @lc code=start
// https://www.youtube.com/watch?v=Ua0GhsJSlWM

// Given two strings text1 and text2, return the length of their longest common subsequence(LCS). If there is no common subsequence, return 0.

// breaking problem: notice that two common subsequence must start (or end) with the same character. If adding "a" to both string "ce" and "bcde", it will make the LCS one letter longer from the original LCS; When two string dont start with the same character, eg ["bce", "abcde"] or ["kabce", "abcde"], in the first case, the LCS is the same as LCS(["bce", "bcde"], remove "a" from string2). In the Second case, the LCS is the same as LCS(["abce", "abcde"], remove "k" from string1). We dont know which case, but since the goal is `Longest`, we can return the max of two solutions.

// define Subproblem: LCS(text1[i:], text2[j:]) int: . Decrease i from the len(text1)-1 to 0, and j from len(text2)-1 to 0 to find the result.

// transition: LCS(text1[i:], text2[j:]) = {
// 	1 + LCS(text1[i+1:], text2[j+1:]), if text1[i] == text2[j];
// 	max(LCS(text1[i+1:], text2[j:]), LCS(text1[i:], text2[j+1:])), if text1[i] != text2[j]
// }

// eg.
// LCS("cecde", "ccde") => 1 + LCS("ecde", "cde")
// LCS("becde", "cde") => max(LCS("ecde", "cde"), LCS("becde", "de"))

// "abcdecfgka", "cedfgbac" => "cdfga"
// "cecde", "cde" => "cde"
// "abeeeeeeed", "abdeeeeeee" => "abeeeeeee"

// pseudocode: start
// Input: text1 = "abcdecfgka", text2 = "cedfgbac"
// table := [len(text1)][len(text1)] of int, init with -1
// LCSHelper(text1, text2 string, i, j int) int {
// 	if i or j are out of bounds (effectly empty string){
// 		return 0
// 	}
// 	// use DP table
// 	if table[i][j] != -1 {
// 		return table[i][j]
// 	}
// 	// recursion
// 	if text1[i] == text2[j] {
// 		table[i][j] = 1 + LCSHelper(text1, text2, i+1, j+1)
// 	} else {
// 		table[i][j] = max(LCSHelper(text1, text2, i+1, j), LCSHelper(text1, text2, i, j+1))
// 	}
// 	return table[i][j]
// }
// longestCommonSubsequenceHelper(text1, text2, 0, 0);
// recursive solution

// func longestCommonSubsequence(text1 string, text2 string) int {
// 	// init DP table
// 	table := make([][]int, len(text1))
// 	for i, _ := range table {
// 		table[i] = make([]int, len(text2))
// 		for j, _ := range table[i] {
// 			table[i][j] = -1
// 		}
// 	}
// 	return LCSHelper(text1, text2, 0, 0, table)
// }
// func LCSHelper(text1, text2 string, i, j int, table [][]int) int {
// 	if i >= len(text1) || j >= len(text2) {
// 		return 0
// 	}

// 	// use DP table
// 	if table[i][j] != -1 {
// 		return table[i][j]
// 	}
// 	// transition/recursion
// 	if text1[i] == text2[j] {
// 		table[i][j] = 1 + LCSHelper(text1, text2, i+1, j+1, table)
// 	} else {
// 		table[i][j] = max(LCSHelper(text1, text2, i+1, j, table), LCSHelper(text1, text2, i, j+1, table))
// 	}
// 	for i := range table {
// 		fmt.Printf("%v\n", table[i])
// 	}
// 	fmt.Println()
// 	return table[i][j]
// }

// func max(a, b int) int {
// 	if a < b {
// 		return b
// 	} else {
// 		return a
// 	}
// }

// iterative solution
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

