/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 */

// @lc code=start

// A message containing letters from A-Z can be encoded into numbers through ('A' -> "1", 'B' -> "2"... 'Z' -> "26"). Given an encoded message, For example, "11106", there are multiple ways to decode it back to letters. Given a string s containing only digits, return the number of ways to decode it. (note: '0' or '06' are not vaild)

// subproblem: numDecodings(s[i:]), at each position we either take one or two digits since all characters are not longer than two digits. Goal: numDecodings(s[:])

// transition: numDecodings(s[i:]) = {
// 	if s[i] == '0', 0 else numDecodings(s[i+1:]) +
// 	if s[i:i+2] is vaild letter, 0 else numDecodings(s[i+2:])
// }

// dpTable: 1D array of len(s)+1, one '1' padding at the end

// iterative solution: right to left
func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[len(s)] = 1
	for i := len(s) - 1; i >= 0; i-- {
		oneDigit := 0
		if s[i] != '0' {
			oneDigit = dp[i+1]
		}
		twoDigit := 0
		if i+1 < len(s) && (s[i] == '1' || (s[i] == '2') && '0' <= s[i+1] && s[i+1] <= '6') {
			twoDigit = dp[i+2]
		}
		dp[i] = oneDigit + twoDigit
	}
	return dp[0]
}

// @lc code=end

