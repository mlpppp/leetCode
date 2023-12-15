/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Merge two sorted linked lists and return it as a sorted list.

// traverse two lists; Two pointers with steps of 1, compare the values and append to the new list. when one pointer reaches the end, append the other list to the new list.

// O(n)

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// dummy always point to the start of solution
	dummy := &ListNode{-1, nil}
	// tail point to the end of current solution
	tail := dummy

	// not nil list
	for list1 != nil && list2 != nil {

		if list1.Val < list2.Val { // 1 is smaller
			tail.Next = list1
			list1 = list1.Next
		} else { // 2 is smaller
			tail.Next = list2
			list2 = list2.Next
		}

		tail = tail.Next
	}

	if list1 == nil {
		tail.Next = list2
	}
	if list2 == nil {
		tail.Next = list1
	}

	return dummy.Next
}

// [1,2,4]\n[1,3,4]
// @lc code=end

