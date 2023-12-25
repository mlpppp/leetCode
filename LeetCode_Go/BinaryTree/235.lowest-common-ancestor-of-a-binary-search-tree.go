/*
 * @lc app=leetcode id=235 lang=golang
 *
 * [235] Lowest Common Ancestor of a Binary Search Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

// Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.

// three possibilities: p, q < root, then find LCA in the left subtree; p, q > root, then find LCA in the right subtree; root = one of p,q or root in the middle of p, q, return root

// O(logn)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// make sure p < q
	if p.Val > q.Val {
		p, q = q, p
	}
	// case 1: p < q < root
	if q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if root.Val < p.Val { // case 2: root < p < q
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return root
	}
}

// @lc code=end

