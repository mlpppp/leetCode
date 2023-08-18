# @before-stub-for-debug-begin
from python3problem315 import *
from typing import *
# @before-stub-for-debug-end

#
# @lc app=leetcode id=315 lang=python3
#
# [315] Count of Smaller Numbers After Self
#

# @lc code=start
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
from math import floor

class Solution:
    counts = None
    class idxPair:
        val = -1
        idx = -1
        def __init__(self, idx, val):
            self.val = val
            self.idx = idx
        def __lt__(self, other):
            return self.val < other.val
        def __gt__(self, other):
            return self.val > other.val
        def __ge__(self, other):
            return self.val >= other.val
        def __le__(self, other):
            return self.val <= other.val

    def mergeSort(self, nums: List[idxPair]):
        if len(nums) == 1: return nums
        mid = floor(len(nums) / 2)
        sortedLeft = self.mergeSort(nums[0:mid])
        sortedRight = self.mergeSort(nums[mid:])
        merged = self.merge(sortedLeft, sortedRight)
        return merged

    def merge(self, sortedLeft: List[idxPair], sortedRight:List[idxPair]):
        leftP = rightP = 0
        merged = []
        while (leftP < len(sortedLeft) and rightP < len(sortedRight)):
            if (sortedLeft[leftP].val <= sortedRight[rightP].val): 
                # sortedRight's left part must all be smaller than leftP's val
                merged.append(sortedLeft[leftP])
                self.counts[sortedLeft[leftP].idx] += rightP
                leftP+=1
                continue
  
            if (sortedLeft[leftP].val > sortedRight[rightP].val):
                merged.append(sortedRight[rightP])
                rightP+=1
      

        if (leftP == len(sortedLeft)):
            merged.extend(sortedRight[rightP:])
        if (rightP == len(sortedRight)):
            merged.extend(sortedLeft[leftP:])
            # the whole right list must be smaller than every remaining slice of left list
            while (leftP < len(sortedLeft)):
                self.counts[sortedLeft[leftP].idx] += len(sortedRight)
                leftP+=1
        return merged



    def countSmaller(self, nums: List[int]) -> List[int]:   
        numPairs = [self.idxPair(val= nums[i], idx = i) for i in range(len(nums))]

        self.counts = [0]*len(nums)
        sorted = self.mergeSort(numPairs)
        print()
        return self.counts
        
        
# @lc code=end

#! not work
# 用priorityQueue, poll(): logn  实现O(nlogn)算法
# 每次poll一个argmax, 并且计算其右边成员数量

#! use merge sort solution
# https://labuladong.github.io/algo/2/21/41/#其他应用

#! may try binary search tree latter