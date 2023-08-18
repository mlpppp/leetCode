#
# @lc app=leetcode id=46 lang=python3
#
# [46] Permutations
#

# @lc code=start
class Solution:
    result = []

    def backtrack(self, nums, partialSolve):
        if len(partialSolve) == len(nums):
            self.result.append(partialSolve.copy())
            return 
        chooseSet = set(nums).difference(set(partialSolve))
        for num in chooseSet:
            partialSolve.append(num)
            print(partialSolve)
            self.backtrack(nums, partialSolve)
            partialSolve.remove(num)

    def permute(self, nums: List[int]) -> List[List[int]]:
        self.result = []
        self.backtrack(nums, [])
        return self.result
        
# @lc code=end

