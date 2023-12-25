/*
 * @lc app=leetcode id=109 lang=golang
 *
 * [109] Convert Sorted List to Binary Search Tree
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Given the head of a singly linked list where elements are sorted in ascending order, convert it to a height-balanced binary search tree.

// use fast/slow pointer to find the middle, or the second middle element. Build root with middle elment, and use middle as the new boundary for the left and the right subtrees. (hard to implement, boundary values pitfalls...)

// solution 2: traverse the list and build an array, then treat it as problem 108

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	return sortedListToBSTHelper(head, nil)

}

func sortedListToBSTHelper(head, tail *ListNode) *TreeNode {
	if head == tail {
		return nil
	}

	mid, fast := head, head                 // mid, slow pointer, search for the middle element
	for fast != tail && fast.Next != tail { // when even number, slow will be the second middle element
		mid = mid.Next
		fast = fast.Next.Next
	}

	root := &TreeNode{Val: mid.Val}
	root.Left = sortedListToBSTHelper(head, mid) // by definition not include the tail.
	root.Right = sortedListToBSTHelper(mid.Next, tail)
	return root
}

// @lc code=end

