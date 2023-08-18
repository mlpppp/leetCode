#
# @lc app=leetcode id=83 lang=python3
#
# [83] Remove Duplicates from Sorted List
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if(not head):
            return None
        pEdit = head 
        pExplore = head
        while (pExplore):
            if(pExplore.val > pEdit.val):
                pEdit.next = pExplore
                pEdit = pEdit.next
            pExplore = pExplore.next
        pEdit.next = None
        return head
# @lc code=end

