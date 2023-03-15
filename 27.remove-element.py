#
# @lc app=leetcode id=27 lang=python3
#
# [27] Remove Element
#

# @lc code=start
class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        pDelete = 0
        pSearch = 0
        while (pSearch < len(nums)):
            if(nums[pDelete] == val and nums[pSearch] != val):
                nums[pDelete] = nums[pSearch]
                pDelete += 1
            if (nums[pDelete] != val):
                pDelete += 1
            pSearch += 1
        if nums[pDelete] == val:
            return pDelete-1
        else:
            return pDelete

        
# @lc code=end

