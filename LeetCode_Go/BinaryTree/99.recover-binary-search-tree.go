/*
 * @lc app=leetcode id=99 lang=golang
 *
 * [99] Recover Binary Search Tree
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

// Given a binary search tree, where the values of exactly two nodes of the tree were swapped by mistake. Recover the tree without changing its structure.

// In-order traversal should return ordered sequence. But the swapped nodes should be out of order. Attempting to identify the two swapped nodes in one in-order traversal. There are two kinds of swapping: adjacent swap (1,3,2,4; 3 and 2), create one out-of-order pair (3,2); or non-adjacent swap (3,2,1; 3 and 1), creates two out-of-order pairs (3,2) and (2,1). We don't know which case when first meeting a out-of-order pair, so we may assume it is the first case. but if we meet a second out-of-order pairs, we know it is the second case, and should update the result.

var lastNode *TreeNode
var swapNodeA *TreeNode
var swapNodeB *TreeNode

func recoverTree(root *TreeNode) {
	lastNode = nil
	swapNodeA = nil
	swapNodeB = nil
	inorderTraverseHelper(root)
	// fmt.Printf("swapNodeA: %d, swapNodeB: %d\n", swapNodeA, swapNodeB)
	swapNodeA.Val, swapNodeB.Val = swapNodeB.Val, swapNodeA.Val
}

func inorderTraverseHelper(root *TreeNode) {
	if root == nil {
		return
	}
	inorderTraverseHelper(root.Left)
	// fmt.Printf("lastNode: %d, curNode: %d\n", lastNode, root)
	if lastNode != nil && root.Val < lastNode.Val { // lastNode is nil until visited the first node
		if swapNodeA == nil {
			swapNodeA = lastNode
			swapNodeB = root // it is possible that current node is also the swapNode, if not, the next encountered swapNode will overwrite it
		} else {
			swapNodeB = root
		}
		// fmt.Printf("swapNodeA: %d, swapNodeB: %d\n", swapNodeA, swapNodeB)
	}
	lastNode = root
	// fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	inorderTraverseHelper(root.Right)
}

// @lc code=end

