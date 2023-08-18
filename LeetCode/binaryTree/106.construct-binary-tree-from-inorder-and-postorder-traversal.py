#
# @lc app=leetcode id=106 lang=python3
#
# [106] Construct Binary Tree from Inorder and Postorder Traversal
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def buildTree(self, inorder: List[int], postorder: List[int]) -> Optional[TreeNode]:
        if len(postorder) == 0: return 

        rootVal = postorder[-1]
        for idx,val in enumerate(inorder):
            if val == rootVal:
                inorderIdx = idx

        # 划分inorder
        inorder_L = inorder[0:inorderIdx] 
        inorder_R = inorder[inorderIdx+1:]  

        # 划分postorder   
        postorder_L = postorder[0:len(inorder_L)]   
        postorder_R = postorder[len(inorder_L):-1]       

        leftRoot = self.buildTree(inorder_L, postorder_L)
        rightRoot = self.buildTree(inorder_R, postorder_R)

        root = TreeNode(val=rootVal, left=leftRoot, right=rightRoot)

        return root
              
# @lc code=end

