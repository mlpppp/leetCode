#
# @lc app=leetcode id=116 lang=python3
#
# [116] Populating Next Right Pointers in Each Node
#

# @lc code=start
"""
# Definition for a Node.
class Node:
    def __init__(self, val: int = 0, left: 'Node' = None, right: 'Node' = None, next: 'Node' = None):
        self.val = val
        self.left = left
        self.right = right
        self.next = next
"""

class Solution:
    
    def connect(self, root: 'Optional[Node]') -> 'Optional[Node]':
        if root is None: return None
        queue = [root]
        while (queue):
            levelLen = len(queue)
            for node in range(levelLen):

                parent = queue.pop(0)  
                if (node == levelLen-1): # right most node in the layer
                    parent.next = None
                elif (node < levelLen-1):
                    parent.next = queue[0]

                if (parent.left):
                    queue.append(parent.left) 
                if (parent.right):
                    queue.append(parent.right)

        return root

# 用队列实现层次遍历  (level order traversal)      
# 应该可以推广到非Perfect tree, ie. irregular levelLen
        
# @lc code=end

