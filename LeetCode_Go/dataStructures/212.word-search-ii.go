/*
 * @lc app=leetcode id=212 lang=golang
 *
 * [212] Word Search II
 */

// @lc code=start

// Given an `m x n board` of characters and a list of strings `words` (for reference), return all words in `words` that exist on the board. A word exist on the board if the word can be build from concatenating adjacent cells horizontally or vertically, at least once at any position.

// prefix tree + backtracking: use prefix tree to store the list of words. Reuse part of DFS solution from (79.word-search): DFS the board with the reference of the prefix tree.

// DFS(board[i][j], visited, curNode), at `board[i][j]` and a trace `visited`, we are also at a trie node curNode. check 4 directions(except for cells in the visited) and DFS the direction `only if the next char at the direction is a child of curNode`. If trieNode is an end-of-word, add the word to the result.

// optimize: remove a word from prefix tree whenever it is found in the board (dunno how to implement, might be circular dependencies => can be done outside the recursion)

func findWords(board [][]byte, words []string) []string {
	// 2D visited matrix
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}

	// init prefix tree with words
	trie := trieConstructor()
	for _, word := range words {
		trie.addNode(word)
	}

	ansMap := make(map[string]bool)
	// perform dfs in every cell
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			char := board[i][j]
			if _, exist := trie.Start.Children[char]; exist {
				DFS(i, j, board, visited, trie.Start.Children[char], []byte{char}, ansMap)
			}
		}
	}

	// transform hashset to array
	ans := []string{}
	for key, _ := range ansMap {
		ans = append(ans, key)
	}

	return ans
}

func DFS(i, j int, board [][]byte, visited [][]bool, curNode *trieNode, word []byte, ansMap map[string]bool) {
	// add to result
	if curNode.IsEndOfWord {
		ansMap[string(word)] = true
	}

	visited[i][j] = true

	directions := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, dir := range directions {
		nextI, nextJ := i+dir[0], j+dir[1]
		// choose set: 1. not out of bounds  2. not in visited 3. is a child of curNode.
		if !(nextI >= 0 && nextI < len(board)) ||
			!(nextJ >= 0 && nextJ < len(board[0])) || visited[nextI][nextJ] {
			continue
		}
		char := board[nextI][nextJ]
		if _, exist := curNode.Children[char]; exist { // 3. is a child of curNode.
			// choose the char
			word = append(word, char)
			DFS(nextI, nextJ, board, visited, curNode.Children[char], word, ansMap)
			// undo choose
			word = word[:len(word)-1]
		}
	}

	visited[i][j] = false
}

// Trie Data structure
type trie struct {
	Start *trieNode
}

type trieNode struct {
	Children    map[byte]*trieNode
	IsEndOfWord bool
}

func trieConstructor() *trie {
	return &trie{
		Start: &trieNode{
			Children:    make(map[byte]*trieNode),
			IsEndOfWord: false,
		},
	}
}

func (t *trie) addNode(word string) {
	curNode := t.Start
	for i, _ := range word {
		char := word[i]
		if nextNode, exists := curNode.Children[char]; exists {
			curNode = nextNode
		} else {
			nextNode := &trieNode{
				Children:    make(map[byte]*trieNode),
				IsEndOfWord: false,
			}
			curNode.Children[char] = nextNode
			curNode = nextNode
		}
	}
	curNode.IsEndOfWord = true
}

// ************************************************************************************************

// (brute force: backtracking: time limit exceed) reuse backtracking solution from (79.word-search), with randomised traverse
// O(wmn*4^mn): where w is the length of the words list.

// DP: exist overlapping subproblem, but impossible to use memo (in worst case we have to create a 2d memo for every possible postfix of the input, space complexity too high)
// DFS(board[i][j], visited, word[i:]), at board[i][j], go 4 directions(except for cells in the visited), to satisfy the word[i:].
// Start with DFS(board[i][j] == word[0], visited, word[1:]), goal: DFS(board[i][j], visited, "")

// At any location we check all word in words. Since we only find a word once, if a word is found, remove it from the words array (O(w)).

// func findWords(board [][]byte, words []string) []string {
// 	// 2D visited matrix
// 	visited := make([][]bool, len(board))
// 	for i, _ := range visited {
// 		visited[i] = make([]bool, len(board[0]))
// 	}

// 	// randomize the traversal of board
// 	coors := make([][2]int, len(board)*len(board[0]))
// 	idx := 0
// 	for i := 0; i < len(board); i++ {
// 		for j := 0; j < len(board[0]); j++ {
// 			coors[idx] = [2]int{i, j}
// 			idx++
// 		}
// 	}
// 	for i := len(coors) - 1; i > 0; i-- {
// 		j := rand.Intn(i + 1)
// 		coors[i], coors[j] = coors[j], coors[i]
// 	}

// 	// use hashset to store words for randomly traverse, and easier removal of found words
// 	remainWords := make(map[string]bool)
// 	for _, word := range words {
// 		remainWords[word] = true
// 	}

// 	ans := []string{}

// 	for _, coor := range coors {
// 		i, j := coor[0], coor[1]
// 		// DFS for every possible word
// 		for word, _ := range remainWords {
// 			if board[i][j] == word[0] {
// 				if DFS(board, visited, i, j, []byte(word)[1:]) {
// 					ans = append(ans, word)   // add word to ans
// 					delete(remainWords, word) // remove word from words
// 				}
// 			}
// 		}

// 	}
// 	return ans
// }

// func DFS(board [][]byte, visited [][]bool, i, j int, postfix []byte) bool {
// 	// when postfix is empty, the `word` string is found
// 	if len(postfix) == 0 {
// 		return true
// 	}
// 	visited[i][j] = true
// 	directions := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
// 	for _, dir := range directions {
// 		nextI, nextJ := i+dir[0], j+dir[1]
// 		// 1. not out of range
// 		// 2. satisfy the next char in the word
// 		// 3. not visited in the current search path
// 		if nextI >= 0 && nextI < len(board) &&
// 			nextJ >= 0 && nextJ < len(board[0]) &&
// 			board[nextI][nextJ] == postfix[0] &&
// 			!visited[nextI][nextJ] {
// 			// make choice

// 			if DFS(board, visited, nextI, nextJ, postfix[1:]) {
// 				visited[i][j] = false // always cleanup before return
// 				return true
// 			}
// 		}
// 	}
// 	visited[i][j] = false
// 	return false
// }
// @lc code=end

