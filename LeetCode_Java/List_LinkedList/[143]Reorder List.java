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
// reverse the second half of the list (middle or the first node in the second half), then two pointers, from two side to the middle, we build the solution.
//1. find the half: fast/slow pointers
//2. reverse(mid)
//3. build the solution
class Solution {

    public void reorderList(ListNode head) {
        if (head == null || head.next == null) return;
        // find middle: fast slow pointers
        ListNode dummy = new ListNode(0, head);
        ListNode slow = dummy, fast = head;
        while (fast != null && fast.next != null) {
            slow = slow.next;
            fast = fast.next.next;
        }
        ListNode list2 = null;
        list2 = slow.next;
        slow.next = null;

        // reverse second half
        ListNode prev =null, curr = list2;
        while (curr != null) {
            ListNode tmp = curr.next;
            curr.next = prev;
            prev = curr;
            curr = tmp;
        }
        list2 = prev;

//        ListNode l1 = head;
//        ListNode l2 = list2;
//        while (l1 != null) {
//            System.out.println(l1.val);
//            l1 = l1.next;
//        }
//        System.out.println("------");
//        while (l2 != null) {
//            System.out.println(l2.val);
//            l2 = l2.next;
//        }
//        System.out.println("--------------------------------");
//
        // concatenate the two
        curr = head;
        while (curr.next != null) {
            ListNode tmp1 = curr.next, tmp2 = list2.next;
            curr.next = list2;
            curr = tmp1;
            list2.next = curr;
            list2 = tmp2;
        }
        if (list2 != null) {
            curr.next = list2;
        }
    }
}
//leetcode submit region end(Prohibit modification and deletion)
