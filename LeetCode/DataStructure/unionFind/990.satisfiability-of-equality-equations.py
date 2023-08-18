#
# @lc app=leetcode id=990 lang=python3
#
# [990] Satisfiability of Equality Equations
#

# @lc code=start
class unionFind:
    def __init__(self, ccNum) -> None:
        self.ccCount = ccNum 
        self.ccParents = [i for i in range(ccNum)] # each node point to it's parent
        self.ccSize = {i:1 for i in range(ccNum)}
    def find (self, x):
        # O(min(|smaller cc|) for first use after a union), O(1) for subsequent use
        # return ccParent of node x
        # compress the route by making all nodes in cc point to the same node
        if (self.ccParents[x] != x):
            self.ccParents[x] = self.find(self.ccParents[x])
        return  self.ccParents[x]

    def union(self, p, q):
        # O(1)
        # union connect component node p and q belong to 
        ccParentP = self.find(p)
        ccParentQ = self.find(q)
        if ccParentP == ccParentQ: # already same CC
            return 
        else:
            # 把size小的cc接到size大的cc
            # ! union加上之后find压缩的cost是 |cc of P|, 因为所有nodes in cc of P在压缩后变成ccParentQ
            smallerCC = ccParentP if self.ccSize[ccParentP] < self.ccSize[ccParentQ] else ccParentQ
            largerCC = ccParentQ if smallerCC is ccParentP else ccParentP
            
            # manage ccSize
            smallerSize = self.ccSize[smallerCC]
            self.ccSize[largerCC] += smallerSize
            self.ccSize.pop(smallerCC)

            self.ccParents[smallerCC] = largerCC
            self.ccCount -= 1

    def isConnect(self, p, q):
        # determine if p and q belongs to the same CC
        ccParentP = self.find(p)
        ccParentQ = self.find(q)
        return ccParentP == ccParentQ

    def connectComponentCount(self):
        # num of connect Component in graph
        return self.ccCount

    def connectComponentSize(self, p):
        # O(1)
        # the size of the connect Component node p belongs to 
        parent = self.find(p)
        return self.ccSize[parent]


import string
class Solution:
    def equationsPossible(self, equations: List[str]) -> bool:
        uf = unionFind(26)
        charList = list(string.ascii_lowercase)
        charToIdx = {char:id for [id, char] in enumerate(charList)}

        for equation in equations:
            if equation[1] == '=':
                uf.union(charToIdx[equation[0]], charToIdx[equation[3]])

        for equation in equations:
            if equation[1] == '!':
                if uf.isConnect(charToIdx[equation[0]], charToIdx[equation[3]]):
                    return False

        return True


        
# @lc code=end

