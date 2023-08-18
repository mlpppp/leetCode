#
# @lc app=leetcode id=26 lang=python3
#
# [26] Remove Duplicates from Sorted Array
#

# @lc code=start
class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        edit_p = 0
        explore_p = 0
        while (explore_p < len(nums)):
            if(nums[explore_p] > nums[edit_p]):
                nums[edit_p+1] = nums[explore_p]
                edit_p+=1
            explore_p+=1


        return edit_p+1

# 2个指针一个负责探索一个负责改
# 探索指针explore_p只要不碰到数组尾部，就不断往前走。只要发现比当前改指针edit_p大的数就修改改指针edit_p的下一个，并且把edit_p加1
# @lc code=end

