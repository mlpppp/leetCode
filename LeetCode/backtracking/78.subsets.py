#
# @lc app=leetcode id=78 lang=python3
#
# [78] Subsets
#

# @lc code=start
class Solution:
    result = []
    def backtrack(self, nums, partialSol, start):
        self.result.append(partialSol.copy())

        for startIdx in range(start, len(nums)):
            partialSol.append(nums[startIdx])
            self.backtrack(nums, partialSol, startIdx+1)
            partialSol.pop(-1)
        
    def subsets(self, nums: List[int]) -> List[List[int]]:
        self.result = []
        self.backtrack(nums, [], 0)
        return self.result

        
# @lc code=end

