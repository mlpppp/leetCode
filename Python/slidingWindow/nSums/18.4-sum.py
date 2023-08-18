#
# @lc app=leetcode id=18 lang=python3
#
# [18] 4Sum
#

# @lc code=start
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
    def threeSums(self, nums: List[int], start, target: int) -> List[List[int]]:

        result = []

        isAppended = False
        appendNum = None
        for idx in range(start, len(nums)):
            num = nums[idx]
            if (isAppended):
                if (num == appendNum):
                    continue
                else:
                    isAppended = False
                    appendNum = None

            twoSumsPairs = self.sortedTwoSumPairs(nums, idx+1, target-num)

            for pair in twoSumsPairs:
                result.append([num, pair[0], pair[1]])
            isAppended = True
            appendNum = num

        return result
    def fourSum(self, nums: List[int], target: int) -> List[List[int]]:
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
            threeSums = self.threeSums(nums, idx+1, target-num)
            for pair in threeSums:
                result.append([num, pair[0], pair[1], pair[2]])
            isAppended = True
            appendNum = num
        return result

        
# @lc code=end


        
        
# @lc code=end

