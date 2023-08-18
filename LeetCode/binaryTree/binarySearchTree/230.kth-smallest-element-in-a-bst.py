#
# @lc app=leetcode id=230 lang=python3
#
# [230] Kth Smallest Element in a BST
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    rank = 0 
    k = -1
    result = None
    def inorderWalk(self, root: Optional[TreeNode]):
        if (self.result): return
        if root is None:
            return 
        self.inorderWalk(root.left)
        self.rank+=1
        if (self.rank == self.k):
            self.result = root.val
            return 
        self.inorderWalk(root.right)
    def kthSmallest(self, root: Optional[TreeNode], k: int) -> int:
        self.rank = 0
        self.k = k
        self.result = None
        self.inorderWalk(root)
        return self.result 
        
# @lc code=end
# 只能用中序遍历， 按照排序的顺序过整个list, 外部变量rank记录已经走过的数量, 如果已经走过的数量正好等于K当前的node就是我们想要的结果