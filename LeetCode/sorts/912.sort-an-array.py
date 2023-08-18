#
# @lc app=leetcode id=912 lang=python3
#
# [912] Sort an Array
#

# @lc code=start
import random
class Solution:
    def swap(self, nums, i, j):
        if i == j:
            return 
        else:
            temp = nums[i]
            nums[i] = nums[j]
            nums[j] = temp
            return 
    def partition(self, nums, left, right):
        pivot = nums[right]
        pEdit = pSearch = left
        while (pSearch < right):
            if (nums[pSearch] <= pivot):
                self.swap(nums, pEdit, pSearch)
                pEdit+=1
            pSearch+=1
        self.swap(nums, pEdit, right)
        return pEdit
    # def partition(self, nums, left, right):
    #     pivot = nums[left]
    #     pLeft = left+1 
    #     pRight = right
    #     while(pLeft <= pRight):
    #         while(pLeft < right and nums[pLeft]<=pivot):
    #             pLeft+=1
    #         while(pRight > left and nums[pRight]>pivot):
    #             pRight-=1
    #         if(pLeft >= pRight):
    #             break
    #         self.swap(nums, pLeft, pRight)
    #     self.swap(nums, left, pRight)          
    #     return pRight

    def quickSort(self, nums, left, right):
        if left >= right: return None
        mid = self.partition(nums, left, right)
        self.quickSort(nums, left, mid-1)
        self.quickSort(nums, mid+1, right)

    def sortArray(self, nums: List[int]) -> List[int]:
        random.shuffle(nums)
        self.quickSort(nums, 0, len(nums)-1)
        return nums


        
# @lc code=end

# it is just quick sort