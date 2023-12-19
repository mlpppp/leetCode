/*
 * @lc app=leetcode id=652 lang=golang
 *
 * [652] Find Duplicate Subtrees
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Given a binary tree, return all duplicate subtrees in an array. For each kind of duplicate subtrees, return the root node of any one of the them. Two trees are duplicate if they have the same structure with the same node values.

// Serialize tree with postorder dfs + null string, determine identicality by comparing the serialized strings. Use a hashmap to store serialized strings and determine duplication.

// O(n) time | O(n) space

var SEP string = ","
var NULL string = "#"

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var sequenceFreqMap map[string]int = make(map[string]int)
	var duplicateSubtrees []*TreeNode = make([]*TreeNode, 0)

	_ = treeSerializerHelper(root, sequenceFreqMap, &duplicateSubtrees)

	// for key, value := range sequenceFreqMap {
	// 	fmt.Printf("key: %s, value: %d\n", key, value)
	// }

	return duplicateSubtrees
}

func treeSerializerHelper(root *TreeNode, sequenceFreqMap map[string]int, duplicateSubtrees *[]*TreeNode) string {
	if root == nil {
		return NULL + SEP
	}

	leftSeq := treeSerializerHelper(root.Left, sequenceFreqMap, duplicateSubtrees)
	rightSeq := treeSerializerHelper(root.Right, sequenceFreqMap, duplicateSubtrees)

	rootSeq := leftSeq + rightSeq + strconv.Itoa(root.Val) + SEP

	if _, exists := sequenceFreqMap[rootSeq]; exists {
		sequenceFreqMap[rootSeq] += 1
		if sequenceFreqMap[rootSeq] == 2 { // first time a duplication found
			*duplicateSubtrees = append(*duplicateSubtrees, root)
		}
	} else {
		sequenceFreqMap[rootSeq] = 1
	}

	return rootSeq
}

// @lc code=end

