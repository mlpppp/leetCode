/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

// @lc code=start

// Given a string s, return true if it is a palindrome, or false otherwise. s consists only of printable ASCII characters.
// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric (letters and number) characters, it reads the same forward and backward.

// Input: s = "A man, a plan, a canal: Panama"
// Output: true

// 1. convert all uppercase letters to lowercase letters: 'a' - 'A' == 32
// 2. remove all non-alphanumeric characters: 'a' < letter < 'z' || '0' < letter < '9'
// 3. determine palindrome: two pointers, start and end toward middle.

func isPalindrome(s string) bool {
	// 1. convert all uppercase letters to lowercase letters
	// 2. remove all non-alphanumeric characters:
	var newS strings.Builder
	for i, _ := range s {
		char := s[i]
		if char >= 'A' && char <= 'Z' {
			newS.WriteByte(char + 32)
			continue
		}
		if char >= 'a' && char <= 'z' || char >= '0' && char <= '9' {
			newS.WriteByte(char)
		}
	}
	s = newS.String()

	// 3. determine palindrome: two pointers
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

// @lc code=end

