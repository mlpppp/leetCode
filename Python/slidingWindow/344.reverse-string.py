#
# @lc app=leetcode id=344 lang=python3
#
# [344] Reverse String
#

# @lc code=start
class Solution:
    def reverseString(self, s: List[str]) -> None:
        """
        Do not return anything, modify s in-place instead.
        """
        leftP = 0
        rightP = len(s)-1
        while(leftP < rightP):
            # sawp()
            temp = s[leftP] 
            s[leftP] = s[rightP]
            s[rightP] = temp
            # move
            leftP += 1
            rightP -= 1




        
# @lc code=end

