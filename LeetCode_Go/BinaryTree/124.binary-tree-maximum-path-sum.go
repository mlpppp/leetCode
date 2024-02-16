/*
 * @lc app=leetcode id=124 lang=golang
 *
 * [124] Binary Tree Maximum Path Sum
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

// A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them. A path does not need to pass through the root. The path sum of a path is the sum of the node's values in the path. Given the root of a binary tree, return the maximum path sum.

// A path must have a highest node := the root of the path
// A path must have two downward side heads (could be size 0): we calculate each head's path sum `maxPathSumSingle`

// self.maxPathSumSingle = 0 if self is null
// leftSum = self.Left.maxPathSumSingle if it > 0 else 0
// RightSum = self.Right.maxPathSumSingle if it > 0 else 0
// self.maxPathSumSingle = max(leftSum, RightSum) + self.Val
// maxPathSum = max(maxPathSum, leftSum + RightSum + self.Val)

func maxPathSum(root *TreeNode) int {
	ans := math.MinInt
	DFS(root, &ans)
	return ans
}

// return maxPathSumSingle
func DFS(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}

	leftSum := max(0, DFS(root.Left, ans))
	RightSum := max(0, DFS(root.Right, ans))

	maxPathSumSingle := max(leftSum, RightSum) + root.Val

	*ans = max(*ans, leftSum+RightSum+root.Val)

	return maxPathSumSingle
}

// @lc code=end

