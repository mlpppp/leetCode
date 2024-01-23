/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// Given a string `s`, find the length of the longest substring without repeating characters. `s` consists of English letters, digits, symbols and `spaces`.

// slide window + hashmap. hashmap record each character's appear number in the window. Expand window and update hashmap. Shrink while there is a character appears more than once. Record maximun length to solution after finsh a shrink.

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	charCount := make(map[byte]int)
	left, right := 0, 0
	for right < len(s) {
		// expand window. add the entry by one (or create the entry with 1 if not exists). Might cause a repeat to char s[right-1]
		charCount[s[right]] += 1
		right++

		// if there is an repeat to char s[right-1], start shrink
		for charCount[s[right-1]] > 1 {
			charCount[s[left]] -= 1
			left++
		}

		// Record maximun length to solution after finsh a shrink.
		maxLength = max(maxLength, right-left)
	}
	return maxLength
}

// @lc code=end

