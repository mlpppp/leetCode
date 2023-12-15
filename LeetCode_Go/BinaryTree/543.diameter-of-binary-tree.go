/*
 * @lc app=leetcode id=543 lang=golang
 *
 * [543] Diameter of Binary Tree
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

// Given a binary tree, return the length of the diameter of the tree (length of the longest path between any two nodes in a tree).

// For each node, a diameter that pass the node is the left tree's depth plus the right tree's depth. Post-order Traversal, calculate the diameter and depth from depth of the left and right subtrees.

// O(n) time | O(n) space

func diameterOfBinaryTree(root *TreeNode) int {
	var maxDiameter int = 0
	_ = traverse(root, &maxDiameter)
	return maxDiameter
}

func traverse(root *TreeNode, maxDiameter *int) int {
	if root == nil {
		return -1
	}

	leftDepth := traverse(root.Left, maxDiameter)
	rightDepth := traverse(root.Right, maxDiameter)

	diameter := leftDepth + rightDepth + 2
	if diameter > *maxDiameter {
		*maxDiameter = diameter
	}
	depth := max(leftDepth, rightDepth) + 1

	// fmt.Printf("cur node: %d, leftDepth: %d, rightDepth: %d, diameter: %v\n", root.Val, leftDepth, rightDepth, diameter)
	return depth
}

// @lc code=end

