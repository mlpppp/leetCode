/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
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

// Given a binary tree, determine if it is height-balanced.

// A `height balanced binary tree` is a binary tree that for any node, the height of the left subtree and right subtree does not differ by more than 1. And the left subtree and right subtree are also height balanced. Following the definition.

// O(n) time | O(h) space

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	_, isBalanced := traverse(root)
	return isBalanced

}

func traverse(root *TreeNode) (depth int, isBalanced bool) {
	if root == nil {
		return 0, true
	}

	lDepth, isBalanced := traverse(root.Left)
	if !isBalanced {
		return -1, false
	}
	rDepth, isBalanced := traverse(root.Right)
	if !isBalanced {
		return -1, false
	}
	// fmt.Printf("root: %d, leftDepth: %d, rightDepth: %d\n", root.Val, lDepth, rDepth)

	if lDepth-rDepth > 1 || rDepth-lDepth > 1 {
		return -1, false
	}

	return max(lDepth, rDepth) + 1, true

}

// @lc code=end

