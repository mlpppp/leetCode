/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start

// Given two strings s and t, return true if t is an anagram of s, and false otherwise. An Anagram is a word formed by rearranging the letters of a different word, using all the original letters exactly once. s and t consist of lowercase English letters.
// Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?

// use two hashmaps, or two array of len 26 to count each char occurance in two string, and then compare the number of occurance. Follow up: use rune type (in golang)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// count occurrences
	sCount, tCount := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(s); i++ {
		sCount[s[i]] += 1
		tCount[t[i]] += 1
	}
	// compare occurances
	for key, _ := range sCount {
		if sCount[key] != tCount[key] {
			return false
		}
	}

	return true
}

// @lc code=end

