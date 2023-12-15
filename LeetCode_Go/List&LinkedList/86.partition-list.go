/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// O(n)

// Given a linked list's and a value x, rearrange the nodes so that those with values less than x precede nodes with values greater than or equal to x. Maintain the original relative order.

// traverse the linked list; build two solution lists: small and large. When finished traversing, append the large list to the small list.
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	largeDummy := &ListNode{-1, nil}
	largeTail := largeDummy
	smallDummy := &ListNode{-1, nil}
	smallTail := smallDummy

	// traverse
	for head != nil {
		if head.Val < x { // go to small
			smallTail.Next = head
			smallTail = smallTail.Next
		} else { // go to large
			largeTail.Next = head
			largeTail = largeTail.Next
		}
		head = head.Next
	}

	smallTail.Next = largeDummy.Next
	largeTail.Next = nil

	return smallDummy.Next
}

// @lc code=end

