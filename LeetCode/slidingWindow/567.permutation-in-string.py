#
# @lc app=leetcode id=567 lang=python3
#
# [567] Permutation in String
#

# @lc code=start
class Solution:
    def isPotentialSol(self, charRequirement):
        for char in charRequirement:
            if charRequirement[char] > 0:
                return False
        return True
    def checkInclusion(self, s1: str, s2: str) -> bool:
        leftP = rightP = 0
        charRequirement = dict.fromkeys(s1, 0)
        for char in s1:
            charRequirement[char]+=1

        while (rightP < len(s2)):

            # update curRChar and solution status 
            curRChar = s2[rightP]
            if curRChar in charRequirement:
                charRequirement[curRChar] -=1  

            while (self.isPotentialSol(charRequirement)):

                if(rightP-leftP == len(s1)-1):
                    return True
                curLChar = s2[leftP]
                if curLChar in charRequirement:
                    charRequirement[curLChar] +=1  
                leftP+=1

            rightP+=1
        return False
                





# 滑动窗口算法
# 右指针滑动扩大到所有计数都符合
# 左指针缩小到和S2相同的长度
        
# @lc code=end

