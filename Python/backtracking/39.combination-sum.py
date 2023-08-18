#
# @lc app=leetcode id=39 lang=python3
#
# [39] Combination Sum
#

# @lc code=start
class Solution:
    result = []
    target = 0
    candidates = []
    def backtrack(self, trackSum, subset, start):
        if trackSum == self.target:
            self.result.append(subset.copy())
        for idx in range(start, len(self.candidates)):
            trackSum += self.candidates[idx]
            if trackSum <= self.target:
                subset.append(self.candidates[idx])
                self.backtrack(trackSum, subset, idx)
                subset.pop(-1)
            trackSum -= self.candidates[idx]

        pass
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        self.result = []
        self.candidates = candidates
        self.candidates.sort()
        self.target = target
        self.backtrack(trackSum = 0, subset=[], start = 0)
        return self.result
        
# @lc code=end

