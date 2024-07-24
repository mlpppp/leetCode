package LeetCode_Java.List_LinkedList;
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */

class Solution {
    public ListNode mergeTwoLists(ListNode list1, ListNode list2) {
        // empty case
        if (list1 == null && list2 == null) {
            return null;
        }
        ListNode dummy = new ListNode(0, null);
        ListNode head = dummy;
        while (list1 != null && list2 != null) {
            if (list2.val < list1.val) {
                head.next = list2;
                list2 = list2.next;
            }
            else {
                head.next = list1;
                list1 = list1.next;
            }
            head = head.next;
        }
        if (list1 == null) {
            head.next = list2;
        } else {
            head.next = list1;
        }
        return dummy.next;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
