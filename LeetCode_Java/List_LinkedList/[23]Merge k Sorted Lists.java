package LeetCode_Java.List_LinkedList;
//leetcode submit region begin(Prohibit modification and deletion)

import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

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
// public class ListNode {
//     int val;
//     ListNode next;
//     ListNode() {}
//     ListNode(int val) { this.val = val; }
//     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
// }

 class Solution {
    public ListNode mergeKLists(ListNode[] lists) {
        if (lists.length == 0  || lists == null) {
            return null;
        }
        List<ListNode> result = new ArrayList<ListNode>(Arrays.asList(lists));
        while (result.size() > 1) {
            int size = result.size();
            List<ListNode> newList = new ArrayList<>();
            int i = 0;
            for (i = 0; i+1 < size; i+=2){
                // merge two lists
                ListNode merged = this.merge2Lists(result.get(i), result.get(i+1));

                // append to the newList
                newList.add(merged);
            }
            // in odd case, there are one more list
            if (i < size) {
                newList.add(result.get(i));
            }
            result = newList;
        }
        return result.get(0);
    }
    public ListNode merge2Lists(ListNode list1, ListNode list2) {
        if (list1 == null && list2 == null) {
            return null;
        }
        ListNode dummy = new ListNode(0, null);
        ListNode cur = dummy;
        while (list1 != null && list2 != null) {
            if (list1.val < list2.val) {
                cur.next = list1;
                list1 = list1.next;
            } else {
                cur.next = list2;
                list2 = list2.next;
            }
            cur = cur.next;
        }
        cur.next = Objects.requireNonNullElse(list1, list2);
        return dummy.next;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
