#
# @lc app=leetcode id=77 lang=python3
#
# [77] Combinations
#

# @lc code=start
class Solution:
    result = []
    k = 0
    n = 0
    def combine(self, n: int, k: int) -> List[List[int]]:
        self.result = []
        self.k = k
        self.n = n
        self.backtrack([], 1)
        return self.result

    def backtrack(self, partialSol, startNum):
        if len(partialSol) == self.k:
            self.result.append(partialSol.copy())
            return 
        for num in range(startNum, self.n+1):
            partialSol.append(num)
            self.backtrack(partialSol, num+1)
            partialSol.pop(-1)

        
        
# @lc code=end

