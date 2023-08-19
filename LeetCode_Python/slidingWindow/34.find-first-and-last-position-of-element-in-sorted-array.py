#
# @lc app=leetcode id=34 lang=python3
#
# [34] Find First and Last Position of Element in Sorted Array
#

# @lc code=start
from math import floor
class Solution:
    def findStart(self, nums, target):
        left = 0
        right = len(nums)-1
        while(left <= right):
            mid = floor(left + (right-left)/2)
            if (nums[mid] == target):
                right = mid-1
            if (nums[mid] > target):
                right = mid-1
            if (nums[mid] < target):
                left = mid+1
        if (left > len(nums)-1): return -1
        if(nums[left] == target):
            return left
        else: return -1      

    def findEnd(self, nums, target):
        left = 0
        right = len(nums)-1
        while(left <= right):
            mid = floor(left + (right-left)/2)
            if (nums[mid] == target):
                left = mid+1
            if (nums[mid] < target):
                left = mid+1
            if (nums[mid] > target):
                right = mid-1
        if (right < 0): return -1
        if(nums[right] == target):
            return right
        else: return -1     

    def searchRange(self, nums: List[int], target: int) -> List[int]:
        return [self.findStart(nums, target), self.findEnd(nums, target)]
        
        
# @lc code=end

# 两边收紧ver
    # 因为我们初始化 right = nums.length - 1
    # 所以决定了我们的「搜索区间」是 [left, right]
    # 所以决定了 while (left <= right), terminate: left == right+1, right == left-1
    # 同时也决定了 left = mid+1 和 right = mid-1

    # findStart() 时返回left (if left is target)
        # 注意向右越界 ( left > nums.length -1 )
    # findEnd() right (if right is target)
        # 注意向左越界 (right < 0 )

# 左边收紧ver
    # 因为我们初始化 right = nums.length
    # 所以决定了我们的「搜索区间」是 [left, right)
    # 所以决定了 while (left < right), terminate: left = right
    # 同时也决定了 left = mid+1 和 right = mid
