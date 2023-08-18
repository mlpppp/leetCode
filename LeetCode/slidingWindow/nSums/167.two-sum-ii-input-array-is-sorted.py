#
# @lc app=leetcode id=167 lang=python3
#
# [167] Two Sum II - Input Array Is Sorted
#

# @lc code=start
class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        leftP = 0
        rightP = len(numbers)-1
        while (leftP < rightP):
            twoSum = numbers[leftP] + numbers[rightP]
            if(twoSum == target):
                return [leftP+1, rightP+1]
            if(twoSum < target):
                leftP+=1
            else:
                rightP-=1
        return [-1, -1]



# 在数组头和数组尾放置窗口此时窗口为最大
# 如果窗口和比target小，从数组头向右缩小窗口以增加sum      
# 如果窗口和比target小，从数组尾向左缩小窗口以减少sum       
# O(N)

# @lc code=end

