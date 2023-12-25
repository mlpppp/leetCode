/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
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

// Given a binary tree, return the level order traversal of its nodes' values. (group values level by level)

// BFS

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var queue []*TreeNode
	var result [][]int

	// initial values
	queue = append(queue, root)
	// var level int = 0

	// iterate till queue is empty
	for len(queue) > 0 {
		var nextLvCandidates []*TreeNode
		var curLevelResult []int
		for len(queue) > 0 {
			curNode := queue[0]
			curLevelResult = append(curLevelResult, curNode.Val)
			// append non null nodes to candidate
			if curNode.Left != nil {
				nextLvCandidates = append(nextLvCandidates, curNode.Left)
			}
			if curNode.Right != nil {
				nextLvCandidates = append(nextLvCandidates, curNode.Right)
			}
			queue = queue[1:]
		}
		// for i := 0; i < len(nextLvCandidates); i++ {
		// 	fmt.Printf("%d ", nextLvCandidates[i].Val)
		// }
		// fmt.Println()
		// fmt.Println(curLevelResult)
		// fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
		// level++
		result = append(result, curLevelResult)
		queue = nextLvCandidates
	}
	return result
}

// @lc code=end

