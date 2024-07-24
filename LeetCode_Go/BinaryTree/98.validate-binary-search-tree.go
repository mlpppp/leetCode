/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
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

// Given a binary tree, determine if it is a valid binary search tree (BST).

// traverse the tree and see if any node break the BST rule：max(leftTree) < root.Val < min(rightTree).

func isValidBST(root *TreeNode) bool {
	_, _, isBST := isValidBSTHelper(root)
	// fmt.Printf("%v\n", isBST)
	return isBST
}

func isValidBSTHelper(root *TreeNode) (min, max *TreeNode, isBST bool) {
	// base case: leaf node
	if root.Left == nil && root.Right == nil {
		return root, root, true
	}

	min, max = root, root // minimum and maximum node in the entire tree
	if root.Left != nil {
		lmin, lmax, isBST := isValidBSTHelper(root.Left)
		if !isBST || root.Val <= lmax.Val { // early termination if not a BST
			return nil, nil, false
		}
		min = lmin // update the minimum node
	}

	if root.Right != nil {
		rmin, rmax, isBST := isValidBSTHelper(root.Right)
		if !isBST || root.Val >= rmin.Val { // early termination if not a BST
			return nil, nil, false
		}
		max = rmax // update the maximum node
	}

	// return min and max of the entire tree where root as the root
	return min, max, true
}

// boundary based solution: Boundary:= an interval that subtree values must be inside. Top-down traversal(pre-order), see if any node break the boundary. Starting slack, dynammically restrict the boundary.
// func isValidBST(root *TreeNode) bool {
// 	isBST := isValidBSTHelper(root, -int(math.Pow(2.0, 32.0)), int(math.Pow(2.0, 32.0)))
// 	return isBST
// }

// func isValidBSTHelper(root *TreeNode, minBound, maxBound int) bool {
// 	if root == nil {
// 		return true
// 	}
// 	if root.Val <= minBound || root.Val >= maxBound {
// 		return false
// 	}
// 	return isValidBSTHelper(root.Left, minBound, root.Val) && isValidBSTHelper(root.Right, root.Val, maxBound)
// }

// @lc code=end

