package LeetCode_Java.List_LinkedList;

import java.util.*;

//leetcode submit region begin(Prohibit modification and deletion)
// 15 Sort the input array. Implement 2Sum(nums, n) to find unique twoSums that sums to n in a sorted array. Iterate the array and for each first encounter of a unique element nums[i] call 2Sum(nums[i+1:], 0-nums[i])
class Solution {
    public List<List<Integer>> twoSum(int[] nums, int i, int target) {
        List<List<Integer>> result = new ArrayList<>();
        int left = i, right = nums.length-1;
        while(left < right) {
            int sum = nums[left] + nums[right];
            if (target > sum) {
                left++;
            } else if (target < sum) {
                right--;
            } else {
                List<Integer> tuple = new ArrayList<>();
                tuple.add(nums[left]);
                tuple.add(nums[right]);
                result.add(tuple);
                // shrink the window without introducing duplicates
                int leftVal = nums[left];
                int rightVal = nums[right];
                while (nums[left] == leftVal && left < right) {
                    left++;
                }
                while (nums[right] == rightVal && left < right) {
                    right--;
                }
            }
        }
        return result;
    }
    public List<List<Integer>> threeSum(int[] nums) {
        Arrays.sort(nums);
        List<List<Integer>> result = new ArrayList<>();
        int lastVal = Integer.MIN_VALUE;
        for (int i = 0; i < nums.length-2; i++) {
            if (nums[i] == lastVal) {  // only compute 2sum for the first unique integer
                continue;
            }
            lastVal = nums[i];
            // try to find the twoSum result
            List<List<Integer>> twoSumResult = twoSum(nums, i+1, 0-nums[i]);
            if (!twoSumResult.isEmpty()) {
                final int num = nums[i];
                twoSumResult.forEach(l -> l.add(0, num));
                result.addAll(twoSumResult);
            }
        }
        return result;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
