package main

import "fmt"

// Given an `m x n board` of characters and a list of strings `words` (for reference), return all words in `words` that exist on the board. A word exist on the board if the word can be build from concatenating adjacent cells horizontally or vertically, at least once at any position.

// prefix tree + backtracking: use prefix tree to store the list of words. Reuse part of DFS solution from (79.word-search): DFS the board with the reference of the prefix tree.

// DFS(board[i][j], visited, curNode), at `board[i][j]` and a trace `visited`, we are also at a trie node curNode. check 4 directions(except for cells in the visited) and DFS the direction `only if the next char at the direction is a child of curNode`. If trieNode is an end-of-word, add the word to the result.

// optimize: remove a word from prefix tree whenever it is found in the board (dunno how to implement, might be circular dependencies)

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

func main() {

	board := [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}
	words := []string{"oath", "pea", "eat", "rain", "o", "oa"}

	// board := [][]byte{{'o', 'a', 'b', 'n'}, {'o', 't', 'a', 'e'}, {'a', 'h', 'k', 'r'}, {'a', 'f', 'l', 'v'}}
	// words := []string{"oa", "oaa"}
	fmt.Println(findWords(board, words))

	// param_3 := obj.StartsWith(prefix);6}, 1))
	// fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 2))
	// fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 3))
	// fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
	// fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 5))
	// fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 6))
	// fmt.maxPathSum(countSubstrings("abcbcddabbdccd"))

	// fmt.Println(isPalindrome("race a car"))

	// fmt.Println(reverseWords("  hello world  "))
	// fmt.Println(reverseWords("a good   example"))
	// fmt.Println(maxProduct([]int{-2, 0, -1}))
	// fmt.Println(maxProduct([]int{2, -2, 3, 6, -8, 2}))
	// fmt.Println(maxProduct([]int{2, -2, 0, 3, 6, -8, 2}))

	// fmt.Println(rangeBitwiseAnd("11", "1"))
	// fmt.Println(rangeBitwiseAnd("0", "0"))
	// fmt.Println(rangeBitwiseAnd("01", "1"))

}
