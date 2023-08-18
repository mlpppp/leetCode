#
# @lc app=leetcode id=86 lang=python3
#
# [86] Partition List
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def partition(self, head: Optional[ListNode], x: int) -> Optional[ListNode]:
        LMin_dummy =  ListNode(-1)
        LMax_dummy =  ListNode(-1)
        LMin_p = LMin_dummy
        LMax_p = LMax_dummy
        Org_p = head
        while (Org_p is not None):
            if(Org_p.val < x):
                LMin_p.next = Org_p  
                LMin_p = LMin_p.next 
            else:
                LMax_p.next = Org_p  
                LMax_p = LMax_p.next 
            temp = Org_p.next
            Org_p.next = None
            Org_p = temp
        LMin_p.next = LMax_dummy.next
        return  LMin_dummy.next

    # 遍历链表如果Node满足前一个要求就放到第一个子表，满足后一个要求就放到第2子表

    # 每把一个Note放到子表后就把这个Node与它在原表中下一个node的连接断开

    # 最后原表被整个打散, 把2个子表连在一起就可以了
    
    # O(n)
# @lc code=end

