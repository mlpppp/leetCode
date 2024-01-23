/*
 * @lc app=leetcode id=424 lang=golang
 *
 * [424] Longest Repeating Character Replacement
 */

// @lc code=start

// You are given a string s and an integer k. you can replace `k` characters in s with any other character. Return the longest substring containing the the same letter after performing the replacements.

// sliding window + hashmap. Given any window(substring), it is valid(can become homogeneous after k replacement) when: the len(window) -  number-of-the-prevalent-character <= k, ie we can replace any other characters to the dominant character in less than k replacement, to make the window homogeneous.
// Expand the window as long as it is vaild. Shrink the window to make it vaild again. Log to result when window is vaild.
// use hashmap to store the frequency of characters in the window to find the prevalent-character

func characterReplacement(s string, k int) int {
	charCount := make(map[byte]int)
	maxLen := 0
	left, right := 0, 0
	for right < len(s) {
		// add char to the window, update window state
		charCount[s[right]] += 1
		right++
		// evaluate vaildity
		maxFreq := 0
		for char := range charCount {
			if charCount[char] > maxFreq {
				maxFreq = charCount[char]
			}
		}
		isVaild := (right-left)-maxFreq <= k
		// if window not vaild, start shrinking
		for !isVaild {
			// remove char from the window, update window state
			charCount[s[left]] -= 1
			left++
			// reevaluate vaildity
			maxFreq := 0
			for char := range charCount {
				if charCount[char] > maxFreq {
					maxFreq = charCount[char]
				}
			}
			isVaild = (right-left)-maxFreq <= k
		}
		// window must be vaild when record the result
		maxLen = max(maxLen, right-left)
	}
	return maxLen
}

// @lc code=end

