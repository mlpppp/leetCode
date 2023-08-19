# @before-stub-for-debug-begin
from python3problem76 import *
from typing import *
# @before-stub-for-debug-end

#
# @lc app=leetcode id=76 lang=python3
#
# [76] Minimum Window Substring
#

# @lc code=start
class Solution:
    def isVaild(self, charRequirement):
        # isVaild() => {return true if charCount's val all == 0}
        # print(charRequirement)
        for char in charRequirement:
            if charRequirement[char] > 0:
                return False
        return True
    def minWindow(self, s: str, t: str) -> str:
        # init optimums
        minLen = 1e9
        minLeftP = minRightP = 0

        leftP = rightP = 0

        # charRequirement
            #  {"charIn_t": "count of the char in substring"; init as required appearance}
        charRequirement = dict.fromkeys(t, 0)
        for char in t:
            charRequirement[char]+=1


        while (rightP < len(s)):
            # print ('start expand')
            curRChar = s[rightP]
            if (curRChar in charRequirement): 
                charRequirement[curRChar]-=1
            # print(s[leftP:rightP+1])
            isVaild = self.isVaild(charRequirement)
            # print('-----')

            while (isVaild):
                # print('start shrink')
                # Update global optimum
                if (rightP-leftP < minLen):
                    minLeftP = leftP
                    minRightP = rightP+1
                    minLen = rightP-leftP

                # shrink window
                curLChar = s[leftP]
                if curLChar in charRequirement:
                    charRequirement[curLChar]+=1
                leftP+=1 
                #re-eval vaildity 
                # print(s[leftP:rightP+1]) 
                isVaild = self.isVaild(charRequirement)
                # print('-----')
   
            
            rightP+=1
        # print(minLeftP)
        # print(minRightP)
        return s[minLeftP: minRightP]
        
# @lc code=end

# 在字符串 S 中使用双指针中的左右指针技巧，初始化 left = right = 0，把索引左闭右开区间 [left, right) 称为一个「窗口」。

# 先不断地增加 right 指针扩大窗口 [left, right)，直到窗口中的字符串符合要求（包含了 T 中的所有字符），此时更新最优的记录。: 相当于在寻找一个「可行解」

# 此时，我们停止增加 right，转而不断增加 left 指针缩小窗口 [left, right)，直到窗口中的字符串不再符合要求（不包含 T 中的所有字符了）。同时，每次增加 left，我们都要更新一轮最优的记录。

# 4、重复第 2 和第 3 步，左右指针轮流前进，窗口大小增增减减，窗口不断向右滑动, 直到 right 到达字符串 S 的尽头。

# O(n), 每一个字符最多进窗口一次然后出窗口一次