/*
 * @lc app=leetcode id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
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

// Given a BST, find the kth(1-indexed) smallest value in the tree.

// in-order traverse of a BST return an ordered sequence. The kth non-null node we traverse is the kth smallest value. Worst case O(n)

// Follow up: If the BST is modified often (i.e., we can do insert and delete operations) and you need to find the kth smallest frequently, how would you optimize?

// Follow up: Add and maintain the an extra field in each node to store number of nodes in the subtree rooted at that node. And use self-balanced BSTs (AVL trees, red-black trees). This can solve the problem in O(log n) time.

func kthSmallest(root *TreeNode, k int) int {
	var rank int = 0
	var result int
	inorderTraverse(root, &rank, &result, k)
	return result
}

func inorderTraverse(root *TreeNode, rank, result *int, k int) {
	if root == nil {
		return
	}
	inorderTraverse(root.Left, rank, result, k)
	*rank++
	if *rank == k {
		*result = root.Val
		return
	}
	inorderTraverse(root.Right, rank, result, k)
}

// stack implementation

// func kthSmallest(root *TreeNode, k int) int {
// 	var rank int = 0
// 	var stack []*TreeNode
// 	var cur *TreeNode = root
// 	for cur != nil || len(stack) > 0 { // stop when cur==nil and stack is empty
// 		for cur != nil {
// 			stack = append(stack, cur)
// 			cur = cur.Left
// 		}
// 		// pop stack and process node
// 		cur = stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]
// 		rank++
// 		fmt.Printf("rank: %d, value: %d\n", rank, cur.Val)
// 		if rank == k {
// 			return cur.Val
// 		}
// 		// finish in-order, go right
// 		cur = cur.Right
// 	}
// 	return 0
// }

// @lc code=end

