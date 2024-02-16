/*
 * @lc app=leetcode id=647 lang=golang
 *
 * [647] Palindromic Substrings
 */

// @lc code=start

// Given a string s, return the number of palindromic substrings in it.

// A string is a palindrome when it reads the same backward as forward. A substring is a contiguous sequence of characters within the string.

// brute force O(n^3): for n(n+1)/2 substrings check if it is palindromic (O(n))

// O(n^2): (There are two kinds of palindrome: odd length with single center, and even length with two center.) Iterate through the string, and for each character in middle(either as a single center or the first character of a dual center), expand it until it is no more a palindrome, while add to the palindrome counter.

func countSubstrings(s string) int {
	count := 0
	bytes := []byte(s)
	for i := 0; i < len(bytes); i++ {
		findPalindrome(bytes, i, i, &count)
		if i+1 < len(bytes) {
			findPalindrome(bytes, i, i+1, &count)
		}
	}
	return count
}

func findPalindrome(s []byte, i, j int, count *int) {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		*count++
		i--
		j++
	}
}

// @lc code=end

