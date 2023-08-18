#
# @lc app=leetcode id=98 lang=python3
#
# [98] Validate Binary Search Tree
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
from logging.config import valid_ident


class Solution:
    def findMin (self, root: Optional[TreeNode]):
        # logn
        originalRoot = root
        while (root.left is not None):
            root = root.left
            if hasattr(root, 'myMin'):
                return root.myMin
        originalRoot.myMin = root.val
        return root.val
    def findMax (self, root: Optional[TreeNode]):
        # logn
        originalRoot = root
        while (root.right is not None): 
            root = root.right
            if hasattr(root, 'myMax'):
                return root.myMax
        originalRoot.myMax = root.val
        return root.val


    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        if root is None: return True
        if root.left is None and root.right is None:    # leaf
            return True
        isValidLeft = self.isValidBST(root.left)
        isValidRight = self.isValidBST(root.right)
        if (isValidLeft and isValidRight):
            if (root.left is None):
                return root.val < self.findMin(root.right)
            if (root.right is None):
                return root.val > self.findMax(root.left)
            else:
                return self.findMax(root.left) < root.val < self.findMin(root.right)
        else:
            return False

            
        
# @lc code=end

# vaild条件:
#     1: left is vaild
#     2. right is valid
#     3. max(left) < root < min (right)

# 更简洁的代码 : 有参数递归: 
# TODO 不太懂

# def isValidBST(root: Optional[TreeNode]) {
#     return self.isValidBSTRecur(root, null, null)
# }

# # /* 限定以 root 为根的子树节点必须满足 max.val > root.val > min.val */
# def isValidBSTRecur(self, root: Optional[TreeNode], min: Optional[TreeNode], max: Optional[TreeNode]) 
#     if (root is None): return True
#     # vialation case
#     if (min is not None and root.val <= min.val): 
#         return False
#     if (max is not None and root.val >= max.val): 
#         return False
#     #  限定左子树的最大值是 root.val，右子树的最小值是 root.val
#     return self.isValidBSTRecur(root.left, min, root) \
#             and self.isValidBSTRecurself(root.right, root, max)
