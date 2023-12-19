/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Given the head of a sorted linked list, delete all duplicates.

// two pointers: nextUnique, for traverse the linked list and find the next unique element; cur, for edit the the linked list and link to next nextUnique;

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur, nextUnique := head, head
	for nextUnique != nil {
		if nextUnique.Val != cur.Val {
			// link cur to nextUnique
			cur.Next = nextUnique
			cur = nextUnique // move cur to nextUnique
		}
		nextUnique = nextUnique.Next
	}
	cur.Next = nil
	return head
}

// @lc code=end

