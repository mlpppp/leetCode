#
# @lc app=leetcode id=27 lang=python3
#
# [27] Remove Element
#

# @lc code=start
class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        pEdit = 0
        pFilter = 0
        while (pFilter < len(nums)):
            if (nums[pFilter] == val):
                pFilter+=1
                continue
            if(pEdit != pFilter):
                nums[pEdit] = nums[pFilter]      
            pFilter+=1
            pEdit+=1

        return pEdit 



    # pFilter的作用类似一个filter一般来说他把所有的input原样发给pEdit但是如果碰到了val他就把这个input给filter掉(Return Nothing )  
# @lc code=end

