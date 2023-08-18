#
# @lc app=leetcode id=3 lang=python3
#
# [3] Longest Substring Without Repeating Characters
#

# @lc code=start
class Solution:
    def isNoRepeat(self, charCounter):
        for key in charCounter:
            if(charCounter[key] > 1):
                return False
            if (charCounter[key] < 0):
                raise ValueError('should not less than 0')
        return True
    def lengthOfLongestSubstring(self, s: str) -> int:
        leftP = rightP = 0
        maxLen = 0
        charCounter = {}

        while (rightP < len(s)):
            curRchar = s[rightP]
            if curRchar in charCounter:
                charCounter[curRchar] += 1
            else:
                charCounter[curRchar] = 1
            isNoRepeat = self.isNoRepeat(charCounter)
            if (isNoRepeat and rightP-leftP+1 > maxLen):
                maxLen = rightP-leftP+1
            
            while (not self.isNoRepeat(charCounter)):
                curLchar = s[leftP]
                if curLchar in charCounter:
                    charCounter[curLchar] -= 1
                else:
                    raise ValueError("curLchar should always be a key")
                leftP+=1

            rightP+=1   

        return maxLen
# @lc code=end

# 移动窗口
#   只要当前窗口中的子串符合条件（no repeat)， 就右指针向右扩张窗口， 当不符合条件就左指针向右缩小窗口， 直到符合条件