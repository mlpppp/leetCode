package LeetCode_Java.List_LinkedList;

import java.util.HashMap;

//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    public boolean containsDuplicate(int[] nums) {
        HashMap<Integer, Integer> counts = new HashMap<Integer, Integer>();
        for (int num : nums) {
            if (counts.containsKey(num)) {
                return true;
            } else {
                counts.put(num, 1);
            }
        }
        return false;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
