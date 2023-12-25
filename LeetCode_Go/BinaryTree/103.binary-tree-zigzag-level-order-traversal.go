/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
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

// Given the root of a binary tree, return the zigzag level order traversal of its nodes' values. (i.e., from left to right, then right to left for the next level and alternate between).

// BFS, record the level. In different level, alternate between append to result and prepend to result. Pop at the same direction.

// O(n) time, O(n) space.

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	result := make([][]int, 0)
	zigzag := true // alternate between zig and zag

	for len(queue) > 0 {
		var nextLevel []*TreeNode
		var levelResult []int

		for len(queue) > 0 {
			// pop left
			var curNode *TreeNode
			curNode = queue[0]
			queue = queue[1:]

			if curNode.Left != nil {
				nextLevel = append(nextLevel, curNode.Left)
			}
			if curNode.Right != nil {
				nextLevel = append(nextLevel, curNode.Right)
			}

			if zigzag { // zig
				levelResult = append(levelResult, curNode.Val)
			} else { // zag
				levelResult = append([]int{curNode.Val}, levelResult...)
			}
		}
		queue = nextLevel
		result = append(result, levelResult)
		zigzag = !zigzag
	}
	// fmt.Println(result)
	return result
}

// @lc code=end

