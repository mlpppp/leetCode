/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
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

// Given a binary tree, return the maximum depth of the tree (length of longest path from the root node to deepest leaf node).

// Travserse the Binary Tree. pre-order operation: add 1 to the depth; post-order operation: subtract 1 from the depth; States: `curDepth` dynamically change, `depth` maintains the maximum depth.

// O(n) time | O(n) space

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxDepth, curDepth := 0, 0
	traverse(root, &maxDepth, &curDepth)
	return maxDepth
}

func traverse(root *TreeNode, maxDepth, curDepth *int) {
	if root == nil {
		return
	}
	*curDepth++
	if *curDepth > *maxDepth {
		*maxDepth = *curDepth
	}
	traverse(root.Left, maxDepth, curDepth)
	traverse(root.Right, maxDepth, curDepth)
	*curDepth--
}

// @lc code=end

