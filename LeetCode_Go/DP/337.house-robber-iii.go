/*
 * @lc app=leetcode id=337 lang=golang
 *
 * [337] House Robber III
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

// 198 house robber in a binary tree, two house is adjacent if there are directly linked in the tree. Entry points is the root of the tree. (notes that the rob dont need to follow one path, rob may rob every node in a level of the tree)

// subproblem: robSubtree(root *TreeNode) int best result rob from a subtree. Goal: robSubtree(root)

// choices: {
// 	rob this root and go 2 steps: left/right subtree of the left/right subtree.
// 	dont rob this root and go 1 step: left/right subtree.
// }

// transition: robSubtree(root) = {
// 	max(
// 		root.Val + robSubtree(root.Left.Left) + robSubtree(root.Left.Right) + robSubtree(root.Right.Left) + robSubtree(root.Right.Right),
// 		robSubtree(root.Left) + robSubtree(root.Left)
// 	)
// }

func rob(root *TreeNode) int {
	// DP table, map[TreeNode]int
	table := make(map[*TreeNode]int)
	return robSubtree(root, table)
}

func robSubtree(root *TreeNode, table map[*TreeNode]int) int {
	// base case: empty tree
	if root == nil {
		return 0
	}
	// query DP table
	if val, exists := table[root]; exists {
		return val
	}
	// recursion
	robThisReturn := root.Val
	if root.Left != nil {
		robThisReturn += robSubtree(root.Left.Left, table) + robSubtree(root.Left.Right, table)
	}
	if root.Right != nil {
		robThisReturn += robSubtree(root.Right.Left, table) + robSubtree(root.Right.Right, table)
	}

	skipThisReturn := robSubtree(root.Left, table) + robSubtree(root.Right, table)

	table[root] = max(robThisReturn, skipThisReturn)
	return table[root]
}

// @lc code=end

