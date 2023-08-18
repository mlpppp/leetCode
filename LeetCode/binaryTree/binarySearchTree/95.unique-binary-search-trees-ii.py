#
# @lc app=leetcode id=95 lang=python3
#
# [95] Unique Binary Search Trees II
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    uniqueDict = {(0, 0):[None]}
    def generateTreesRecur(self, low: int, high: int) -> List[Optional[TreeNode]]:
        if (low, high) in self.uniqueDict: 
            return self.uniqueDict[(low, high)]
        else: 
            for rootVal in range(low, high+1):
                leftTreeRange = (low, rootVal-1)
                rightTreeRange = (rootVal+1, high)
                if rootVal == low:
                    leftTreeRange = (0, 0)
                if rootVal == high:
                    rightTreeRange = (0, 0)

                leftTreesList = self.generateTreesRecur(leftTreeRange[0], leftTreeRange[1])
                rightTreesList = self.generateTreesRecur(rightTreeRange[0], rightTreeRange[1])


                for treeL in leftTreesList:
                    for treeR in rightTreesList:
                        root = TreeNode(rootVal)
                        root.left = treeL
                        root.right = treeR
                        if ((low, high) not in self.uniqueDict):
                            self.uniqueDict[(low, high)] = [root]
                        else:
                            self.uniqueDict[(low, high)].append(root)

            return self.uniqueDict[(low, high)]


    def generateTrees(self, n: int) -> List[Optional[TreeNode]]:
        
        return self.generateTreesRecur(1, n)



        
# @lc code=end

# 如果要优化的话可以用二叉树sequence(codec)算法把每一个结构unique的树编码