/*
 * @lc app=leetcode id=572 lang=golang
 *
 * [572] Subtree of Another Tree
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

// Given the roots of two binary trees root and subRoot, return true if there is a subtree of root with the same structure and node values of subRoot and false otherwise. (note: A tree could also be considered as a subtree of itself.)

// (brute force O(n*m)) DFS root, when there is a node of same value as subRoot, Determine if they are the same structure via a DFS (same val, same left tree, same right tree).

// better name: containSubtree
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// edge case: a null tree is subtree of any tree
	if subRoot == nil {
		return true
	} else if root == nil {
		return false
	}

	// preorder: compare two trees
	if root.Val == subRoot.Val && isSameStructure(root, subRoot) {
		return true
	}

	// DFS
	if isSubtree(root.Left, subRoot) {
		return true
	}
	if isSubtree(root.Right, subRoot) {
		return true
	}
	return false
}

// compare any tree root to a reference tree ref. O(m), return true if they match
func isSameStructure(root *TreeNode, ref *TreeNode) bool {
	if root == nil && ref == nil {
		return true
	} else if root == nil && ref != nil || root != nil && ref == nil {
		return false
	} else {
		if root.Val != ref.Val {
			return false
		} else {
			return isSameStructure(root.Left, ref.Left) && isSameStructure(root.Right, ref.Right)
		}
	}
}

// @lc code=end

