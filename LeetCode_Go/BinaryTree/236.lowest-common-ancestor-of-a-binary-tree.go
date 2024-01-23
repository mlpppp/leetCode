/*
 * @lc app=leetcode id=236 lang=golang
 *
 * [236] Lowest Common Ancestor of a Binary Tree
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

// Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

// Implement DFS to find either of the two nodes in left and right subtrees. As a root, if I can find both nodes in my left and right subtree, I am the LCA, return me as result. If only one nodes is found in my subtrees: either the nodes is the LCA (one or both nodes are hidden in the tree owned by LCA); or the LCA is in levels above me. Either way, I would return the found node to let functions in higher call stack to decide.

// 重点是lca和pq的处理逻辑是一样的,都是把自己返回回去。因为我们要找的是最上层的LCA。

// O(n) time, O(n) space

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	leftFound := lowestCommonAncestor(root.Left, p, q)
	rightFound := lowestCommonAncestor(root.Right, p, q)
	if leftFound != nil && rightFound != nil {
		return root
	}
	if leftFound != nil {
		return leftFound
	}
	if rightFound != nil {
		return rightFound
	}
	return nil // LCA is not in this subtree

}

// @lc code=end

