# @before-stub-for-debug-begin
from python3problem5 import *
from typing import *
# @before-stub-for-debug-end

#
# @lc app=leetcode id=5 lang=python3
#
# [5] Longest Palindromic Substring
#

# @lc code=start



class Solution:
    # 遍历每一个位置i搜索i的Palindrome
    def longestPalindrome(self, s: str) -> str:
        longestP = s[0]
        for i in range(len(s)):
            paliA = self.searchPalindrome(s, i, i)
            paliB = self.searchPalindrome(s, i, i+1)
            longestP = paliA if (len(paliA) > len(longestP)) else longestP
            longestP = paliB if (len(paliB) > len(longestP)) else longestP
        return longestP

    # 在s搜索中心处包含index i的Palindrome, 返回搜索结果
    def searchPalindrome(self, string: str, l: int, r:int) -> str:  
        while (l >= 0 and r <= len(string)-1 and string[l] == string[r]):
            l-=1
            r+=1
        return string[l+1:r]




        
# @lc code=end

