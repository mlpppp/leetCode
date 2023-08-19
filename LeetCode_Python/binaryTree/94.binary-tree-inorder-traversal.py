#
# @lc app=leetcode id=94 lang=python3
#
# [94] Binary Tree Inorder Traversal
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def traverse(self, root, result):
        if (root is None): return
        self.traverse(root.left, result)
        result.append(root.val)
        self.traverse(root.right, result)

    def inorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        result = []
        self.traverse(root, result)
        return result 
# @lc code=end

