#
# @lc app=leetcode id=704 lang=python3
#
# [704] Binary Search
#

# @lc code=start
import math
class Solution:
    def search(self, nums: List[int], target: int) -> int:
        ub = len(nums)-1
        lb = 0
        while (lb <= ub):
            mid = math.floor(lb + (ub - lb) / 2)
            if(nums[mid] == target):
                return mid
            elif (nums[mid] > target):
                ub = mid-1
            elif (nums[mid] < target):
                lb = mid+1
        return -1  # target比所有数都大或小，越界
        
# @lc code=end

# > 小心整形溢出情况，用 mid = left + (right - left) / 2; 代替 mid = (left + right)/ 2;

#  两种情况下停止循环第一是搜索区间为空，失败=>while(lb <= ub))；第二是找到结果

# 因为使用了 lb <= ub 作为搜索循环条件，希望失败时 lb > ub, 因此应该在 lb+1 或者 up-1 取新的区间, 以使其最终会错位


# [寻找边界的二分搜索]
    # 比如说给你有序数组 nums = [1,2,2,2,3]，target 为 2，此算法返回的索引是 2，没错。但是如果我想得到 target 的左侧边界，即索引 1，或者我想得到 target 的右侧边界，即索引 3

def left_bound(self, nums: List[int], target: int) -> int:
    ub = len(nums)-1
    lb = 0
    while (lb <= ub):
        mid = math.floor(lb + (ub - lb) / 2)
        if(nums[mid] == target):
            ub = mid-1  #!
        elif (nums[mid] > target):
            ub = mid-1
        elif (nums[mid] < target):
            lb = mid+1
    if (lb == len(nums)) return -1  # 此时 target 比所有数都大，返回 -1
    if (lb == -1) return -1  # 此时 target 比所有数都小，返回 -1
    if (nums[lb] == target) return lb

def right_bound(self, nums: List[int], target: int) -> int:
    ub = len(nums)-1
    lb = 0
    while (lb <= ub):
        mid = math.floor(lb + (ub - lb) / 2)
        if(nums[mid] == target):
            lb = mid+1  #!
        elif (nums[mid] > target):
            ub = mid-1
        elif (nums[mid] < target):
            lb = mid+1
    if (ub == len(nums)) return -1  # 此时 target 比所有数都大，返回 -1
    if (ub == -1) return -1  # 此时 target 比所有数都小，返回 -1
    if (nums[ub] == target) return ub

