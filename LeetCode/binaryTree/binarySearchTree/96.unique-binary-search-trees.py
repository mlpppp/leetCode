#
# @lc app=leetcode id=96 lang=python3
#
# [96] Unique Binary Search Trees
#

# @lc code=start
class Solution:
    uniqueDict = {0:1, 1:1, 2:2, 3:5}
    # store #unique BST in any continuous list of size idx (idx: #unique BST), given few init cases 
    def numTrees(self, n: int) -> int:
        if n in self.uniqueDict: 
            return self.uniqueDict[n]
        else: 
            result = 0
            for idx in range(n):
                numNodesLeft = idx
                numNodesright = n - (idx+1)
                result += self.numTrees(numNodesLeft)*self.numTrees(numNodesright)
            self.uniqueDict[n] = result
            return result

      
# @lc code=end

# https://appktavsiei5995.pc.xiaoe-tech.com/detail/i_62987940e4b01a4852072f8c/1?from=p_629871eee4b01a48520729f7&type=6&parent_pro_id=

