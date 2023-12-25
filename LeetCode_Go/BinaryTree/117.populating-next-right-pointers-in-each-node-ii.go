/*
 * @lc app=leetcode id=117 lang=golang
 *
 * [117] Populating Next Right Pointers in Each Node II
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
// Given a binary tree, populate a new field `Next` in every node pointing to the next node in the same level (or Null if there is no next node).

// Iterative solution (BFS). Assuming a level is fully connected, we traverse the level with Next pointer(as if it is a linked list), build an array of all non-nil children in the next level, and populate the Next pointer for the next level . The entry point to the next level is the first non-null child pointer of the current level.

// O(n) time and O(n) space.

// func connect(root *Node) *Node {
// 	if root == nil {
// 		return nil
// 	}
// 	// initialize
// 	curNode := root

// 	// nextCurNode: first non-nul child of root
// 	var nextCurNode *Node = nil
// 	if curNode.Left != nil {
// 		nextCurNode = curNode.Left
// 	} else if curNode.Right != nil {
// 		nextCurNode = curNode.Right
// 	}

// 	for nextCurNode != nil {
// 		nextCurNode = nil // reset nextCurNode until an entry point to next level is found

// 		// connect next level:
// 		var nonNilChildren []*Node
// 		for curNode != nil { // traverse current level and add children in next level
// 			if curNode.Left != nil {
// 				nonNilChildren = append(nonNilChildren, curNode.Left)
// 				if nextCurNode == nil { // update entry point to next level
// 					nextCurNode = curNode.Left
// 				}
// 			}
// 			if curNode.Right != nil {
// 				nonNilChildren = append(nonNilChildren, curNode.Right)
// 				if nextCurNode == nil {
// 					nextCurNode = curNode.Right
// 				}
// 			}
// 			curNode = curNode.Next
// 		}

// 		// connect all nodes
// 		for i := 0; i < len(nonNilChildren)-1; i++ {
// 			nonNilChildren[i].Next = nonNilChildren[i+1]
// 		}
// 		curNode = nextCurNode
// 	}

// 	return root

// }

// follow up: use constant space: Iterative solution (BFS). Assuming a level is fully connected, we traverse the level with Next pointer(as if it is a linked list), at the same time we also connect all non-nil children in the next level: Use a tail pointer pointing to the last non-nil child that has been connected, and when a new non-nil child is found by a node in the current level, connect the tail and move tail to the node. The entry point to the next level is the first non-null child pointer of the current level.

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	// initialize
	curNode := root

	// nextCurNode: first non-nul child of root
	var nextCurNode *Node = nil
	if curNode.Left != nil {
		nextCurNode = curNode.Left
	} else if curNode.Right != nil {
		nextCurNode = curNode.Right
	}

	for nextCurNode != nil {
		nextCurNode = nil // reset nextCurNode until an entry point to next level is found
		var tail *Node    // edit nodes in next level
		// connect next level:
		for curNode != nil { // traverse current level and add children in next level
			if curNode.Left != nil {
				if nextCurNode == nil { // update entry point to next level
					nextCurNode = curNode.Left
					tail = curNode.Left
				}
				tail.Next = curNode.Left
				tail = curNode.Left
			}
			if curNode.Right != nil {
				if nextCurNode == nil {
					nextCurNode = curNode.Right
					tail = curNode.Right
				}
				tail.Next = curNode.Right
				tail = curNode.Right
			}
			curNode = curNode.Next
		}
		if tail != nil {
			tail.Next = nil
		}
		curNode = nextCurNode
	}

	return root

}

// [1,2,3,4,5,null,7,8,9,null,null,null,null,14,15]
// @lc code=end

