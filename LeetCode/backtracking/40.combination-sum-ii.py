#
# @lc app=leetcode id=40 lang=python3
#
# [40] Combination Sum II
#

# @lc code=start
class Solution:
    result = []
    candidates = []
    target = 0
    def backtrack(self, subset, start, trackSum):
        if trackSum == self.target:
            self.result.append(subset.copy())
        for idx in range(start, len(self.candidates)):  
            if idx > start and self.candidates[idx] == self.candidates[idx-1]:
                continue
            trackSum += self.candidates[idx]
            if trackSum <= self.target:
                subset.append(self.candidates[idx])
                self.backtrack(subset, idx+1, trackSum)
                subset.pop(-1)
            trackSum -= self.candidates[idx]
        
    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        self.result = [] 
        self.candidates = candidates
        self.candidates.sort()
        self.target = target
        self.backtrack([], start = 0, trackSum = 0)
        return self.result



# @lc code=end

