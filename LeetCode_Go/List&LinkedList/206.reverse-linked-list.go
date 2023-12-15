/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// reverse a linked list

// iterative approach: prev pointer one step before cur pointer. cur point rewire the Next filed to the prev pointer, and then move the prev and cur pointer one step forward.

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev, cur *ListNode
	prev, cur = nil, head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

// recursive approach: base case: last node. return: reversed list' head. operator: reverse the direction of the current node, and return the reversed list' head.
func reverseLinkedList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	reversedHead := reverseLinkedList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return reversedHead
}

// @lc code=end

