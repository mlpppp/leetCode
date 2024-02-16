/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */

// @lc code=start

// Given an m x n grid of characters board and a string word, return true if word can be found by concatenating adjacent cells horizontally or vertically

// DP: exist overlapping subproblem, but impossible to use memo (in worst case we have to create a 2d memo for every possible postfix of the input, space complexity too high)

// backtracking: DFS(board[i][j], visited, word[i:]), at board[i][j], go 4 directions(except for cells in the visited), to satisfy the word[i:].
// Start with DFS(board[i][j] == word[0], visited, word[1:]), goal: DFS(board[i][j], visited, "")
func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i, _ := range visited {
		visited[i] = make([]bool, len(board[0]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			// DFS for every start char word[0]
			if board[i][j] == word[0] {
				if DFS(board, visited, i, j, []byte(word)[1:]) {
					return true
				}
			}
		}
	}
	return false
}

func DFS(board [][]byte, visited [][]bool, i, j int, postfix []byte) bool {
	// when postfix is empty, the `word` string is found
	if len(postfix) == 0 {
		return true
	}
	visited[i][j] = true
	directions := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	for _, dir := range directions {
		nextI, nextJ := i+dir[0], j+dir[1]
		// 1. not out of range
		// 2. satisfy the next char in the word
		// 3. not visited in the current search path
		if nextI >= 0 && nextI < len(board) &&
			nextJ >= 0 && nextJ < len(board[0]) &&
			board[nextI][nextJ] == postfix[0] &&
			!visited[nextI][nextJ] {
			// make choice

			if DFS(board, visited, nextI, nextJ, postfix[1:]) {
				visited[i][j] = false // always cleanup before return
				return true
			}
		}
	}
	visited[i][j] = false
	return false
}

// @lc code=end

