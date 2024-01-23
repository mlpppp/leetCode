/*
 * @lc app=leetcode id=567 lang=golang
 *
 * [567] Permutation in String
 */

// @lc code=start

// Given two strings `s1` and `s2`, return true/false if s2 contains a substring which is a permutation of s1. (aka, a substring in `s2` such that contains all characters same number of time as it is in `s1`, and contains no other characters)

// note: similar to 76.minimum-window-substring

// Slidewindow + hashmap. Use hashmap to record remain times required for each character in `s1`. `Expand` the window to the right until all characters in s1 is satisfied. Then while the windows is still satisfied, `shrink` the window(substring). Repeat the process and return true if at any time the substring is satisfied and the len(substring) == len(s1).

// s = "ab", s2 = "eidbaooo"
// eidba => ba => found

func checkInclusion(s1 string, s2 string) bool {
	// init a counter of remain characters required in the window
	satCount := make(map[byte]int)
	for i, _ := range s1 {
		satCount[s1[i]] += 1
	}
	// number of remain characters to be satisfied
	remainUnsat := len(satCount)

	// slidewindow start
	left, right := 0, 0
	for right < len(s2) {
		// add to window, check if the character is of interest
		if _, exists := satCount[s2[right]]; exists {
			satCount[s2[right]]--
			if satCount[s2[right]] == 0 {
				remainUnsat--
			}
		}
		right++
		// while keep the window satisfy, shirking it and check the length equality
		for remainUnsat == 0 {
			if right-left == len(s1) {
				return true
			}
			// shrink the window
			if _, exists := satCount[s2[left]]; exists {
				satCount[s2[left]]++
				if satCount[s2[left]] > 0 {
					remainUnsat++
				}
			}
			left++
		}

	}

	// fmt.Printf("%v, %v", satCount, remainUnsat)
	return false
}

// @lc code=end

