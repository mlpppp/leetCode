/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
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

// implements inorder traversal to a binary tree
// func inorderTraversal(root *TreeNode) []int {
// 	var result []int
// 	inorderTraversalHelper(root, &result)
// 	return result
// }

// func inorderTraversalHelper(root *TreeNode, result *[]int) {
// 	if root == nil {
// 		return
// 	}
// 	inorderTraversalHelper(root.Left, result)
// 	*result = append(*result, root.Val)
// 	inorderTraversalHelper(root.Right, result)
// }

// iterative implementation: 1. push nodes along left path to stack 2. pop the stack, append to the result 3. go to the right subtree

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	stack := []*TreeNode{}
	var cur *TreeNode = root

	for len(stack) > 0 || cur != nil {
		// push nodes along left path
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		// pop the stack, `append` to the result
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		// go to right subtree
		cur = cur.Right
	}
	return result
}

// @lc code=end

