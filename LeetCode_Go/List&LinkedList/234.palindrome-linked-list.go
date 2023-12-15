/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Given a linked list, determine if it is a palindrome

// Brute Force: copy the linked list into an array, and determine if the array is a palindrome./ reversely traversed the linked list and see if it outputs the same value as forward direction.

// Optimal: reverse the second half of the linked list (toward the center). Two pointers from both ends traverse the linked list to the center and determine if it is a palindrome. Then reverse back the second half of the linked list.

// O(n) time | O(1) space

func isPalindrome(head *ListNode) bool {
	// find the center or the second middle node(if even nodes)
	secondHalfStart, temp := head, head
	for temp != nil && temp.Next != nil {
		secondHalfStart = secondHalfStart.Next
		temp = temp.Next.Next
	}

	// reverse the second half
	reversedHead := reverseList(secondHalfStart)

	// determine Palindrome
	p1, p2 := head, reversedHead
	for p1 != nil && p2 != nil {
		// fmt.Printf("p1: %d, p2: %d\n", p1.Val, p2.Val)
		if p1.Val != p2.Val {
			// reverse back the linked list (optional)
			// _ = reverseList(reversedHead)
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	// reverse back the linked list (optional)
	// _ = reverseList(reversedHead)
	return true

}

func reverseList(head *ListNode) *ListNode {
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

// @lc code=end

