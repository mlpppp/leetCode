/*
 * @lc app=leetcode id=72 lang=golang
 *
 * [72] Edit Distance
 */

// @lc code=start

// Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2. Three operations are permitted on: Insert, Delete, Replace

// Distance from word1 to word2 is the same with distant from word2 to word1 (through reversed operations).

// subproblem: editDistance(word1[i:], word2[j:]): edit distance to change string word1[i:], at a time, we edit word1[i] and word2[j], which make them the same, and the subproblem becomes another editDistance problem. Different opertaion has different implications. We want to calculate the full strings editDistance(word1[0:], word2[0:]) from the right hand side.

// transitions: editDistance(word1[i:], word2[j:]) = {
// 	if word1[i]==word1[j] {
// 		0 + editDistance(word1[i+1:], word2[j+1:]), do nothing, the letter is already the same.
// 	}
// 	if word1[i]!=word1[j] {
// 		min(
// 			1 + editDistance(word1[i+1:], word2[j+1:]), replace operation:  replacing word1[i] with word2[j] move both words letter-of-interest to the right by 1;
// 			1 + editDistance(word1[i+1:], word2[j:]), deletion operation, delete word1[i] move word1 letter-of-interest to the right by 1;
// 			1 + editDistance(word1[i:], word2[j+1:]), insertion operation, inserting word2[j] before word1[i], the word1 letter-of-interest stay, and word2 letter-of-interest can be moved to the right by 1;
// 		)
// 	}
// }

// base case: {
// 	editDistance(emptyString, word2[j:]) = len(word2[j:]);
// 	editDistance(word1[i+1:], emptyString) = word1[i+1:]
// }

// DP table
// a[i,j] -> a[i,j+1]
// |     \
// v        a[i+1,j+1]
// a[i+1,j]

// recursion implements
// func minDistance(word1 string, word2 string) int {
// 	// DP table, filled with -1 (distance cannot be negative)
// 	table := make([][]int, len(word1))
// 	for i, _ := range table {
// 		table[i] = make([]int, len(word2))
// 		for j, _ := range table[i] {
// 			table[i][j] = -1
// 		}
// 	}
// 	return editDistance(word1, word2, 0, 0, table)
// }

// func editDistance(word1, word2 string, i, j int, table [][]int) int {
// 	if i >= len(word1) { // basecase1: convert emptyString to word2[j:]
// 		return len(word2) - j
// 	}
// 	if j >= len(word2) { // basecase2: convert word1[i:] to emptyString
// 		return len(word1) - i
// 	}
// 	// use DP table
// 	if table[i][j] != -1 {
// 		return table[i][j]
// 	}
// 	// recursion call
// 	if word1[i] == word2[j] {
// 		table[i][j] = editDistance(word1, word2, i+1, j+1, table)
// 	} else {
// 		table[i][j] = min(
// 			1+editDistance(word1, word2, i+1, j+1, table),
// 			1+editDistance(word1, word2, i+1, j, table),
// 			1+editDistance(word1, word2, i, j+1, table))
// 	}

// 	return table[i][j]
// }

// func min(a, b, c int) int {
// 	if a <= b && a <= c {
// 		return a
// 	}
// 	if b <= a && b <= c {
// 		return b
// 	}
// 	return c
// }

// Iterative implements
func minDistance(word1 string, word2 string) int {
	// DP table, 1 row/column padding, filled with basecases
	table := make([][]int, len(word1)+1)
	for i, _ := range table {
		table[i] = make([]int, len(word2)+1)
		for j, _ := range table[i] {
			table[i][j] = -1
		}
	}
	// last row basecases
	for i := 0; i <= len(word2); i++ {
		table[len(word1)][i] = len(word2) - i
	}
	// last column basecases
	for i := 0; i <= len(word1); i++ {
		table[i][len(word2)] = len(word1) - i
	}
	// fill table from bottom to top, right to left
	for i := len(word1) - 1; i >= 0; i-- {
		for j := len(word2) - 1; j >= 0; j-- {
			if word1[i] == word2[j] {
				table[i][j] = table[i+1][j+1]
			} else {
				table[i][j] = min(
					1+table[i+1][j],
					1+table[i+1][j+1],
					1+table[i][j+1],
				)
			}
		}

	}
	for row, _ := range table {
		fmt.Printf("%v\n", table[row])
	}
	return table[0][0]
}
func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

// @lc code=end

