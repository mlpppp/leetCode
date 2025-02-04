/*
 * @lc app=leetcode id=389 lang=golang
 *
 * [389] Find the Difference
 */

// @lc code=start

// You are given two strings s and t. String t is generated by random shuffling string s and then add one more letter at a random position. Return the letter that was added to t. (eg. s = "abcd", t = "cebad", return 'e')

// concatenate the two string and use XOR to eliminate duplicates characters. (ref: 136, 268). Notes that byte can XOR with byte just like int.

func findTheDifference(s string, t string) byte {
	s = s + t
	var ans byte
	for i, _ := range s {
		ans ^= s[i]
	}
	return ans
}

// @lc code=end

