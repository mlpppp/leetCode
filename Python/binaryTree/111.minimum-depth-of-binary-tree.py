#
# @lc app=leetcode id=111 lang=python3
#
# [111] Minimum Depth of Binary Tree
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def minDepth(self, root: Optional[TreeNode]) -> int:
        if root is None:
            return 0
        queue = [root]
        layer = 1
        while(len(queue) > 0):
            layerSize = len(queue)
            for nodeIdx in range(layerSize):
                curNode = queue.pop(0)
                if curNode.left is None and curNode.right is None:
                    return layer
                if curNode.left is not None:
                    queue.append(curNode.left)               
                if curNode.right is not None:
                    queue.append(curNode.right)
            layer+=1

        
# @lc code=end

