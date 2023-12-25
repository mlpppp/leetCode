/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Given a binary search tree (BST) and a value to insert into the tree, Return the root node of the BST after the insertion of the value. It is guaranteed that the new value does not exist in the original BST.

// Insert the node as a leaf. similar to search. O(log n)

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode(val, nil, nil)
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	} else if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}