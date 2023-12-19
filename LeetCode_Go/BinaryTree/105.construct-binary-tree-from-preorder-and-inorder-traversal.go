/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
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

// build a binary tree from preorder and inorder traversal integer arrays

// Given a tree in preorder, the first number is the root. then use the root index to find the left and right subtrees' boundary in inorder. (preorder: 3|9|20,15,7| inorder: 9|3|15,20,7), and recursively build the left and right subtrees. Use a hashmap to store inorder index for faster slicing inorder.

// O(n) time | O(n) space

func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderIdxDict := make(map[int]int)
	for i, v := range inorder {
		inorderIdxDict[v] = i
	}
	return buildSubTree(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1, inorderIdxDict)
}

func buildSubTree(preorder, inorder []int, preLeft, preRight, inLeft, inRight int, inorderIdxDict map[int]int) *TreeNode {
	if preLeft > preRight || inLeft > inRight {
		return nil
	}

	// build root
	root := &TreeNode{Val: preorder[preLeft]}

	// find slicing boundaries
	inorderIdx := inorderIdxDict[preorder[preLeft]]
	leftTreeSize := inorderIdx - inLeft   // (1-0)
	rightTreeSize := inRight - inorderIdx // (4-1)

	// build subtrees
	leftRoot := buildSubTree(preorder, inorder, preLeft+1, preLeft+leftTreeSize, inLeft, inorderIdx-1, inorderIdxDict)
	rightRoot := buildSubTree(preorder, inorder, preRight-rightTreeSize+1, preRight, inorderIdx+1, inRight, inorderIdxDict)

	// connect to root
	root.Left = leftRoot
	root.Right = rightRoot

	return root
}

// @lc code=end

