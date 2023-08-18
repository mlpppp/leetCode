#
# @lc app=leetcode id=15 lang=python3
#
# [15] 3Sum
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
    def sortedTwoSumPairs(self, nums: List[int], start: int, target: int):     
        # find all unique value pairs and save these pairs to an array 
        pLeft = start
        pRight = len(nums)-1
        result = []

        while (pLeft < pRight):
            sumTwo = nums[pLeft] + nums[pRight]
            valLeft = nums[pLeft]
            valRight = nums[pRight]

            if (sumTwo == target):
                result.append([valLeft, valRight])
                while (nums[pLeft]==valLeft and pLeft < pRight):
                    pLeft+=1
                while (nums[pRight]==valRight and pLeft < pRight):
                    pRight-=1    

            elif (sumTwo < target):
                while (nums[pLeft]==valLeft and pLeft < pRight):
                    pLeft+=1
            elif (sumTwo > target):
                while (nums[pRight]==valRight and pLeft < pRight):
                    pRight-=1
        return result
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        result = []
        isAppended = False
        appendNum = None
        for idx, num in enumerate(nums):

            if (isAppended):
                if (num == appendNum):
                    continue
                else:
                    isAppended = False
                    appendNum = None

            twoSumsPairs = self.sortedTwoSumPairs(nums, idx+1, 0-num)

            for pair in twoSumsPairs:
                result.append([num, pair[0], pair[1]])
            isAppended = True
            appendNum = num

        return result



            


        
        
# @lc code=end

