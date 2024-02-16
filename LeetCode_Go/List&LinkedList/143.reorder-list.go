/*
 * @lc app=leetcode id=143 lang=golang
 *
 * [143] Reorder List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// You are given the head of a singly linked-list. The list can be represented as: L0 → L1 → … → Ln - 1 → Ln
// Reorder the list to be on the following form: L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → … You may not modify the values in the list's nodes.

// reverse the second half of the list (middle or the first node in the second half), then two pointers, from two side to the middle, we build the solution.
// 1. find the half: fast/slow pointers
// 2. reverse(mid)
// 3. build the solution

func reorderList(head *ListNode) {
	// find the middle or the first node in the second half
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// reverse the second half
	p2 := reverse(slow)

	// build the solution
	p1 := head
	for p2 != nil { // odd or even case both terminate with p2 == nil
		p1Next := p1.Next
		p2Next := p2.Next
		p1.Next = p2
		p2.Next = p1Next
		p1 = p1Next
		p2 = p2Next
	}
	if p1 != nil { // in even case p1 might point the last node which point to itself
		p1.Next = nil
	}
}

func reverse(head *ListNode) *ListNode {
	var cur, prev *ListNode
	cur, prev = head, nil
	for cur != nil {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}
	return prev
}

// @lc code=end

