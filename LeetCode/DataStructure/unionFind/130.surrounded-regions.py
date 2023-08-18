#
# @lc app=leetcode id=130 lang=python3
#
# [130] Surrounded Regions
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



class Solution:
    row = -1
    col = -1
    safeZone = -1

    def map2to1(self, x, y):
        return x*self.col + y

    def solve(self, board: List[List[str]]) -> None:
        """
        Do not return anything, modify board in-place instead.
        """
        if len(board) == 0:
            return 
        self.row = len(board)
        self.col = len(board[0])
        uf = unionFind(self.row*self.col+1)
        self.safeZone = self.row*self.col #last node is a dummy 'safeZone'

        # union 'o' in four surrounding edges
        for colIdx in range(self.col):
            if board[0][colIdx] == 'O':
                uf.union(self.safeZone, self.map2to1(0, colIdx))
            if board[self.row-1][colIdx] == 'O':
                uf.union(self.safeZone, self.map2to1(self.row-1, colIdx))

        for rowIdx in range(self.row):
            if board[rowIdx][0] == 'O':
                uf.union(self.safeZone, self.map2to1(rowIdx, 0))
            if board[rowIdx][self.col-1] == 'O':
                uf.union(self.safeZone, self.map2to1(rowIdx, self.col-1))

        # ! union middle part
        # direction matrix
        dirMat = [(1,0), (0,-1), (-1,0), (0,1)]
        for x in range(1, self.row-1):
            for y in range(1, self.col-1):
                if board[x][y] == 'O':
                    for dir in dirMat:
                        if board[x+dir[0]][y+dir[1]] == 'O':
                            uf.union(self.map2to1(x, y), self.map2to1(x+dir[0], y+dir[1]))

        # filp 'O's not connected to safeZone
        for x in range(1, self.row-1):
            for y in range(1, self.col-1):
                if board[x][y] == 'O':
                    if not uf.isConnect(self.safeZone, self.map2to1(x, y)):
                        board[x][y] = 'X'
        
        
# @lc code=end

