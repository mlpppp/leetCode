/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
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

// Given a binary tree, find its minimum depth, aka shortest path from the root node down to the nearest leaf. (a leaf's depth is 1)

// (Top-Down) to avoid going to deep with necessity, increase the probe range: a node check if it and its children is a possible minDepth. If any one of the left/right children is a leaf, update the globalMinDepth and return. Otherwise recursivly check the left/right subtrees

// alt solution: BFS, whenever a node is discovered as a leaf during BFS, the minDepth the is current traverse level

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	var globalMinDepth int = math.MaxInt32
	minDepthHelper(root, 1, &globalMinDepth)
	return globalMinDepth

}

func minDepthHelper(root *TreeNode, depth int, globalMinDepth *int) {
	// fmt.Printf("root: %d, depth: %d, globalMinDepth: %d\n", root.Val, depth, *globalMinDepth)

	// if any one of the left/right children is a leaf, update the globalMinDepth and return
	if (root.Left != nil && root.Left.Left == nil && root.Left.Right == nil) || (root.Right != nil && root.Right.Left == nil && root.Right.Right == nil) {
		if *globalMinDepth > (depth + 1) {
			*globalMinDepth = depth + 1
		}
		return
	}

	// if not, recursivly find in the left/right subtrees
	if root.Left != nil {
		minDepthHelper(root.Left, depth+1, globalMinDepth)
	}
	if root.Right != nil {
		minDepthHelper(root.Right, depth+1, globalMinDepth)
	}
}

// @lc code=end

