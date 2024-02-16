/*
 * @lc app=leetcode id=211 lang=golang
 *
 * [211] Design Add and Search Words Data Structure
 */

// @lc code=start

// 208.implement-trie-prefix-tree but with wildcard: `.` match any single character
// Search("ap.le") -> true
// Search("a..") -> true

// `WordDictionary()` Initializes the object.
// `void addWord(word)` Adds word to the data structure, it can be matched later.
// `bool search(word)` Returns true if there is any string in the data structure that matches word or false otherwise. word may contain dots '.' where dots can be matched with any letter.

// TrieNode: use HashMap [char]:[TrieNode] to store childrens of a node. Use a bool variable `isEnd` to indicate if that node is a end of a word.

// AddWord: traverse the string and use current char to traverse the tree. If a child does not exist create it, otherwise go to the the child.

// Search: traverse the string and use current char to traverse the tree. Return true <=> finished the word, and the final node is marked isEnd.

type TrieNode struct {
	Children map[byte]*TrieNode
	isEnd    bool
}

type WordDictionary struct {
	Start *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{
		Start: &TrieNode{
			Children: make(map[byte]*TrieNode),
			isEnd:    false,
		},
	}
}

func (this *WordDictionary) AddWord(word string) {
	curNode := this.Start
	for i, _ := range word {
		char := word[i]

		if child, exists := curNode.Children[char]; !exists {
			newNode := &TrieNode{
				Children: make(map[byte]*TrieNode),
			}
			curNode.Children[char] = newNode
			curNode = newNode
		} else {
			curNode = child
		}
	}
	curNode.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {

}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end

