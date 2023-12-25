/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
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

// Given the root of a binary tree, return the bottom-up level order traversal of its nodes' values. (i.e., from left to right, level by level from leaf to root).

// BFS, record the level. But instead of append, prepend result of level to the overall result. (you can also reverse the standard BFS result).

// O(n) time, O(n) space.

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	var result [][]int
	for len(queue) > 0 {
		var nextLevel []*TreeNode
		var curResult []int
		for len(queue) > 0 {
			// pop
			curNode := queue[0]
			queue = queue[1:]
			curResult = append(curResult, curNode.Val)
			// add to next level
			if curNode.Left != nil {
				nextLevel = append(nextLevel, curNode.Left)
			}
			if curNode.Right != nil {
				nextLevel = append(nextLevel, curNode.Right)
			}
		}
		queue = nextLevel
		result = append([][]int{curResult}, result...)
	}
	return result
}

// @lc code=end

