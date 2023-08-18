# @before-stub-for-debug-begin
from python3problem23 import *
from typing import *
# @before-stub-for-debug-end

#
# @lc app=leetcode id=23 lang=python3
#
# [23] Merge k Sorted Lists
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
# from queue import PriorityQueue
import heapq

class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        setattr(ListNode, "__lt__", lambda self, other: self.val < other.val)
        setattr(ListNode, "__gt__", lambda self, other: self.val > other.val)
        setattr(ListNode, "__ge__", lambda self, other: self.val >= other.val)
        setattr(ListNode, "__le__", lambda self, other: self.val <= other.val)


        PriorQue = []
        for linkList in lists:
            if linkList is not None:
                heapq.heappush(PriorQue, linkList)
        dummy = ListNode(-1)
        p = dummy
        while PriorQue:
            argmin = heapq.heappop(PriorQue)
            if (argmin and argmin.next):
                heapq.heappush(PriorQue, argmin.next)

            p.next = argmin
            p = p.next 
            p.next = None

        return dummy.next

        

# 用binary heap实现的PriorityQueue来对N个input进行快速的取出最小值, 

# N个指针分别指向N个Link list, p指向output链表

# 把N个指针指向的当前对象全部加到PQ里然后让PQ出队列
#   pq出队列的k:=argMax(pq)加入到p (output)的队列里
#   再把k.next 加入到pq里

#complexity:
    # pq.poll/add: o(logn)
    # 每一个节点都会被poll一遍
    # 如果ListNode总和为K， O(Klog)

# @lc code=end

