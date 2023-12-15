/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{-1, nil}
	dummy.Next = head
	tail := dummy
	for i := 0; i < n+1; i++ {
		tail = tail.Next
	}
	predelete := dummy
	for tail != nil {
		tail = tail.Next
		predelete = predelete.Next
	}
	// now predelete find next node to be the deleted node
	delete := predelete.Next
	// delete the nth node
	predelete.Next = delete.Next
	delete = nil

	return dummy.Next
}

// @lc code=end

