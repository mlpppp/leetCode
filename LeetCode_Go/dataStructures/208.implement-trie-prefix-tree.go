/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 */

// @lc code=start

// A trie (pronounced as "try") or prefix tree is a tree data structure. It can efficiently check if a string is a prefix of another string. There are various applications such as autocomplete and spellchecker. Implement the Trie class:

// Trie(): Initializes the trie object.
// void insert(String word): Inserts the string word into the trie.
// boolean search(String word): Returns true if the string word is in the trie (i.e., was inserted before), and false otherwise.
// boolean startsWith(String prefix): Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.

// TrieNode: use HashMap [char]:[TrieNode] to store childrens of a node. Use a bool variable `isEnd` to indicate if that node is a end of a word.

// Insert: traverse the string and use current char to traverse the tree. If a child does not exist create it, otherwise to the the child.

// Search: traverse the string and use current char to traverse the tree. Return true <=> finished the word, and the final node is marked isEnd.

// StartsWith: almost identical to Search. but the final Node does not need to be isEnd.

type Trie struct {
	Start *TrieNode
}

type TrieNode struct {
	Children map[byte]*TrieNode
	isEnd    bool
}

func Constructor() Trie {
	return Trie{
		Start: &TrieNode{
			Children: make(map[byte]*TrieNode),
			isEnd:    false,
		},
	}
}

func (this *Trie) Insert(word string) {
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

func (this *Trie) Search(word string) bool {
	curNode := this.Start

	for i, _ := range word {
		char := word[i]
		if child, exists := curNode.Children[char]; exists {
			curNode = child
		} else {
			return false
		}
	}
	return curNode.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	curNode := this.Start

	for i, _ := range prefix {
		char := prefix[i]
		if child, exists := curNode.Children[char]; exists {
			curNode = child
		} else {
			return false
		}
	}
	return true
}

// @lc code=end

