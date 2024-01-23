/*
 * @lc app=leetcode id=410 lang=golang
 *
 * [410] Split Array Largest Sum
 */

// @lc code=start

// Given an integer array nums and an integer k, split nums into k non-empty subarrays (i.e. non-empty slices), such that the largest sum of any subarray is minimized. Return the minimized largest sum.

// brute force, given k, there are combination(len(nums), k-1) ways of splitting, enumerate all splitting and calculate sum for each subarray to find the smallest. O(n*combination(n, k-1))

// Assuming we have the largest sum of any subarray (think capacity) from a already splited array, it is possible to find the `minimal number of subarrays` through a determinstic function: findMinSplits(nums, capacity) -> k. The input capacity is strictly negatively correlated to k. Run binary search in capacity-domain [max(nums), sum(nums)] for findMinSplits to find the left boundary of capacity that yields the target k.

func splitArray(nums []int, k int) int {
	// initialize search range [max(nums), sum(nums)]
	minCapacity, maxCapacity := 0, 0
	for i := range nums {
		minCapacity = max(minCapacity, nums[i])
		maxCapacity += nums[i]
	}

	// start left boundary binary search
	result := math.MaxInt
	for minCapacity <= maxCapacity {
		mid := (minCapacity + maxCapacity) / 2
		minSplit := findMinSplits(nums, mid)
		if minSplit > k { // too much split, increase capacity
			minCapacity = mid + 1
		} else if minSplit <= k { // still has room for capacity optimizations
			result = mid
			maxCapacity = mid - 1
		}
	}
	return result
}

func findMinSplits(nums []int, capacity int) int {
	k := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if i+1 >= len(nums) || sum+nums[i+1] > capacity {
			k++
			sum = 0
		}
	}
	return k
}

// @lc code=end

