/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 */

// @lc code=start

// Given two strings `s` and `t`, return a substring in `s` such that contains all characters in `t`, and the length of the substring is minimized. If there is no such substring, return the empty string "". The testcases will be generated such that the answer is unique.

// Use hashmap to record remain times required for each character in `t`. Use a int to log number of characters currently satisfing. Run SlideWindow in `s`. 1.Expand right until all remaining time requirement is satisfied, then 2.shrink left until causing unsatisfications. Repeat the expansion and shrinking. Record the min satisfy length when shrinking.

// s = "ADOBECODEBANC", t = "ABC"
// ADOBEC => DOBEC(unsat) => DOBECODEBA => ODEBA(unsat) => ODEBANC => ANC(unsat) => <end>

func minWindow(s string, t string) string {
	// global answer
	var substring [2]int
	var minLength int = math.MaxInt

	// init hash map with number of appearance in t. goal: make all of them 0 or negative
	satisfyCount := make(map[byte]int)
	for i, _ := range t {
		satisfyCount[t[i]] += 1
	}

	// number of characters to be satisfied. goal: 0
	remainUnsat := len(satisfyCount)

	// slidewindow start
	left, right := 0, 0
	for right < len(s) {
		// expand right, if current char is one of t, update satisfyCount and remainUnsat
		if _, exists := satisfyCount[s[right]]; exists {
			satisfyCount[s[right]]--
			// this character just got satisfied
			if satisfyCount[s[right]] == 0 {
				remainUnsat--
			}
		}
		right++

		// when all satisfied, start log answer and start shrinking
		for remainUnsat == 0 {
			// log answer
			curLength := right - left
			if curLength < minLength {
				minLength = curLength
				substring[0], substring[1] = left, right-1
			}
			// remove a char, if the char is one of t, update satisfyCount and remainUnsat
			if _, exists := satisfyCount[s[left]]; exists {
				satisfyCount[s[left]]++
				if satisfyCount[s[left]] > 0 {
					remainUnsat++ // cause unsatisfiction
				}
			}
			left++
		}
	}

	if minLength == math.MaxInt {
		return ""
	}
	return s[substring[0] : substring[1]+1]
}

// @lc code=end

