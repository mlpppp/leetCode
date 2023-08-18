#
# @lc app=leetcode id=222 lang=python3
#
# [222] Count Complete Tree Nodes
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def countNodes(self, root: Optional[TreeNode]) -> int:
        if root is None:
            return 0

        leftHeight = rightHeight = 0
        leftPointer = rightPointer = root

        while (leftPointer is not None):
            leftHeight+=1
            leftPointer = leftPointer.left
       
        while (rightPointer is not None):
            rightHeight+=1
            rightPointer = rightPointer.right

        if (leftHeight == rightHeight):
            return 2**leftHeight-1
        else:
            return 1 + self.countNodes(root.left) + self.countNodes(root.right)

        
# @lc code=end
# O(logn*logn)
# https://labuladong.github.io/algo/2/21/48/

# counting #nodes in a general binary tree
    # def countNodes(root: Optional[TreeNode]) -> int:
    #     if root is None:
    #         return 0
    #     return 1+countNodes(root.left)+countNodes(root.right)
    # O(n)

# count #nodes in a general perfect binary tree:
    # def countNodes(root: Optional[TreeNode]) -> int:
    #     height = 0
    #     while (root is not None):
    #         height+=1
    #         root = root.right
    #     return 2^height-1
    # O(logn)

# Complete binary tree is the combination of the two algorithm, using counting sub-perfect-binary-tree to optimize the complexity