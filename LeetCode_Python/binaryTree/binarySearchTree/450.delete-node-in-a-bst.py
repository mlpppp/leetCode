#
# @lc app=leetcode id=450 lang=python3
#
# [450] Delete Node in a BST
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def findArgmin(self, root): 
        while (root.left != None):
             root = root.left
        return root
    

    def deleteNode(self, root: Optional[TreeNode], key: int) -> Optional[TreeNode]:
        if (root is None): return root
        if (root.val > key):
            treeLeftDeleted = self.deleteNode(root.left, key)
            root.left = treeLeftDeleted
        if (root.val < key):
            treeRightDeleted = self.deleteNode(root.right, key)
            root.right = treeRightDeleted
        if (root.val == key):
            if (not root.left):
                return root.right 
            if (not root.right):
                return root.left 
            successor = self.findArgmin(root.right)
            treeRightDeleted = self.deleteNode(root.right, successor.val)
            root.right = treeRightDeleted

            successor.left = root.left
            successor.right = root.right
            root = successor
            
            

        return root

        
# @lc code=end

