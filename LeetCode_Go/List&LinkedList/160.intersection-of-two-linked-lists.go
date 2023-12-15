/*
 * @lc app=leetcode id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Given two acyclic lists, return the node at which the two lists intersect. If the two linked lists have no intersection at all, return null.

// two pointers, same speed, each start from the head of one list, finishing traverse the list, and then start to traverse the other list, (at the end both pointer will traverse each nodes in both lists). If they both reached nil without meeting, there is no intersection. Otherwise, they will meet at the same node, which is the intersection node.

// O(m+n) time, O(1) space

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pointA, pointB := headA, headB
	for pointA != nil || pointB != nil {
		// fmt.Printf("pointA: %v, pointB: %v\n", pointA.Val, pointB.Val)
		if pointA == pointB { //found intersection
			return pointA
		}
		pointA = pointA.Next
		pointB = pointB.Next
		if pointA == nil && pointB != nil {
			pointA = headB
		}
		if pointB == nil && pointA != nil {
			pointB = headA
		}
	}
	return nil
}

// @lc code=end

