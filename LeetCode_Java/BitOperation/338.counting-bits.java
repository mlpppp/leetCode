package LeetCode_Java.BitOperation;
/*
 * @lc app=leetcode id=338 lang=java
 *
 * [338] Counting Bits
 */

// @lc code=start
class Solution {
    public int[] countBits(int n) {
        int[] result = new int[n+1];
        for (int i = 0; i < result.length; i++) {
            result[i] = 0;
        }

        for (int i = 0; i < result.length; i++) {
            result[i] = result[i>>1] + (i&1);
        }
        return result; 
    }
    public int[] brutalForce(int n) {
        int[] result = new int[n+1];
        for (int i=0; i < n+1; i++) {
            int count = 0;
            int currentNum = i;
            while (currentNum > 0) {
                count += currentNum&1;
                currentNum >>= 1;
            }
            result[i] = count;
        }
        return result;
    }
}
// @lc code=end

