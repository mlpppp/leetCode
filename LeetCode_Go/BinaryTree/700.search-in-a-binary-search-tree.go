/*
 * @lc app=leetcode id=700 lang=golang
 *
 * [700] Search in a Binary Search Tree
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

// Given a binary search tree (BST) and an integer val. Find the node with the given val in the BST and return the node. If such a node does not exist, return null

// iteration. compare val with the current node, if val is smaller, go left otherwise go right. Hit the null means node does not exist.

// height of the BST, O(log(n)) for balanced BST, worst case O(n)

func searchBST(root *TreeNode, val int) *TreeNode {
	curNode := root
	for curNode != nil {
		if val < curNode.Val {
			curNode = curNode.Left
		} else if val == curNode.Val {
			return curNode
		} else if val > curNode.Val {
			curNode = curNode.Right
		}
	}
	return nil
}

// @lc code=end

