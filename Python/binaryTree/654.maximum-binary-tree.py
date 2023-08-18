# @before-stub-for-debug-begin
from python3problem654 import *
from typing import *
# @before-stub-for-debug-end

#
# @lc app=leetcode id=654 lang=python3
#
# [654] Maximum Binary Tree
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def constructMaximumBinaryTree(self, nums: List[int]) -> Optional[TreeNode]:
        if (not nums): return None

        # rootIdx = nums.index(max(nums)) # find argmax
        rootIdx = -1
        maxVal = -1
        for index, num in enumerate(nums):
            if num > maxVal:
                rootIdx = index
                maxVal = num

        leftTree = nums[0:rootIdx]
        rightTree = nums[rootIdx+1:]

        leftSubTree = self.constructMaximumBinaryTree(leftTree)
        rightSubTree = self.constructMaximumBinaryTree(rightTree)

        root = TreeNode(val=nums[rootIdx], left=leftSubTree, right=rightSubTree)

        return root   
# @lc code=end

