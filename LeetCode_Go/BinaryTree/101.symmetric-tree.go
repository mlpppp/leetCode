/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
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

// Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

// iterative solution (BFS): BFS traverse two subtrees of root. For each level, the left subtree from left to right, and the right subtree from right to left. See if the current node are the same.

// recursive (preorder): DFS traverse two subtrees of root. The left subtree: root -> root.Left -> root.Right. The right subtree: root -> root.Right -> root.Left.

// recursive (from definition): to be symmetric, 1. the leftRoot and the rightRoot should be the same. 2. leftRoot.Left and rightRoot.Right should be symmetric. 3. leftRoot.Right and rightRoot.Left should be symmetric

func isSymmetric(root *TreeNode) bool {
	return isSymmetricHelper(root.Left, root.Right)
}

func isSymmetricHelper(left, right *TreeNode) bool {
	// case 1. when there exists a nil in either or both tree
	if left == nil {
		return right == nil
	}
	if right == nil {
		return left == nil
	}

	if left.Val != right.Val {
		return false
	}

	return isSymmetricHelper(left.Left, right.Right) && isSymmetricHelper(left.Right, right.Left)
}

// @lc code=end

