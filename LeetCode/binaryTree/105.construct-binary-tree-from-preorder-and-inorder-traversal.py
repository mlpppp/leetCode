#
# @lc app=leetcode id=105 lang=python3
#
# [105] Construct Binary Tree from Preorder and Inorder Traversal
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:

    def buildTree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        if len(preorder) == 0: return 

        rootVal = preorder[0]
        for idx,val in enumerate(inorder):
            if val == rootVal:
                inorderIdx = idx

        # 划分inorder
        inorder_L = inorder[0:inorderIdx] 
        inorder_R = inorder[inorderIdx+1:]  

        # 划分preorder    
        preorder_L = preorder[1:1+len(inorder_L)]   
        preorder_R = preorder[1+len(inorder_L):]       

        leftRoot= self.buildTree(preorder_L, inorder_L)
        rightRoot= self.buildTree(preorder_R, inorder_R)

        root = TreeNode(val=rootVal, left=leftRoot, right=rightRoot)

        return root
        
# @lc code=end

# 想办法通过rootVal把preorder和inorder各自划分成    
    #inorder_L; inorder_R
        # inorder:= L + rootVal + R
        # rootVal左边的子数组就是L右边的子数组就是R
    #preorder_L; preorder_R
        # preorder:= rootVal + L + R
        # 注意到inorder_L和preorder_L的长度是相同的， ez

# 建立L, R 2个子树
    # leftRoot = buildTree(preorder_L, inorder_L)
    # rightRoot = buildTree(preorder_R, inorder_R)

