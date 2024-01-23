/*
 * @lc app=leetcode id=712 lang=golang
 *
 * [712] Minimum ASCII Delete Sum for Two Strings
 */

// @lc code=start

// Given two strings s1 and s2, return the lowest ASCII sum of deleted characters to make two strings equal.

// review 1143.longest-common-subsequence before this problem (https://www.youtube.com/watch?v=Ua0GhsJSlWM)

// firstly, it is a common subsequence problem, so the same framework as 1143.longest-common-subsequence. The target is different. The target common subsequence is not only the longest, if there are multiple LCS, the chosen LCS should `maximize the ASCII sum`. (When remaining string's ASCII sum is maximized, the ASCII sum of deleted characters is minimized). the result is asciiSum(s1) + asciiSum(s2) - 2 * maxASCS(s1, s2)

// subproblem maxASCIISumCommonSubsequence(s1[i:], s2[j:]) int maxASCIISum; aka `maxASCS`.

// transition: maxASCS(s1[i:], s2[j:]) int max: {
// 	maxASCS(s1[i+1:], s2[j+1:]) + ascii(s1[i]), if s1[i] == s2[j];
// 	max(maxASCS(s1[i+1:], s2[j:]), maxASCS(s1[i:], s2[j+1:])), if s1[i] == s2[j]
// }

// // pseudocode(iterative)
// table := [len(s1)+1][len(s2)+1]int initialized with 0
// for i from len(s1)-1 to 0, step -1:
// 	for j from len(s2)-1 to 0, step -1:
// 		if s1[i] == s2[j]:
// 			table[i][j] = table[i+1][j+1] + ascii(s1[i])
// 		else:
// 			table[i][j] = max(table[i+1][j], table[i][j+1])

func minimumDeleteSum(s1 string, s2 string) int {
	// initalize DP table with one extra row/column for padding
	table := make([][]int, len(s1)+1)
	for i, _ := range table {
		table[i] = make([]int, len(s2)+1)
	}
	for i := len(s1) - 1; i >= 0; i-- {
		for j := len(s2) - 1; j >= 0; j-- {
			if s1[i] == s2[j] {
				table[i][j] = table[i+1][j+1] + int(s1[i])
			} else {
				table[i][j] = max(table[i+1][j], table[i][j+1])
			}
		}
		// for k, _ := range table {
		// 	fmt.Printf("%v\n", table[k])
		// }
	}
	// fmt.Printf("s1 AsciiSum: %v\n, s2 AsciiSum: %v\n", asciiSum(s1), asciiSum(s2))
	return asciiSum(s1) + asciiSum(s2) - 2*table[0][0]
}

func asciiSum(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += int(s[i])
	}
	return sum
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

// @lc code=end

