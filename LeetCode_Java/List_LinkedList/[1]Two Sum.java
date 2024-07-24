package LeetCode_Java.List_LinkedList;

import java.util.Arrays;
import java.util.Comparator;

//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    // O(nlogn) sorted then two pointers
    // O(n) hashmap save target-nums[i]

    // hashmap solution
//    public int[] twoSum(int[] nums, int target) {
//        Map<Integer, Integer> remainderToIdx = new HashMap<Integer, Integer>();
//        for (int i = 0; i < nums.length; i++) {
//            if (remainderToIdx.containsKey(nums[i])) {
//                return new int[]{remainderToIdx.get(nums[i]), i};
//            } else {
//                remainderToIdx.put(target - nums[i], i);
//            }
//        }
//        return new int[]{-1, -1}; // throws an exception. Not a valid solution
//    isSubtree}

    // sorted solution
    public int[] twoSum(int[] nums, int target) {
        int[][] tuples = new int[nums.length][2];
        for (int i = 0; i < nums.length; i++) {
            tuples[i][0] = nums[i];
            tuples[i][1] = i;
        }
        Arrays.sort(tuples, new Comparator<int[]>(){

            @Override
            public int compare(int[] o1, int[] o2) {
                return Integer.compare(o1[0], o2[0]);
            }
        });
//        System.out.println(Arrays.deepToString(tuples));
        int left = 0, right = tuples.length-1;
        int[] result = new int[2];
        while(left < right) {
            int sum = tuples[left][0] + tuples[right][0];
            if (target < sum) {
                right--;
            } else if (target > sum) {
                left++;
            } else {
//                System.out.printf("left: %d, right: %d", left, right);
                result[0] = tuples[left][1];
                result[1] = tuples[right][1];
                break;
            }
        }
        return result;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
