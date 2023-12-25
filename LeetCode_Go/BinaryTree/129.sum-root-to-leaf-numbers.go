/*
 * @lc app=leetcode id=129 lang=golang
 *
 * [129] Sum Root to Leaf Numbers
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

// Each root-to-leaf path in the tree represents a number, through append nodes value along the path as digits to the integer. Return the total sum of all root-to-leaf numbers. The answer will fit in a 32-bit integer.

// Tree traverse. Append nodes value as string in preorder place, and remove the nodes value from the string in postorder place. Add a complete pathValue at a leaf node to the sum.

// O(n) time, O(n) space.

// func sumNumbers(root *TreeNode) int {
// 	sum := 0
// 	sumNumbersHelper(root, &sum, "")
// 	return sum
// }

// func sumNumbersHelper(root *TreeNode, sum *int, curPathValue string) {
// 	if root == nil {
// 		return
// 	}
// 	curPathValue = curPathValue + strconv.Itoa(root.Val)
// 	if root.Left == nil && root.Right == nil { // add to result set if it is a leaf node
// 		intVal, err := strconv.Atoi(curPathValue)
// 		if err == nil {
// 			*sum += intVal
// 		} else {
// 			log.Println(err)
// 		}
// 	}
// 	sumNumbersHelper(root.Left, sum, curPathValue)
// 	sumNumbersHelper(root.Right, sum, curPathValue)
// 	curPathValue = curPathValue[:len(curPathValue)-1]
// 	return
// }

// followup: replaced append string with pure int operations.
func sumNumbers(root *TreeNode) int {
	sum := 0
	sumNumbersHelper(root, &sum, 0)
	return sum
}

func sumNumbersHelper(root *TreeNode, sum *int, curPathValue int) {
	if root == nil {
		return
	}
	curPathValue = curPathValue*10 + root.Val
	if root.Left == nil && root.Right == nil { // add to result set if it is a leaf node
		*sum += curPathValue
	}

	sumNumbersHelper(root.Left, sum, curPathValue)
	sumNumbersHelper(root.Right, sum, curPathValue)
	curPathValue = (curPathValue - root.Val) / 10
	return
}

// @lc code=end

