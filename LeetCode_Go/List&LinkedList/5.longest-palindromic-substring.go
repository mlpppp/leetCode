/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start

// return the longest palindromic substring in a string

// (There are two kinds of palindrome: odd length with single center, and even length with two center.) Iterate through the string, and for each character (either as a single center or the first character of a dual center), find the longest palindromic substring. Maintain globally the longest palindromic substring.

// O(n^2) time | O(1) space

func longestPalindrome(s string) string {
	var longestPalindrome string = ""
	for i := 0; i < len(s); i++ {
		// fmt.Printf("i = %v \n", i)
		findLongestPalindrome(s, i, i, &longestPalindrome)
		findLongestPalindrome(s, i, i+1, &longestPalindrome)
		// fmt.Printf("single center: %v \n", longestPalindrome)
		// fmt.Printf("dual center: %v \n", longestPalindrome)
	}
	return longestPalindrome
}
func findLongestPalindrome(s string, centerA, centerB int, longestPalindrome *string) {
	left := centerA
	right := centerB
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	if right-left-1 > len(*longestPalindrome) {
		*longestPalindrome = s[left+1 : right]
	}
}

// @lc code=end

