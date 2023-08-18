#
# @lc app=leetcode id=101 lang=python3
#
# [101] Symmetric Tree
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
from math import ceil
class Solution:
    class NullNode:
        def __init__(self, val=None):
            self.val = val
        val = None

    def isSymmetric(self, root: Optional[TreeNode]) -> bool:
        levelQueue = [root]
        while(levelQueue):
            levelLen = len(levelQueue)
            # compare left and right
            leftTreeP = 0
            rightTreeP = len(levelQueue)-1
            while (leftTreeP < levelLen/2):
                # both not None
                if levelQueue[leftTreeP].val != levelQueue[rightTreeP].val:
                    return False
                leftTreeP+=1
                rightTreeP-=1
         
            # make the next level
            for nodeIdx in range(levelLen): 
                curNode = levelQueue.pop(0)
                if(curNode.val is None): continue
                if(curNode.left):
                    levelQueue.append(curNode.left)
                else:
                    levelQueue.append(self.NullNode)
                if(curNode.right):
                    levelQueue.append(curNode.right)
                else:
                    levelQueue.append(self.NullNode)

        return True
# @lc code=end

# 层次遍历ver
    # 层次遍历每一次循环在levelQueue中取得当前层次的所有节点
        # 用了辅助类NullNode来代表空的节点
    # 用2个指针分别从两端向中心比较levelQueue是否是对称，如果有任意pair值不相等则不是对称

# recursion ver
class Solution {
    public boolean isSymmetric(TreeNode root) {
        if (root == null) return true;
        # 检查两棵子树是否对称
        return check(root.left, root.right);
    }

    boolean check(TreeNode left, TreeNode right) {
        if (left == null || right == null) return left == right;
        # 两个根节点需要相同
        if (left.val != right.val) return false;
        # 左右子节点需要对称相同
        return check(left.right, right.left) && check(left.left, right.right);
    }
}