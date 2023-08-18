#
# @lc app=leetcode id=1 lang=python3
#
# [1] Two Sum
#

# @lc code=start
class idxPair:
        def __init__(self, idx, val) -> None:
            self.idx = idx
            self.val = val
        def __lt__(self, other):
            return self.val < other.val
        def __gt__(self, other):
            return self.val > other.val
        def __ge__(self, other):
            return self.val >= other.val
        def __le__(self, other):
            return self.val <= other.val
        def __eq__(self, other):
            return self.val == other.val
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:    
        numsPair = []
        for idx, val in enumerate(nums):
            numsPair.append(idxPair(idx, val))

        numsPair.sort()
        # same as twoSum sorted (167)
        pLeft = 0
        pRight = len(numsPair)-1
        while (pLeft < pRight):
            sumTwo = numsPair[pLeft].val + numsPair[pRight].val
            if (sumTwo == target):
                return [numsPair[pLeft].idx, numsPair[pRight].idx]
            elif (sumTwo < target):
                pLeft+=1
            elif (sumTwo > target):
                pRight-=1
           

# @lc code=end

# Two pointers point to the left most and the right most of the array
# When the left pointer moved to right the two sum increase by its value
# When the right pointer moved to the left the two sum decrease by its value
# we dynamically adjust the left and left pointers according to how the current sum comparing tothe target
