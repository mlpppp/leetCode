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
    public ListNode reverseList(ListNode head) {
        ListNode first = null, second = head;
        while (second != null) {
            ListNode tmp = second.next;
            second.next = first;
            first = second;
            second = tmp;
        }
        head = first;
        return head;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
