package LeetCode_Java.sorts;

/*
 * @lc app=leetcode id=948 lang=java
 *
 * [912] Sort an Array
 */

// @lc code=start

import java.util.Arrays;

class Solution {
    public int[] sortArray(int[] nums) {
        quickSort(nums, 0, nums.length-1);
        return nums;
    }

    // Perform quickSort for index start to index end in nums[].
    public void quickSort(int[] nums, int start, int end) {
        if (start >= end) return;
        int pivotIdx = partition(nums, start, end);
        quickSort(nums, start, pivotIdx-1);        
        quickSort(nums, pivotIdx+1, end);
    }

    public int partition(int[] nums, int start, int end) {
        int pivotVal = nums[end];
        int pEdit = start, pSearch = start;
        while (pSearch < end) {
            if (nums[pSearch] <= pivotVal) {
                swap(nums, pEdit, pSearch);
                pEdit ++;
            }
            pSearch ++;
        }
        swap(nums, pEdit, end);
        return pEdit;
    }

    public void swap(int[] nums, int i, int j) {
        if (i == j) return;
        int temp =  nums[i];
        nums[i] = nums[j];
        nums[j] =  temp;
        return;
    }

    public static void main(String[] args) {
        int[] nums = {5,1,1,2,0,0};
        Solution sol = new Solution();
        sol.sortArray(nums);
        System.out.println(Arrays.toString(nums));
    }
}