#
# @lc app=leetcode id=90 lang=python3
#
# [90] Subsets II
#

# @lc code=start
class Solution:
    result = []
    def backtrack(self, nums, subset, start):
        self.result.append(subset.copy())
        for idx in range(start, len(nums)):
            if idx > start and nums[idx] == nums[idx-1]:
                continue 
            subset.append(nums[idx])
            self.backtrack(nums, subset, idx+1)
            subset.pop(-1)
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        self.result = []
        nums.sort()
        self.backtrack(nums, [], 0)
        return self.result
        
# @lc code=end

