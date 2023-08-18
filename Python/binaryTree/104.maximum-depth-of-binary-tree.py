#
# @lc app=leetcode id=104 lang=python3
#
# [104] Maximum Depth of Binary Tree
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        if root is None: return 0
        if ((root.left is None) and (root.right is None)): return 1
        childDepth = max(self.maxDepth(root.left), self.maxDepth(root.right))
        return childDepth+1


# alt: 自顶向下
# // 记录最大深度
# int res = 0;
# # // 记录遍历到的节点的深度
# int depth = 0;

# # // 主函数
# int maxDepth(TreeNode root) {
# 	traverse(root);
# 	return res;
# # }

# # // 二叉树遍历框架
# void traverse(TreeNode root) {
# 	if (root == null) {
# 		return;
# 	}
# 	# // 前序位置
# 	depth++;
#     if (root.left == null && root.right == null) {
#         // 到达叶子节点，更新最大深度
# 		res = Math.max(res, depth);
#     }
# 	traverse(root.left);
# 	traverse(root.right);
# 	# // 后序位置
# 	depth--;
# }
        
# @lc code=end

