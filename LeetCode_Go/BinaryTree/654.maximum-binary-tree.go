/*
 * @lc app=leetcode id=654 lang=golang
 *
 * [654] Maximum Binary Tree
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

// Given an array of unique integers, construct a maximum binary tree using the following algorithm: create a root node with the maximum value in the array, then recursively build left and right subtrees from the subarrays to the left and right of the maximum value. Return the resulting maximum binary tree.

// follow the algorithm

// O(nlogn) time and O(n) space

func constructMaximumBinaryTree(nums []int) *TreeNode {
	root := constructMaximumBinaryTreeBetween(nums, 0, len(nums)-1)
	return root
}

func constructMaximumBinaryTreeBetween(nums []int, i, j int) *TreeNode {
	idx := findMaximumIdxBetween(nums, i, j)
	if idx == -1 {
		return nil
	}

	root := TreeNode{
		Val:   nums[idx],
		Left:  constructMaximumBinaryTreeBetween(nums, i, idx-1),
		Right: constructMaximumBinaryTreeBetween(nums, idx+1, j),
	}
	return &root
}

func findMaximumIdxBetween(nums []int, i, j int) int {
	if i > j {
		return -1
	}
	max := -1
	idx := -1
	for cur := i; cur <= j; cur++ {
		if nums[cur] > max {
			max = nums[cur]
			idx = cur
		}
	}
	return idx
}

// @lc code=end

