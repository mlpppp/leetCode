/*
 * @lc app=leetcode id=187 lang=golang
 *
 * [187] Repeated DNA Sequences
 */

// @lc code=start

// A DNA sequence is composed of a series of nucleotides abbreviated as 'A', 'C', 'G', and 'T'. eg "ACGAATTCCG". Given a DNA string `s`, return all the `10-letter-long substrings` that occur more than once.

// slide window + hashmap. Maintain a hashmap of seen patterns. Use fix size slide window of size 10. For each i, check if substring s[i:i+10] has been seen once. If has been seen once, add to result, if has been seen more than once, do nothing.

// O(n*10)

// func findRepeatedDnaSequences(s string) []string {
// 	seenPatterns := make(map[string]int)
// 	var result []string
// 	for i := 0; i < len(s)-9; i++ {
// 		if val, exists := seenPatterns[s[i:i+10]]; exists {
// 			if val == 1 {
// 				result = append(result, s[i:i+10])
// 			}
// 			seenPatterns[s[i:i+10]] += 1
// 		} else {
// 			seenPatterns[s[i:i+10]] = 1
// 		}
// 	}
// 	return result
// }

// optimization: Use int32 to encode a 10-letter substrings ('A', 'C', 'G', and 'T' in quaternary numbers). Dynamically encode the substrings when moving (ie. remove left and add right) in O(1). reduce the complexity to O(n)

func findRepeatedDnaSequences(s string) []string {
	seenPatterns := make(map[int]int)
	var result []string

	// convert string to quaternary representation
	sQuater := make([]int, len(s))
	for i := 0; i < len(sQuater); i++ {
		switch s[i] {
		case 'A':
			sQuater[i] = 0
		case 'G':
			sQuater[i] = 1
		case 'C':
			sQuater[i] = 2
		case 'T':
			sQuater[i] = 3
		}
	}

	windowHash := 0
	left, right := 0, 0
	for right < len(sQuater) {
		// expand windows right by one and update the hash
		windowHash = windowHash*4 + sQuater[right]
		right++

		if right-left == 10 {
			// check DNA sequence hash
			if val, exists := seenPatterns[windowHash]; exists {
				if val == 1 {
					result = append(result, s[left:right])
				}
				seenPatterns[windowHash] += 1
			} else {
				seenPatterns[windowHash] = 1
			}
			// shrink windows left by one and update the hash
			windowHash -= sQuater[left] * int(math.Pow(4, float64(9)))
			left++
		}
	}
	return result
}

// @lc code=end

