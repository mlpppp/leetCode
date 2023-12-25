/*
 * @lc app=leetcode id=222 lang=golang
 *
 * [222] Count Complete Tree Nodes
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
// Given the root of a complete binary tree, return the number of the nodes in the tree.

// Notice that no matter how the leaves align in the last level, either The left subtree or the right subtree must be a perfect binary tree, and the other one is a complete/perfect binary tree. For a perfect binary tree we can solve the problem with formula (2^n - 1) in O(1). Then we can recursivly solve the other subtree.

// T(n) = 2log(n) + T(n/2) = ... = log(n)*log(n)

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var leftCount int
	var rightCount int
	// check if the left subtree is a perfect binary tree
	left, right := root.Left, root.Left
	leftHeight, rightHeight := 0, 0
	for left != nil {
		leftHeight++
		left = left.Left
	}
	for right != nil {
		rightHeight++
		right = right.Right
	}
	if leftHeight == rightHeight { // left subtree is a perfect binary tree
		leftCount = 1<<leftHeight - 1
		rightCount = countNodes(root.Right)
	} else { // right subtree is a perfect binary tree of one level less than the left subtree (happens to be rightHeight)
		leftCount = countNodes(root.Left)
		rightCount = 1<<rightHeight - 1
	}
	return leftCount + rightCount + 1
}

// @lc code=end

