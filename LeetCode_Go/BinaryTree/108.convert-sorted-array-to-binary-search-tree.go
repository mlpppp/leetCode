/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
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

// Given an integer array nums where the elements are sorted in ascending order, convert it to a height-balanced binary search tree.

// use the middle of the array (or the first middle element) as the root. Then recursively build the left/right tree with left and right subarrays separated by the middle element. (empty subarrays are represented as nil)

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := (len(nums) - 1) / 2
	// fmt.Printf("cur index: %d\n", mid)
	// fmt.Printf("cur num: %v\n", nums)

	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}

// @lc code=end

