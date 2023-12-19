/*
 * @lc app=leetcode id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
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

// Flatten a binary tree into a linked list using its pre-order traversal, where the linked list uses the TreeNode class with Right child pointing to the next node and Left child always null.

// DFS, post-order traversal. Each node receives a flattened left tree and a flattened right tree, with pointers to the last node (tails) in the linked lists. Then append the left linked list to the Right of root, and append the right linked list to the last node in the left linked list. Return the correct tails from leftTail, rightTail or root.

// O(n) time, O(n) space.

func flatten(root *TreeNode) {
	_ = flattenAndReturnTail(root)
	return
}

func flattenAndReturnTail(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	leftTail := flattenAndReturnTail(root.Left)
	rightTail := flattenAndReturnTail(root.Right)

	left := root.Left
	right := root.Right

	root.Left = nil

	// if there is a left tree, append the right flattened tree to the leftTail, and reconnect the root.Right to the left tree
	if left != nil {
		leftTail.Right = right
		root.Right = left
	}

	// if there is a right tree, return the right tree's tail as rightTail, otherwise return leftTail
	if rightTail != nil {
		return rightTail
	} else if leftTail != nil {
		return leftTail
	} else {
		return root
	}
}

// @lc code=end

