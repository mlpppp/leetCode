#
# @lc app=leetcode id=438 lang=python3
#
# [438] Find All Anagrams in a String
#

# @lc code=start
class Solution:
    def isPotentialSol(self, charRequirement):
        for char in charRequirement:
            if charRequirement[char] > 0:
                return False
        return True

    def findAnagrams(self, s: str, p: str) -> List[int]:
        leftP = rightP = 0
        anaList = []
        charRequirement = dict.fromkeys(p, 0)
        for char in p:
            charRequirement[char]+=1

        while (rightP < len(s)):

            # update curRChar and solution status 
            curRChar = s[rightP]
            if curRChar in charRequirement:
                charRequirement[curRChar] -=1  

            while (self.isPotentialSol(charRequirement)):

                if(rightP-leftP == len(p)-1):
                    anaList.append(leftP)
                    break

                curLChar = s[leftP]
                if curLChar in charRequirement:
                    charRequirement[curLChar] +=1  
                leftP+=1

            rightP+=1
        return anaList

        
# @lc code=end

# 与 567 寻找permutation是同一道题