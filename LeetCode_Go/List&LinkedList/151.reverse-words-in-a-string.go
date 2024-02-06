/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 */

// @lc code=start

// Given an input string s, reverse the order of the words. Reversed string should remove leading or trailing spaces or rudundant spaces. Solve it in-place with O(1) extra space.

// reverse the whole string, then scen the reversed string and for each reversed word reverse the word.

func reverseWords(s string) string {
	str := []byte(s)
	// reverse the string
	i, j := 0, len(str)-1
	for i < j {
		str[i], str[j] = str[j], str[i]
		i++
		j--
	}
	// reverse each world and save to the answer
	// i find start and end of a word (start: a space followed by a char, end: a char followed by a space/out of bound)
	var ans strings.Builder
	str = append([]byte{' '}, str...) // for edge cases
	for i := 1; i < len(str); i++ {
		// find the start and end of a word
		if str[i-1] == ' ' && str[i] != ' ' {
			start := i - 1
			var end int
			for j := start + 1; j < len(str); j++ {
				if str[j] != ' ' && (j+1 == len(str) || str[j+1] == ' ') {
					end = j
					break
				}
			}
			// copy to the answer string
			for char := end; char >= start; char-- {
				ans.WriteByte(str[char])
			}
			// reset i
			i = end + 1
		}
	}
	ansStr := ans.String()
	if ansStr[len(ansStr)-1] == ' ' {
		ansStr = ansStr[:len(ansStr)-1]
	}
	return ansStr
}

// @lc code=end

