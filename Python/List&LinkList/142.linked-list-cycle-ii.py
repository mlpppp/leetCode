#
# @lc app=leetcode id=142 lang=python3
#
# [142] Linked List Cycle II
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def detectCycle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        p_fast = head
        p_slow = head
        while (p_fast and p_fast.next):
            p_fast = p_fast.next.next
            p_slow = p_slow.next
            if (p_fast is p_slow):  # exist loop
                break   
        if (p_fast is None or p_fast.next is None):
            return None
        p_fast = head
        while (p_slow is not p_fast):
            p_slow = p_slow.next 
            p_fast = p_fast.next

        return p_slow









# 快慢指针快指针走二慢指针走一，
# 若fast遇到null节点，则表示链表无环，结束；
# 若链中有环，fast和slow一定会再次相遇，相遇时，slow不可能绕环超过1圈。；
    # FAST肯定会首先进入环里面并且有可能在环里面loop
    # slow到达环的起点的时候FAST要么在还的起点(满足环判断标准)，要么在环内的某一处
    # 假定环长为r
    # slow到达环的起点时，FAST在环内某一处，并且和Slow的顺时针距离为n (n<r)。 那么问题变成fast追slow，追击距离为(m = l-n) FAST和Slow每走一个回合他们之间的距离都会变成m-1， 并且最终一定会变成0 => 再次相遇  
    # 也就是fast和slow在slow进入后m个回合后相遇.
    # 因为m < r，相遇时，slow不可能绕环超过1圈。 

# 当fast和slow相遇时，额外创建指针ptr，并指向链表头部，且每次移1步，最终slow和ptr会在入环点相遇。
    # 假设Slow进入环以后再走了m与fast相遇，此时slow总共走了k步，则环的起点距离head必为k-m
    # fast相遇时必走了2k步，假设环长为r, fast相遇时已经在环内loop n 圈，有
        # 2k = nr + (k-m) + m 
        # 消元 => k = nr 
        # => 可知k定为环长度的整数倍 
        
    # 额外创建指针ptr重新走k-m步到环入口，由已知有
        # k-m = nr-m = (n-1)r + (r-m)
    # r-m 正好可以让slow走回入口，(n-1)r 让它再loop (n-1) 圈      

# complexity: 
    # 假设链表有N个元素
    # slow在第一次相遇时不会走完环的第一圈，也即 < N
    # fast: < 2N, ptr < N
    # so O(N)
# @lc code=end

