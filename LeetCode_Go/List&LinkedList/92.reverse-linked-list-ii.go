/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// reverse a slice of linked list from the inputed left position to right position. Retain the order of the other parts.

// find the pointer to the positions: node preceding the left position; left position; node following the right position; right position. Use reverse utility function. then reconnect the parts preceeding the the left position and the part following the the right position.

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{-1, head}
	preLeftNode := dummy
	rightNode := dummy
	for i := 0; i < right; i++ {
		rightNode = rightNode.Next
	}
	for i := 0; i < left-1; i++ {
		preLeftNode = preLeftNode.Next
	}
	leftNode := preLeftNode.Next
	rightNodeNext := rightNode.Next

	// partial reverse
	rightNode.Next = nil
	reversedHead := reverseLinkedList(preLeftNode.Next)

	// reconnect
	preLeftNode.Next = reversedHead
	leftNode.Next = rightNodeNext
	return dummy.Next
}

// reverse linked list with iterative approach
func reverseLinkedList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev := nil
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// @lc code=end

