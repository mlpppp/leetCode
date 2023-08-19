#
# @lc app=leetcode id=889 lang=python3
#
# [889] Construct Binary Tree from Preorder and Postorder Traversal
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def constructFromPrePost(self, preorder: List[int], postorder: List[int]) -> Optional[TreeNode]:
        if len(preorder)==0:
            return 
        if len(preorder)==1:
            root = TreeNode(val=preorder[0], left=None, right=None)
            return root
        if len(preorder) != len(postorder):
            raise ValueError('len should be identical')

        rootVal = preorder[0]
        lTreeVal = preorder[1]
        
        for idx, val in enumerate(postorder):
            if val == lTreeVal:
                lTreePostIdx = idx 
                
        # construct postorder_L
        postorder_L = postorder[0:lTreePostIdx+1]
        # construct preorder_L
        preorder_L = preorder[1:1+len(postorder_L)]
        # construct postorder_R
        postorder_R = postorder[lTreePostIdx+1:-1]
        # construct preorder_R
        preorder_R = preorder[1+len(postorder_L):]
        

        # construct preorder_L

        leftRoot = self.constructFromPrePost(preorder_L, postorder_L)
        rightRoot = self.constructFromPrePost(preorder_R, postorder_R)

        root = TreeNode(val=rootVal, left=leftRoot, right=rightRoot)
        return root

        
        
# @lc code=end

