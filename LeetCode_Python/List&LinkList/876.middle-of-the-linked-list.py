#
# @lc app=leetcode id=876 lang=python3
#
# [876] Middle of the Linked List
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def middleNode(self, head: Optional[ListNode]) -> Optional[ListNode]:
        p_fast = head
        p_slow = head
        while(p_fast and p_fast.next):
            p_slow = p_slow.next
            p_fast = p_fast.next.next
        return p_slow
        
# 快慢指针慢指针走一，快指针走二，快指针走到的时候，慢指针真指向就是终点
# @lc code=end

