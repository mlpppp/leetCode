/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Reverse every k nodes in a linked list, preserving the original order of any remaining nodes.


// iterative solution. traverse the linked list and reverse the nodes in k groups. use a utility function "reverseList" to reverse a group, then reconnect it from/to the previous and next group.



func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{-1, head}
	prev := dummy // a node previous to current group
	groupStart := head
	for groupExists(groupStart, k) {
		// printList(dummy.Next)
		// gather pointers
		groupEnd := groupStart
		for i := 0; i < k-1; i++ {
			groupEnd = groupEnd.Next
		}
		// fmt.Printf("prev:%v, groupStart:%v, groupEnd:%v \n", prev.Val, groupStart.Val, groupEnd.Val)

		nextGroupStart := groupEnd.Next

		// reverse group
		groupEnd.Next = nil
		reversedHead := reverseList(groupStart)

		// reconnect
		prev.Next = reversedHead         // link from the prevGroup
		groupStart.Next = nextGroupStart // link to the next group

		// move to next group
		prev = groupStart
		groupStart = nextGroupStart

	}
	return dummy.Next
}

func groupExists(head *ListNode, k int) bool {
	if head == nil {
		return false
	}
	for i := 0; i < k-1; i++ {
		head = head.Next
		if head == nil {
			return false
		}
	}
	return true
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev, curr, next *ListNode
	prev = nil
	curr = head
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// @lc code=end

