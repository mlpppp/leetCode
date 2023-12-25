/*
 * @lc app=leetcode id=173 lang=golang
 *
 * [173] Binary Search Tree Iterator
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

// https://www.youtube.com/watch?v=RXy5RzGF5wo

// Implement the BSTIterator class that represents an iterator over the in-order traversal of a binary search tree (BST)

// use inorder traverse iterative implementation. Break the loop in the algorithm into interator function calls.

// Alt simple solution: O(n)/O(1): iterate BST inorder, save results in an array. implement iterator in the array (trivial).

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	stack := []*TreeNode{}
	curNode := root
	for curNode != nil {
		stack = append(stack, curNode)
		curNode = curNode.Left
	}
	return BSTIterator{stack: stack}
}

func (this *BSTIterator) Next() int {
	// pop stack
	curNode := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	nextVal := curNode.Val
	// go to right subtree and append along the left path of the subtree
	curNode = curNode.Right
	for curNode != nil {
		this.stack = append(this.stack, curNode)
		curNode = curNode.Left
	}
	return nextVal
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

// reference: inorder traverse iterative implementation

// func inorder(root *TreeNode) {
// 	if root == nil {
// 		return
// 	}
// 	stack := []*TreeNode{}
// 	curNode := root
// 	for curNode != nil || len(stack) > 0 {
// 		// append left path
// 		for curNode != nil {
// 			stack = append(stack, curNode)
// 			curNode = curNode.Left
// 		}
// 		// ! pop
// 		curNode = stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]

// 		// go to right subtree
// 		curNode = curNode.Right
// 	}
// }

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end

