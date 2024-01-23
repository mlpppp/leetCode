/*
 * @lc app=leetcode id=1888 lang=golang
 *
 * [1888] Minimum Number of Flips to Make the Binary String Alternating
 */

// @lc code=start

// The string is called alternating if no two adjacent characters are equal (eg. "01010101"). You are given a binary string `s`. You are allowed to do two operations: Type-1: pop char from the head and append it to the end. Type-2: pick any digit in `s` and flip its value (0 to 1, 1 to 0). Return the minimum number of Type-2 filp you need to make `s` alternating.

// to make a alternating string, there are only two possibilities: start with 1, 1 in the even position, 0 in the odd position, or start with 0, 0 in the even position, 1 in the odd position. For example, given input s = "111000", we want to compare it with "101010" or "010101" of length 6.

// brute force: O(n^2): given a string of length n, there are `n possible Type-1 processed strings`. The optimal type-2 operations number should be in one of them. Compare each of them to alternating string "101..." and "010..." and find the minimum number of difference. In brute force approach there are overlapping subproblems. Eg. compare "1|11000" with "1|01010" is the same as compare "11000|1" with "01010|1".

// sliding window: concatenate two `s` (s = "111000" => s = "111000111000"). Use two alternating string "101010" and "010101" of length 6 as two windows. Each window maintain the number of difference to current string. When a window move, its new state is: altWindowState - diff(altWindow[0], s[left]) + diff(thisWindow[-1], s[right])

func minFlips(s string) int {
	minT2 := math.MaxInt
	// initialize window templates
	var builder0, builder1 strings.Builder
	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			builder0.WriteByte(byte('0'))
			builder1.WriteByte(byte('1'))
		} else {
			builder0.WriteByte(byte('1'))
			builder1.WriteByte(byte('0'))
		}
	}
	alt0, alt1 := builder0.String(), builder1.String()
	// init window state
	diffAlt0, diffAlt1 := 0, 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] != alt0[i] {
			diffAlt0++
		}
		if s[i] != alt1[i] {
			diffAlt1++
		}
	}
	// start sliding window
	sCat := s + s
	left, right := 0, len(s)-1
	for right < len(sCat) {
		// expand right, update window state
		if alt0[len(s)-1] != sCat[right] {
			diffAlt0 += 1
		}
		if alt1[len(s)-1] != sCat[right] {
			diffAlt1 += 1
		}
		right++

		// if the window size equals len(s), record the result, and shrink left
		if right-left == len(s) {
			minT2 = min(minT2, diffAlt0, diffAlt1)
			// shrink left
			diffAlt0, diffAlt1 = diffAlt1, diffAlt0
			if alt1[0] != sCat[left] {
				diffAlt0 -= 1
			}
			if alt0[0] != sCat[left] {
				diffAlt1 -= 1
			}
			left++
		}

	}

	return minT2
}

// @lc code=end

