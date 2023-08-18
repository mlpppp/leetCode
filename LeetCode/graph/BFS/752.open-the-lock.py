#
# @lc app=leetcode id=752 lang=python3
#
# [752] Open the Lock
#

# @lc code=start
class Solution:
    addMap = {}
    minusMap = {}
    stateChanges = []

    def initSolver(self):
        self.addMap = {str(i):str(i+1) for i in range(9)} 
        self.addMap["9"] = "0"
        self.minusMap = {str(i):str(i-1) for i in range(1, 10)} 
        self.minusMap["0"] = "9"
        self.stateChanges = [[direction,i] for direction in ['minus', 'add'] for i in range(4)]

    
    def flipOneDigit(self, code, stateChange):
        flipedCode = self.strToCode(code)
        if stateChange[0] == "add":
            flipedCode[stateChange[1]] = self.addMap[code[stateChange[1]]]
            return self.codeToStr(flipedCode)
        
        if stateChange[0] == "minus":
            flipedCode[stateChange[1]] = self.minusMap[code[stateChange[1]]]
            return self.codeToStr(flipedCode)


    def codeToStr(self, code):
        return ''.join(code)

    def strToCode(self, codeStr):
        return [*codeStr]

    def BFS_solver (self, initCode, deadends, target):
        if initCode in deadends:
            return -1
        visitedOrToVisit = {initCode}  # set of string-code
        toVisit = [initCode] # queue of string-code
        step = 0
        while len(toVisit) != 0:
            xStepSize = len(toVisit)
            for idx in range(xStepSize):
                curCode = toVisit.pop(0) #'0000'
                if curCode == target:
                    return step
                for stateChange in self.stateChanges:
                    newCode = self.flipOneDigit(curCode, stateChange)
                    if newCode not in deadends and newCode not in visitedOrToVisit:
                        toVisit.append(newCode)
                        visitedOrToVisit.add(newCode)
            step += 1


    def openLock(self, deadends: List[str], target: str) -> int:
        self.initSolver()
        initCode = '0000'
        step = self.BFS_solver(initCode, set(deadends), target)
        return step if step is not None else -1
        
# @lc code=end

