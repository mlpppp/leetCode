/*
 * @lc app=leetcode id=199 lang=golang
 *
 * [199] Binary Tree Right Side View
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Given the root of a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.

//			  1
//			/   \
//		   2     3
//		  / \   / \
//		 4   5 6   7
//		/     \ \
//	   8       9 10
//		\		\
//		 11		12
//		  \
//		  13
//
// 1,3,7,10,12,13

// DFS: use a map to record the rightmost node of each level. For each node, visit right before visit left.

func rightSideView(root *TreeNode) []int {
	result := make(map[int]int)
	rightSideViewHelper(root, result, 1)
	arrryResult := make([]int, len(result))
	for i := 1; i <= len(result); i++ {
		arrryResult[i-1] = result[i]
	}
	return arrryResult
}

func rightSideViewHelper(root *TreeNode, result map[int]int, rootLevel int) {
	if root == nil {
		return
	}
	// record the first encounter node of each level
	if _, existed := result[rootLevel]; !existed {
		result[rootLevel] = root.Val
	}
	// fmt.Printf("curLevel: %d, cur.Val: %d, result: %v\n", rootLevel, root.Val, result)

	rightSideViewHelper(root.Right, result, rootLevel+1)
	rightSideViewHelper(root.Left, result, rootLevel+1)
}

// BFS: add the last node (node that can be viewed) in a level to the result array. (similar to BFS reversed/zigzag traverse)

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{root}
	result := []int{}
	for len(queue) > 0 {
		nextLevel := []*TreeNode{}
		// add the last node in a level to the result array.

		levelWidth := len(queue)
		for i := 0; i < levelWidth; i++ {
			// pop the queue
			curNode := queue[0]
			queue = queue[1:]
			if i == 0 {
				result = append(result, curNode.Val) // ! add to result if it is the first node
			}
			// add to queue in reversed order
			if curNode.Right != nil {
				nextLevel = append(nextLevel, curNode.Right)
			}
			if curNode.Left != nil {
				nextLevel = append(nextLevel, curNode.Left)
			}
		}
		queue = nextLevel
	}
	return result
}

// @lc code=end

