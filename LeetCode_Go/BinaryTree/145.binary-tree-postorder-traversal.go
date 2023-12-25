/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
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

// Given the root of a binary tree, return the postorder traversal of its nodes' values.

// basic algorithm

// func postorderTraversal(root *TreeNode) []int {
// 	var result []int
// 	postorderTraversalHelper(root, &result)
// 	return result
// }
// func postorderTraversalHelper(root *TreeNode, result *[]int) {
// 	if root == nil {
// 		return
// 	}
// 	postorderTraversalHelper(root.Left, result)
// 	postorderTraversalHelper(root.Right, result)
// 	*result = append(*result, root.Val)
// 	return
// }

// iterative implementation

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result []int
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		// pop the stack, `prepend` to the result
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append([]int{cur.Val}, result...)

		// add children from left to right (or right to left for reversed subtree traversal)
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
	}
	return result
}

// @lc code=end

