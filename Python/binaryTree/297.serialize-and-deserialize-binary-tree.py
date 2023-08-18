#
# @lc app=leetcode id=297 lang=python3
#
# [297] Serialize and Deserialize Binary Tree
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Codec:
    sequence = ''
    def serialize(self, root):
        """Encodes a tree to a single string.
        
        :type root: TreeNode
        :rtype: str
        """
        if(root is None): 
            self.sequence+='#,'       
            return self.sequence
        self.sequence+=str(root.val) +','

        self.serialize(root.left)
        self.serialize(root.right)
        return self.sequence
        

    def deserializeRecur(self, seqList):
        """Decodes your encoded data to tree.
        
        :type data: str
        :rtype: TreeNode
        """    
        if(not seqList): return None
        rootVal = seqList.pop(0)
        if rootVal == '#': return None 
        # this is the key: return None two time we get a leaf, return None one times we get a single leaf tree
        root = TreeNode(val=rootVal)
        root.left = self.deserializeRecur(seqList)
        root.right = self.deserializeRecur(seqList)

        return root

    def deserialize(self, data):
        print(data)
        seqList = [d for d in data.split(',') if d]
        return self.deserializeRecur(seqList)

# Your Codec object will be instantiated and called as such:
# ser = Codec()
# deser = Codec()
# ans = deser.deserialize(ser.serialize(root))
# @lc code=end

