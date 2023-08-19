#
# @lc app=leetcode id=652 lang=python3
#
# [652] Find Duplicate Subtrees
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    uniqueStructure = {}
    dupStrucutre = []

    def serialization(self, root: Optional[TreeNode]) -> string:
        if root is None: return '#,'
        leftSerial = self.serialization(root.left)
        rightSerial = self.serialization(root.right)

        # got unique structure ID
        mySerial = leftSerial + rightSerial + str(root.val) + ','

        # compare mySerial with global record
        if(mySerial in self.uniqueStructure): # only append to res in first reappear
            if self.uniqueStructure[mySerial] == 1:
                self.uniqueStructure[mySerial]+=1
                self.dupStrucutre.append(root)
        else: # first time: create record
            self.uniqueStructure[mySerial] = 1

        return mySerial

    def findDuplicateSubtrees(self, root: Optional[TreeNode]) -> List[Optional[TreeNode]]:
        # re-init global vars
        self.uniqueStructure = {}
        self.dupStrucutre = []
        self.serialization(root)
        print(self.uniqueStructure)

        return self.dupStrucutre




# @lc code=end

# 序列化遍历整个树得到以任意一个node为树根的子树的结构, 序列化 :=> 任何相同的结构应该都有相同的序列
# 全局变量uniqueStructure (hash map)存储出现过的，unique的结构（序列），
# 如果任意一个node的子树序列在uniqueStructure出现过一次，就在结果中加上她

# 后序遍历！！ 这样保证获得整个子树结构后，在树根进行判断

