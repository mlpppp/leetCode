#
# @lc app=leetcode id=543 lang=python3
#
# [543] Diameter of Binary Tree
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    depthTable = {}  # use this to avoid repeating maxDepth() on  same nodes 
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        if root is None: return 0
        if ((root.left is None) and (root.right is None)): return 1
        childDepth = max(self.maxDepth(root.left), self.maxDepth(root.right))
        self.depthTable[root] = childDepth+1
        return childDepth+1

    def diameterOfBinaryTree(self, root: Optional[TreeNode]) -> int:
        if root is None or (root.left is None and root.right is None): return 0
        # dynamic progromming table
        if root.left in self.depthTable:
            depthL = self.depthTable[root.left]
        else: 
            depthL = self.maxDepth(root.left)
            self.depthTable[root.left] = depthL

        if root.right in self.depthTable:
            depthR = self.depthTable[root.right]
        else: 
            depthR = self.maxDepth(root.right)
            self.depthTable[root.right] = depthR


        diameter = max(self.diameterOfBinaryTree(root.left),
                       self.diameterOfBinaryTree(root.right),
                        depthR+depthL )
        return diameter

    # 每一个树都有一个diameterOfBinaryTree(Root)
    #     这个diameter要么经过树根R要么不经过树根R
    # 如果经过R, 有
    #     diameterdiainaryTree(Root) == maxDepth(L) + maxDepth(R) 
    # 如果不经过R, 有   
    #     diameterdiainaryTree(Root) == max(diameterdiainaryTree(L), diameterdiainaryTree(R))  
    # maxDepth() 函数在leetcode: 104

#! alt solution: optimization design, modify maxDepth directly
    # 解决这题的关键在于，每一条二叉树都只有「直径」长度(通过root)，就是一个节点的左右子树的最大深度之和。
    # 问题要求的是在整棵树中最长的一个这样的直径也就是优化一个全局值

class Solution:

    maxDia = 0
    def diameterOfBinaryTree(self, root: Optional[TreeNode]) -> int:
        self.calculateMaxDia(root)
        return self.maxDia


    def maxDepth(self, root: Optional[TreeNode]) -> int:
        if root is None: return 0
        if ((root.left is None) and (root.right is None)): return 1

        leftDepth = self.maxDepth(root.left)
        rightDepth = self.maxDepth(root.right)

        rootRoute = leftDepth + rightDepth
        if (rootRoute > self.maxDia): self.maxDia = rootRoute

        return max(leftDepth, rightDepth) + 1



# @lc code=end

