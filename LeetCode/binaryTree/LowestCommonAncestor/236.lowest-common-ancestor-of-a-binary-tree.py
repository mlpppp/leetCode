# @before-stub-for-debug-begin
from platform import node
from unittest import result
from python3problem236 import *
from typing import *
# @before-stub-for-debug-end

#
# @lc app=leetcode id=236 lang=python3
#
# [236] Lowest Common Ancestor of a Binary Tree
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    # def lowestCommonAncestorRecur(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode'):


    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        if root is None: return None
        elif (root is p) or (root is q): 
            return root
        else:
            leftResult = self.lowestCommonAncestor(root.left, p, q)
            rightResult = self.lowestCommonAncestor(root.right, p, q)
            if (leftResult and rightResult):
                return root
            if ((not leftResult) and (not rightResult)):
                return None
            return leftResult if leftResult else rightResult 




# @lc code=end

# https://appktavsiei5995.pc.xiaoe-tech.com/detail/i_62987959e4b01a4852072fa5/1?from=p_629871eee4b01a48520729f7&type=6&parent_pro_id=




# similar to 函数findEitherOne的变体 （
    # findEitherOne(root, p, q) : Find the node in tree who is p or q, A reference to P&Q is given, return either or the found node. you can return either of them

    def findEitherOne(root, p, q):
        if root is None: return None

        # return me if i am
        if root is p or root is q:
            return root

        # otherwise query children
        leftResult = findEitherOne(root.left, p, q)
        rightResult = findEitherOne(root.right, p, q)

        #  process the result
        if (not leftResult and not rightResult): return None # none the them
        if (leftResult and rightResult): return leftResult  # all of them, can also return rightResult
        return leftResult if leftResult else rightResult # either of them, 其实不必写前两行（这一行cover所以case，为了逻辑的清晰全写了


# findLCA logic:
# 对任意node:
def findLCA(root, p, q):
    if root is None: return None
    if root is either p or q:
        return root  # as either LCA or LCAchild
    if root is neither p or q:
            # query l, r trees
            leftResult = findLCA(root.left, p, q)
            rightResult = findLCA(root.right, p, q)
            results = (leftResult, rightResult)
            if (results have both p&q):
                return root as LCA
            if (either of results is q or q):
                return the result
            if (either of results is LCA):
                return the result
            if (results are both None):
                return None





###############################################################
# mey redundant sol...
# 31/31 cases passed (162 ms)
# Your runtime beats 16.11 % of python3 submissions
# Your memory usage beats 6.72 % of python3 submissions (28.7 MB)
class Solution:
    sol = None
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        self.sol = None
        self.lowestCommonAncestorRecur(root, p, q, False, False)
        return self.sol

    def lowestCommonAncestorRecur(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode', ancestorIsP, ancestorIsQ):
        if root is None: return (False, False)
        if self.sol: # skip recur
            return (True, True)

        if root is p:
            amAncestorP = True  
            if(ancestorIsQ):
                return (True, False)
        else:
            amAncestorP = False

        if root is q: 
            amAncestorQ = True 
            if(ancestorIsP):
                return (False, True) 
        else:
            amAncestorQ = False

        ######## query Left
        (leftIsAncestorP, leftIsAncestorQ) = self.lowestCommonAncestorRecur(root.left, p, q, amAncestorP, amAncestorQ)

        if (amAncestorP and leftIsAncestorQ):
            # found as root and child
            self.sol = root
            return (True, True)  # skip recur
        elif (amAncestorQ and leftIsAncestorP):
            # found as root and child
            self.sol = root
            return (True, True)  # skip recur
        
        if (leftIsAncestorP and leftIsAncestorQ):
            return (True, True)
        ######## query Right
        (rightIsAncestorP, rightIsAncestorQ) = self.lowestCommonAncestorRecur(root.right, p, q, amAncestorP, amAncestorQ)
        if (amAncestorP and rightIsAncestorQ):
            # found as root and child
            self.sol = root
            return (True, True)  # skip recur
         
        elif (amAncestorQ and rightIsAncestorP):
            # found as root and child
            self.sol = root
            return (True, True) # skip recur
        
        if (rightIsAncestorP and rightIsAncestorQ):
            return (True, True)

        ####### combine Left and Right
        if ((leftIsAncestorP and rightIsAncestorQ)  \
            or (leftIsAncestorQ and rightIsAncestorP)):
            # found as child and child
            self.sol = root
            return (True, True) 
            
        if amAncestorP or rightIsAncestorP or leftIsAncestorP:
            return (True, False)  # firstArg: contain P, secondArg: contain Q
        if amAncestorQ or rightIsAncestorQ or leftIsAncestorQ:
            return (False, True)
        else:
            return (False, False)

