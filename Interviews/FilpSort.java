package Interviews;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

// imagine a sub-array exist, that is. We call it "sa"
    // partially naturally-sorted, 
    // starting with 1
    // either filped or not 
// we use a data structure to track the sub-array
    // saIsFlip: Boolean
    // saBeginIdx: Integer
    // salength: Integer

// the FilpSort algorithm dynamically build the sub-array to sort the whole array

// to increase the size of "sa", there are 4 scenerio
    // 1. "idx" of next number after "sa"
        // flip "sa", then 2,
    // 2. "idx" of next number after "sa", "sa" fliped
        // flip everything till idx-1, then length of sa + 1, update "sa" data
    // 3. "idx" of next number before "sa"
        // flip everything till the end of "sa", then 2
    // 4. "idx" of next number before "sa", "sa" fliped
        // flip "sa", then 3




public class FilpSort {
    private List<Integer> kHistory = new ArrayList<>();
    private Boolean saIsFlip;
    private Integer saBeginIdx;
    private Integer salength;
    private Integer nextIdx;

    private void partialFlip(int[] nums, int leftIdx, int rightIdx) {
        while (leftIdx <= rightIdx) {
            int temp = nums[rightIdx];
            nums[rightIdx] = nums[leftIdx];
            nums[leftIdx] = temp;
            leftIdx++; rightIdx--;
        }
    }

    private boolean isSorted(int[] nums, int leftIdx, int rightIdx) {
        if (nums.length <= 0) return true;
        for (int i = leftIdx; i < rightIdx; i++) {
            int j = i+1;
            if (nums[j] < nums[i]) {
                return false;
            }
        }
        return true;
    }

    private void filpSort(int[] nums) {
        // find initial sa
        for (int i=0; i<nums.length; i++) {
            if (nums[i] == 1) {
                saIsFlip = true;
                saBeginIdx = i;
                salength = 1;
            }
        }
        while (salength < nums.length) {
            // find the next integer's idx
            int nextInt = salength+1;
            for (int i=saBeginIdx; i>=0; i--) {
                if (nums[i] == nextInt) {
                    nextIdx = i;
                }
            }
            for (int i=saBeginIdx+salength; i<nums.length; i++) {
                if (nums[i] == nextInt) {
                    nextIdx = i;
                }
            }

            // 1. "idx" of next number after "sa"
                // flip "sa", then 2,
            if (nextIdx > saBeginIdx && !saIsFlip) {
                partialFlip(nums, saBeginIdx, saBeginIdx+salength-1);
                saIsFlip = !saIsFlip;
            }
            // 2. "idx" of next number after "sa", "sa" fliped
                // flip everything till idx-1, then length of sa + 1, update "sa" data
            if (nextIdx > saBeginIdx && saIsFlip) {
                partialFlip(nums, 0, nextIdx-1);
                salength++;
                saBeginIdx = nextIdx-salength+1;
                saIsFlip = !saIsFlip;
            }
            // 3. "idx" of next number before "sa"
                // flip everything till the end of "sa", then 2
            if (nextIdx < saBeginIdx && !saIsFlip) {    
                partialFlip(nums, 0, saBeginIdx+salength-1);
                saIsFlip = !saIsFlip;
                saBeginIdx = 0;
            }
            // 4. "idx" of next number before "sa", "sa" fliped
                // flip "sa", then 3
            if (nextIdx < saBeginIdx && saIsFlip) {    
                partialFlip(nums, saBeginIdx, saBeginIdx+salength-1);
                saIsFlip = !saIsFlip;
            }
        }

    }
    
    public static void main(String[] args) {
        FilpSort filpSort = new FilpSort();
        int[] nums = {1,3,6,7,9,2,5,4,8, 10};
        filpSort.filpSort(nums);
        System.out.println(Arrays.toString(nums));
    }

    // write some unit test
    private void partialFlip_test() {
        int[] nums = {3,2,1,4};
        partialFlip(nums, 0, 3);
        System.out.println(Arrays.toString(nums));
    }

    private void isSorted_test() {
        int[] nums = {3,2,1,4};
        int[] nums_sorted = {1,2,3,4};
        System.out.println(isSorted(nums, 0, 3));
        System.out.println(isSorted(nums_sorted, 0, 3));
    }

    private void filpSort_test() {
        int[] nums = {3,2,1,4};
        filpSort(nums);
        System.out.println(Arrays.toString(nums));
        System.out.println(kHistory.toString());
    }
}