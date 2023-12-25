/*
 * @lc app=leetcode id=112 lang=golang
 *
 * [112] Path Sum
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
// https://www.youtube.com/watch?v=LSKQyOz_P8I

// Given a binary tree and an integer targetSum, return true if the tree has a root-to-leaf path such that adding up to targetSum.

// In a node, there is only one path from the root to it (aka, only one sum). Use the path sum as the recursive input. If the node is a leaf, return true/false, otherwise accumulates the path sum in the two subtrees. Base case: if the node is a leaf.

// O(n) time worst case
func hasPathSum(root *TreeNode, targetSum int) bool {
	// edge case: empty tree
	if root == nil {
		return false
	}
	return pathSumHelper(root, 0, targetSum)

}

func pathSumHelper(root *TreeNode, pathSum, targetSum int) bool {
	// base case: leaf node
	// check pathSum
	pathSum += root.Val
	if root.Left == nil && root.Right == nil {
		return pathSum == targetSum
	}
	// recursive problem: left/right subtrees
	var lHasPathSum bool = false
	var rHasPathSum bool = false
	if root.Left != nil {
		lHasPathSum = pathSumHelper(root.Left, pathSum, targetSum)
	}
	if root.Right != nil {
		rHasPathSum = pathSumHelper(root.Right, pathSum, targetSum)
	}
	return lHasPathSum || rHasPathSum
}

// @lc code=end

