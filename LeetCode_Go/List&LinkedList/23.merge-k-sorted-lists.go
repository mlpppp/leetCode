/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Merge k sorted linked lists and return it as one sorted linked list.

// solution 2: use heap/priority queue: maintain a min heap of max size k, while heap is not empty poll one node to new list and then push the next node.
// Solution 1: merge lists in two lists' groups. Evolute the working_list of lists until it becomes len==1: merge two lists' group and populates the merged list to the next working_list generation

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	for len(lists) > 1 {
		var mergedLists []*ListNode
		for i := 0; i < len(lists); i += 2 {
			if i+1 >= len(lists) { // odd number of list
				mergedLists = append(mergedLists, lists[i])
				continue
			}
			mergedLists = append(mergedLists, mergeList(lists[i], lists[i+1]))
		}
		lists = mergedLists
	}

	return lists[0]
}

func mergeList(left *ListNode, right *ListNode) *ListNode {
	dummy := &ListNode{-1, nil}
	tail := dummy

	for left != nil && right != nil {
		if left.Val < right.Val {
			tail.Next = left
			tail = left
			left = left.Next
		} else {
			tail.Next = right
			tail = right
			right = right.Next
		}
	}

	if right == nil {
		tail.Next = left
	}
	if left == nil {
		tail.Next = right
	}
	return dummy.Next
}

// @lc code=end

