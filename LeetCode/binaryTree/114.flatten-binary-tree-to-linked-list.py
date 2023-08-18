#
# @lc app=leetcode id=114 lang=python3
#
# [114] Flatten Binary Tree to Linked List
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:

    def flatten(self, root: Optional[TreeNode]) -> None:
        """
        Do not return anything, modify root in-place instead.
        """
        if (root is None): return 

        rightLinkList = self.flatten(root.right) 
        leftLinkList = self.flatten(root.left)  

        if (leftLinkList):
            # 连接root和左链表
            root.right  = leftLinkList 
            leftLinkList.left = None
            # 连接左右链表
            while(leftLinkList.right):
                leftLinkList = leftLinkList.right
            leftLinkList.right = rightLinkList
        else: 
            root.right = rightLinkList    
        root.left = None 
        return root
        
# @lc code=end


    # 如果左边的链表存在把root连接到左边链表的表头
    # 接着遍历到左边链表的结尾，把左边链表的结尾连接到右边链表的表头