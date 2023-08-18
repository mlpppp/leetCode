#
# @lc app=leetcode id=19 lang=python3
#
# [19] Remove Nth Node From End of List
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        dummy = ListNode(-1)
        dummy.next = head
        p_fast = dummy
        p_slow = dummy
        for i in range(n+1):
            p_fast = p_fast.next
        while (p_fast):
            p_fast = p_fast.next
            p_slow = p_slow.next

        next = p_slow.next.next
        del p_slow.next
        p_slow.next = next

        return dummy.next

# 快慢2个指针快指针先走n+1步, 然后慢指针再跟着快指针一起走，在快指针到None的时候慢指针的下一个就是目标
# 使用dummy节点可以避免n == len(list) 时快指针在第n步就到达None
# @lc code=end

