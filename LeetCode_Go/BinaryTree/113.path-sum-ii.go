/*
 * @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
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

// Given a binary tree, find paths (might be multiple) that satisfy that pathSum == targetSum. pathSum: sum of a root-to-leaf paths. Return paths as a list of the node values, not node references.

// We need to reach every leaf node (all paths) to find all pathSum satisfactory paths. => O(n). Use an array to store the current path, add current node in preorder place and remove it in postorder place. When reaches a leaf, determine if it satisfies the pathSum, if so create a copy of the path to the result.

// O(n) need to traverse all nodes.

func pathSum(root *TreeNode, targetSum int) [][]int {
	// edge case: empty tree
	if root == nil {
		return [][]int{}
	}

	var ansPaths [][]int // aggregate target paths
	var path []int       // find target paths
	pathSumHelper(root, 0, targetSum, path, &ansPaths)
	return ansPaths
}

func pathSumHelper(root *TreeNode, pathSum, targetSum int, path []int, ansPaths *[][]int) {
	// preorder processing
	pathSum += root.Val
	path = append(path, root.Val)

	// base case: leaf node
	// check targetSum, aggregate ans if satisfied
	if root.Left == nil && root.Right == nil {
		if pathSum == targetSum {
			*ansPaths = append(*ansPaths, append([]int{}, path...)) // create a copy
		}
	}

	// recursive problem: left/right subtrees
	if root.Left != nil {
		pathSumHelper(root.Left, pathSum, targetSum, path, ansPaths)

	}
	if root.Right != nil {
		pathSumHelper(root.Right, pathSum, targetSum, path, ansPaths)
	}

	// postorder place processes: recover pathSum and path
	// subtract pathSum
	pathSum -= root.Val
	// pop path
	path = path[:len(path)-1]
	return
}

// @lc code=end

