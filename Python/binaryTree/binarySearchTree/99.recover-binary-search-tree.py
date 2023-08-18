#
# @lc app=leetcode id=99 lang=python3
#
# [99] Recover Binary Search Tree
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    keyValue1 = None
    keyValue2 = None
    lastNode = None
    def recoverTree(self, root: Optional[TreeNode]) -> None:
        """
        Do not return anything, modify root in-place instead.
        """
        self.keyValue1 = None
        self.keyValue2 = None
        self.lastNode = None

        self.traverseInorderToFind(root)
        temp = self.keyValue1.val
        self.keyValue1.val = self.keyValue2.val
        self.keyValue2.val = temp  

    def traverseInorderToFind(self, root: Optional[TreeNode]):
        if root is None: return 
        self.traverseInorderToFind(root.left)
        
        if self.lastNode is not None: 
            if root.val < self.lastNode.val: 
                if self.keyValue1 is None:
                    self.keyValue1 = self.lastNode
                    self.keyValue2 = root # 2 value may be adjacent to eachother 
                else: 
                    self.keyValue2 = root  
                    # if we find (root.val < self.lastNode.val) the second time, then the swapped two must not be adjacent      

        self.lastNode = root

        self.traverseInorderToFind(root.right)
        
# @lc code=end

