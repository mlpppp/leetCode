/*
 * @lc app=leetcode id=450 lang=golang
 *
 * [450] Delete Node in a BST
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

// Given a BST and a key, find and delete the node with the given key in the BST. After the deletion return the root of the new BST.

// Delete the node: find the maximun node in the left subtree (rightmost), or the minimum node in the right subtree (leftmost) as substitution, depending on the left/right subtree availability.

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	// fmt.Printf("root.Val: %d\n", root.Val)
	if root.Val == key {
		// fmt.Println("found node")
		// case 1: single node (leaf node)
		if root.Left == nil && root.Right == nil {
			return nil
		}
		// case 2: single subtree
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// case 3: two subtrees, find substitution in the Right subtree (argmin(root.Right))
		if root.Right != nil && root.Left != nil {
			newRoot := root.Right
			for newRoot.Left != nil {
				newRoot = newRoot.Left
			}
			// fmt.Printf("newRoot.Val: %d\n", newRoot.Val)
			root.Right = deleteNode(root.Right, newRoot.Val) // recursively delete the substitution node in the original right subtree
			newRoot.Left = root.Left
			newRoot.Right = root.Right
			root = newRoot
		}

	} else if root.Val > key {
		newRoot := deleteNode(root.Left, key)
		root.Left = newRoot
	} else if root.Val < key {
		newRoot := deleteNode(root.Right, key)
		root.Right = newRoot
	}
	return root
}

// @lc code=end

