/*
 * @lc app=leetcode id=908 lang=golang
 *
 * [876] Middle of the Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Given a linked list, return a middle node of linked list.

// fast and slow pointers. both start from the head. fast pointer 2* speed. when fast pointer reaches the end of the linked list, slow pointer reaches the middle of the linked list
func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// @lc code=end

