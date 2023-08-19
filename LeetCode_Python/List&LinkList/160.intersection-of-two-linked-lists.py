# @before-stub-for-debug-begin
from python3problem160 import *
from typing import *
# @before-stub-for-debug-end

#
# @lc app=leetcode id=160 lang=python3
#
# [160] Intersection of Two Linked Lists
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> Optional[ListNode]:
        p1 = headA 
        p2 = headB 
        while(p1 is not p2):
            if (p1 is None):
                p1 = headB
            else:
                p1 = p1.next

            if (p2 is None):
                p2 = headA
            else:
                p2 = p2.next

        return p1
        

  # 2个相同速度的指针 p1, p2
  # 设定2个链表的共同部分为k
  # 设定a链表的前置 skA, 设定b链表的前置 skB
  # p1: skA -> a - > skB -> c1
  # p2: skB -> a - > skA -> c1
  # p1 与 p2 必在c1相遇
  # 如果2个链表没有相交 (a为空）他们会在None相遇

# @lc code=end

