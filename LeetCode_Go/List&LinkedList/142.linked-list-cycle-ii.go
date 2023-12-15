/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// given a linked list, return the node "pos" where the cycle begins. pos is relative to the head(pos=0). If there is no cycle, return pos=-1

// fast and slow pointers both start from the head. fast pointer 2* speed. No circle if any of the pointers reaches null. Otherwise they will meet. When two pointers meets for the first time, put a pointer at the head and start traverse in parallel, same speed, with a pointer at the meeting point. The next meeting point will be the cycle start point.

// O(n)

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil
}

// @lc code=end

