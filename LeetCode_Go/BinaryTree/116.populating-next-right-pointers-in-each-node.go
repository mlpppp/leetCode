/*
 * @lc app=leetcode id=116 lang=golang
 *
 * [116] Populating Next Right Pointers in Each Node
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

// Given a `perfect binary tree` where all leaves are on the same level, and every parent has two children. Populate a new field `Next` in every node pointing to the next node in the same level (or Null if there is no next node).

// Make use of the Next pointer. Each iteration we populate all Next pointers in the next level, so in the next iteration we traverse the next level with Next pointer, just like travsersing a linked list.

// O(n) time and O(1) Space.

// Solution 2: BFS. Solution 3: traverse with two nodes group [link](https://labuladong.github.io/algo/di-yi-zhan-da78c/shou-ba-sh-66994/dong-ge-da-cbce8/#%E7%AC%AC%E4%BA%8C%E9%A2%98%E3%80%81%E5%A1%AB%E5%85%85%E8%8A%82%E7%82%B9%E7%9A%84%E5%8F%B3%E4%BE%A7%E6%8C%87%E9%92%88)

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	var cur *Node = root
	var nextCur *Node = root.Left

	for nextCur != nil {
		for cur != nil {
			// fmt.Printf("cur: %d, nextCur: %d\n", cur.Val, nextCur.Val)
			cur.Left.Next = cur.Right
			if cur.Next != nil {
				cur.Right.Next = cur.Next.Left
			}
			cur = cur.Next
		}
		cur = nextCur
		nextCur = cur.Left

	}
	return root

}

// @lc code=end

