/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
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

// Given the root of a binary tree, return the preorder traversal of its nodes' values.

// basic algorithm

// func preorderTraversal(root *TreeNode) []int {
// 	var result []int
// 	preorderTraversalHelper(root, &result)
// 	return result
// }

// func preorderTraversalHelper(root *TreeNode, result *[]int) {
// 	if root == nil {
// 		return
// 	}
// 	*result = append(*result, root.Val)
// 	preorderTraversalHelper(root.Left, result)
// 	preorderTraversalHelper(root.Right, result)
// 	return
// }

// iterative implementation

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result []int
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		// pop the stack, `append` to the result
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)

		// add children from right to left (or left to right for a reversed subtree traversal)
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}

	return result
}

// @lc code=end

