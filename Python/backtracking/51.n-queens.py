#
# @lc app=leetcode id=51 lang=python3
#
# [51] N-Queens
#

# @lc code=start
class Solution:
    result = []
    n = -1
    def diagRowIntersectionLeft(self, diagIdx, rowIdx):
        # return colIdx where diag of diagIdx and row rowIdx intersect
        val = (self.n-1)+rowIdx-diagIdx
        if 0 <= val and val <= (self.n-1):
            return val
        else:
            return None


    def diagRowIntersectionRight(self, diagIdx, rowIdx):
        # return colIdx where diag of diagIdx and row rowIdx intersect
        val = diagIdx-rowIdx
        if 0 <= val and val <= self.n-1:
            return val
        else:
            return None
        
    
    def myLeftUpdiagIdx(self, x, y):
        return (self.n-1)+x-y
        
    def myRightUpdiagIdx(self, x, y):
        return x+y
        
    def diagRowIntersection_batch(self, diagIndices, rowIdx, dir):
        # given a list of diagIndices, return set of colIdx: colList where diagIdx intersect rowIdx
        # dir: left-up; right-up

        colList = []
        if len(diagIndices) == 0:
            return colList
        if dir == 'left-up':
            for diagIdx in diagIndices.keys():
                col = self.diagRowIntersectionLeft(diagIdx, rowIdx)
                if col is not None:
                    colList.append(col)
        elif dir == 'right-up':
            for diagIdx in diagIndices.keys():
                col = self.diagRowIntersectionRight(diagIdx, rowIdx)
                if col is not None:
                    colList.append(col)

        return colList

    def formatBoardToResult(self, board):
        for rowIdx in range(len(board)):
            board[rowIdx] = ''.join(board[rowIdx])
        return board # board formatted as result


    def backtrack(self, board, solutionRow, usedCol, usedLeftDiag, usedRightDiag):
        if (solutionRow == len(board)):
            self.result.append(self.formatBoardToResult(board.copy()))
            return 
            

        # find chooseSet
        chooseSet = [idx for idx in range(self.n)]
   
        diagLeftUnavaiable = self.diagRowIntersection_batch(usedLeftDiag, solutionRow, 'left-up')

        diagRightUnavaiable = self.diagRowIntersection_batch(usedRightDiag, solutionRow, 'right-up')  

        chooseSet = set(chooseSet).difference(set(usedCol.keys())).difference(set(diagLeftUnavaiable)).difference(set(diagRightUnavaiable))

        for colIdx in chooseSet:
            #! draw the queen
            # known: colIdx, leftUpdiagIdx, rightUpdiagIdx
            leftUpdiagIdx = self.myLeftUpdiagIdx(solutionRow, colIdx)
            rightUpdiagIdx = self.myRightUpdiagIdx(solutionRow, colIdx)


            # change 'Q' in current board 
            board[solutionRow][colIdx] = "Q"
            # update: usedLeftDiag, usedRightDiag,usedCol
            usedLeftDiag[leftUpdiagIdx] = True
            usedRightDiag[rightUpdiagIdx] = True
            usedCol[colIdx] = True

            # ! recursion
            self.backtrack(board, solutionRow+1, usedCol, usedLeftDiag, usedRightDiag)
            # ! UNDOs
            # undo change 'Q' in board
            board[solutionRow][colIdx] = "."
            # undo: usedCol, usedLeftDiag, usedRightDiag
            
            usedLeftDiag.pop(leftUpdiagIdx)
            usedRightDiag.pop(rightUpdiagIdx)
            usedCol.pop(colIdx)

    def solveNQueens(self, n: int) -> List[List[str]]:
        self.result = []
        self.n = n
        board = [['.' for i in range(n)] for i in range(n)]
        self.backtrack(board, solutionRow = 0, usedCol={}, usedLeftDiag={}, usedRightDiag={})
        return self.result
        
# @lc code=end

