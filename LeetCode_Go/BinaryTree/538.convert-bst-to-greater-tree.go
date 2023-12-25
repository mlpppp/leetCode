/*
 * @lc app=leetcode id=538 lang=golang
 *
 * [538] Convert BST to Greater Tree
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
// Convert the given Binary Search Tree (BST) into a Greater Tree by updating each node's key to the sum of its original key and the sum of all keys greater than it in the BST.

// reversed inorder traverse a BST return an reversed-ordered traverse sequence, fitting the problem definition (O(n) time)

// brute force: inorder traverse the whole tree construct ordered array, then consutruct array of new values from the ordered array. Then traverse the tree again to assign the new values. (O(n) time)

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	reversedInorderTraverse(root, &sum)
	return root
}

func reversedInorderTraverse(root *TreeNode, sum *int) {
	if root == nil {
		return
	}

	reversedInorderTraverse(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	reversedInorderTraverse(root.Left, sum)
}

// stack implementation
// func convertBST(root *TreeNode) *TreeNode {
// 	cur := root
// 	stack := make([]*TreeNode, 0)
// 	sum := 0
// 	for cur != nil || len(stack) != 0 {
// 		for cur != nil {
// 			stack = append(stack, cur)
// 			cur = cur.Right
// 		}
// 		// pop and process node
// 		node := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]
// 		sum += node.Val
// 		node.Val = sum
// 		cur = node.Left
// 	}
// 	return root
// }

// @lc code=end

