/*
 * @lc app=leetcode id=95 lang=golang
 *
 * [95] Unique Binary Search Trees II
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

// Given an integer n, return all the structurally unique BST's, which has exactly n nodes of unique values from 1 to n in an array.

// recursive function build all possible BSTs from continuous valued node from `start` to `end`. For any possible `root`, we build all possible BSTs for the left subtree and the right subtree, and then do a combination to connect the root and the left/right subtrees.
// rewrite to DP: use a read-only DP table to store the result (`templates`), and use a function to copy trees from the DP table.

// O(n^2)
func generateTrees(n int) []*TreeNode {
	subStructures := make(map[[2]int][]*TreeNode)
	return generateTreesHelper(1, n, subStructures)
}

func generateTreesHelper(start, end int, subStructures map[[2]int][]*TreeNode) []*TreeNode {
	// base case
	if start > end {
		return []*TreeNode{nil}
	}

	var result []*TreeNode

	// check DP table
	key := [2]int{start, end}
	if value, exists := subStructures[key]; exists {
		for _, root := range value {
			result = append(result, cloneTree(root))
		}
		return result
	}

	// iterate subtrees with mid as root
	for mid := start; mid <= end; mid++ {
		leftStructures := generateTreesHelper(start, mid-1, subStructures)
		rightStructures := generateTreesHelper(mid+1, end, subStructures)

		// combination, connecting sub-structures
		for i := 0; i < len(leftStructures); i++ {
			for j := 0; j < len(rightStructures); j++ {
				root := &TreeNode{Val: mid}
				root.Left = leftStructures[i]
				root.Right = rightStructures[j]
				result = append(result, root)
			}
		}
	}
	subStructures[key] = result
	return result

}

func cloneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	newRoot := &TreeNode{Val: root.Val}
	newRoot.Left = cloneTree(root.Left)
	newRoot.Right = cloneTree(root.Right)
	return newRoot
}

// @lc code=end

