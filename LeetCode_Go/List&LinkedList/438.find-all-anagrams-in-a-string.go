/*
 * @lc app=leetcode id=438 lang=golang
 *
 * [438] Find All Anagrams in a String
 */

// Given two strings `s` and `p`, return an array of all the start indices of substrings which is a permutation of `p`. (aka, a substring in `s` such that contains all characters same number of time as it is in `p`, and contains no other characters)

// same problem as 567.permutation-in-string

// Slidewindow + hashmap. Use hashmap to record remain times required for each character in p. In s, Expand the window to the right until all characters in p is satisfied. Then while the windows is still satisfied, shrink the window(substring). Repeat the process and save window's left indices to solution when the substring is satisfied and the len(substring) == len(p).

// @lc code=start
func findAnagrams(s string, p string) []int {
	var solutions []int
	// init a counter of remain characters required in the window
	satCount := make(map[byte]int)
	for i, _ := range p {
		satCount[p[i]] += 1
	}
	// number of remain characters to be satisfied
	remainUnsat := len(satCount)

	// slidewindow start
	left, right := 0, 0
	for right < len(s) {
		// add to window, check if the character is of interest
		if _, exists := satCount[s[right]]; exists {
			satCount[s[right]]--
			if satCount[s[right]] == 0 {
				remainUnsat--
			}
		}
		right++
		// while keep the window satisfy, shirking it and check the length equality
		for remainUnsat == 0 {
			if right-left == len(p) {
				solutions = append(solutions, left)
			}
			// shrink the window
			if _, exists := satCount[s[left]]; exists {
				satCount[s[left]]++
				if satCount[s[left]] > 0 {
					remainUnsat++
				}
			}
			left++
		}

	}

	// fmt.Printf("%v, %v", satCount, remainUnsat)
	return solutions
}

// @lc code=end

