#
# @lc app=leetcode id=283 lang=python3
#
# [283] Move Zeroes
#

# @lc code=start
class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        pEdit = 0     
        pFilter = 0   
        while(pFilter < len(nums)):
            if(nums[pFilter] != 0):
                nums[pEdit] = nums[pFilter]
                pEdit+=1
            pFilter+=1
        for i in range(pEdit, len(nums)):
            nums[i] = 0

# pEdit应该永远指向第一个零的位置 
# pFilter寻找非0元素当找到一个非0元素它就与pEdit交换， 并且把pEdit指向下一个0
# @lc code=end

# 010312
# 100312
# 130012
# 131002
# 131200