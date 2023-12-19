/*
 * @lc app=leetcode id=925 lang=golang
 *
 * [889] Construct Binary Tree from Preorder and Postorder Traversal
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

// build a binary tree from postorder and inorder traversal integer arrays. There are multiple solutions, return any of them.

// Given a tree in preorder, the first number is the root. The second number is the left subtree's root, and also the last number in left subtree's slice in the postorder. (1|245|36, 452|63|1, look for `2`). use the left subtree root to find the boundary of left subtree and right subtree in the two arrays. and recursively build the left and right subtrees. Use a hashmap to store inorder index for faster slicing inorder.

// O(n) time | O(n) space

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	postorderIdxMap := make(map[int]int)
	for idx, val := range postorder {
		postorderIdxMap[val] = idx
	}
	return buildSubTree(preorder, postorder, 0, len(preorder)-1, 0, len(postorder)-1, postorderIdxMap)
}

func buildSubTree(preorder, postorder []int, preLeft, preRight, postLeft, postRight int, postorderIdxMap map[int]int) *TreeNode {
	if postLeft > postRight || preLeft > preRight {
		return nil
	}

	// build root
	root := &TreeNode{Val: preorder[preLeft]}
	if preLeft == preRight { // single node case
		root.Left, root.Right = nil, nil
		return root
	}
	// find slicing boundaries
	leftRootVal := preorder[preLeft+1] // assume the second number is the left subtree's root. But left subtree's root can also be nil, which make the solution not unique.

	// postorder
	leftTreePostLeft := postLeft
	leftTreePostRight := postorderIdxMap[leftRootVal]
	rightTreePostLeft := postorderIdxMap[leftRootVal] + 1
	rightTreePostRight := postRight - 1
	// preorder
	leftTreePreLeft := preLeft + 1
	leftTreePreRight := preLeft + 1 + leftTreePostRight - leftTreePostLeft
	rightTreePreLeft := preRight - (rightTreePostRight - rightTreePostLeft)
	rightTreePreRight := preRight

	// build subtrees
	// connect to root
	root.Left = buildSubTree(preorder, postorder, leftTreePreLeft, leftTreePreRight, leftTreePostLeft, leftTreePostRight, postorderIdxMap)
	root.Right = buildSubTree(preorder, postorder, rightTreePreLeft, rightTreePreRight, rightTreePostLeft, rightTreePostRight, postorderIdxMap)

	return root
}

// @lc code=end

