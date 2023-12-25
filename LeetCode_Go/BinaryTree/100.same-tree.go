/*
 * @lc app=leetcode id=100 lang=golang
 *
 * [100] Same Tree
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

// easy, either traverse the trees at the same time, or traverse each tree and return and encoded string, then compare the strings. Encoding: 1. inorder + preorder/postorder. 2. preorder + null node 3. postorder + null node 4. level order

func isSameTree(p *TreeNode, q *TreeNode) bool {
	pstr := preOrderEncode(p)
	qstr := preOrderEncode(q)
	fmt.Printf("pstr: %s, qstr: %s\n", pstr, qstr)
	return pstr == qstr
}

func preOrderEncode(root *TreeNode) string {
	if root == nil {
		return "#,"
	}
	rootStr := strconv.Itoa(root.Val) + ","
	leftString := preOrderEncode(root.Left)
	rightString := preOrderEncode(root.Right)
	return rootStr + leftString + rightString
}

// @lc code=end

