#
# @lc app=leetcode id=47 lang=python3
#
# [47] Permutations II
#

# @lc code=start
class Solution:
    result = []
    nums =[]
    used = []
    def backtrack(self, partialSol):
        if len(partialSol) == len(self.nums):
            self.result.append(partialSol.copy())

        for (idx, num) in enumerate(self.nums):
            if self.used[idx]:
                continue
            if (idx > 0 and self.nums[idx] == self.nums[idx-1] and (not self.used[idx-1])): # 这里是关键 (not self.used[idx-1])
                continue
            self.used[idx] = True
            partialSol.append(num)
            self.backtrack(partialSol)
            self.used[idx] = False
            partialSol.pop(-1)
        
    def permuteUnique(self, nums: List[int]) -> List[List[int]]:
        self.result = []
        self.nums = nums
        self.nums.sort()
        self.used = [False for i in range(len(nums))]
        self.backtrack([])
        return self.result
        
# @lc code=end

# https://labuladong.github.io/algo/1/9/#排列元素可重不可复选

# 另一剪法 
class Solution:
    result = []
    nums =[]
    used = []
    def backtrack(self, partialSol):
        if len(partialSol) == len(self.nums):
            self.result.append(partialSol.copy())

        prevNum = -100  # 记录同一层左边sibling元素的值
        for (idx, num) in enumerate(self.nums):
            if self.used[idx]:
                continue
            if (self.nums[idx] == prevNum): 
                continue
            self.used[idx] = True
            partialSol.append(num)
            prevNum = num
            self.backtrack(partialSol)
            self.used[idx] = False
            partialSol.pop(-1)
        
    def permuteUnique(self, nums: List[int]) -> List[List[int]]:
        self.result = []
        self.nums = nums
        self.nums.sort()
        self.used = [False for i in range(len(nums))]
        self.backtrack([])
        return self.result