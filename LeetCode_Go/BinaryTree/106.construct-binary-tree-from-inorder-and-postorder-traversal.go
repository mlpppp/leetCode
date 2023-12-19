/*
 * @lc app=leetcode id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
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

// build a binary tree from postorder and inorder traversal integer arrays

// Given a tree in postorder, the last number is the root. then use the root index to find the left and right subtrees' boundary in inorder. (postorder: 9|15,7,20|3 inorder: 9|3|15,20,7), and recursively build the left and right subtrees. Use a hashmap to store inorder index for faster slicing inorder.

// O(n) time | O(n) space
func buildTree(inorder []int, postorder []int) *TreeNode {
	inorderIdxDict := make(map[int]int)
	for i, v := range inorder {
		inorderIdxDict[v] = i
	}
	return buildSubTree(inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1, inorderIdxDict)
}

func buildSubTree(inorder, postorder []int, inLeft, inRight, postLeft, postRight int, inorderIdxDict map[int]int) *TreeNode {
	if postLeft > postRight || inLeft > inRight {
		return nil
	}

	// build root
	root := &TreeNode{Val: postorder[postRight]}

	// find slicing boundaries
	inorderIdx := inorderIdxDict[postorder[postRight]]
	leftTreeSize := inorderIdx - inLeft // (3-0)
	// rightTreeSize := inRight - inorderIdx // (5-3)

	// build subtrees
	leftRoot := buildSubTree(inorder, postorder, inLeft, inorderIdx-1, postLeft, postLeft+leftTreeSize-1, inorderIdxDict)
	rightRoot := buildSubTree(inorder, postorder, inorderIdx+1, inRight, postLeft+leftTreeSize, postRight-1, inorderIdxDict)

	// connect to root
	root.Left = leftRoot
	root.Right = rightRoot

	return root
}

// @lc code=end

