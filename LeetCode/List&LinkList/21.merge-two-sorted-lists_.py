# @before-stub-for-debug-begin
from python3problem21 import *
from typing import *
# @before-stub-for-debug-end


# @lc app=leetcode id=21 lang=python3
#
# [21] Merge Two Sorted Lists
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode(-1)
        p = dummy
        while ((list1 != None) and (list2 != None)):
            if (list1.val < list2.val):
                p.next = list1
                list1 = list1.next
            else:
                p.next = list2
                list2 = list2.next
            p = p.next
        if (list1 != None) :
            p.next = list1
        if (list2 != None) :
            p.next = list2
        return dummy.next

    # 2个指针p1, p2, 一个A头一个B头
    # 当指针不null, 比较2个指针
    # 比较小的一个插到新的链表并且把指针指向当前的下一个
    # O(n)

# @lc code=end

