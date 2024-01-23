/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start

// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays in O(log (m+n))
// https://www.youtube.com/watch?v=q6IEA26hvXc
// To find the median of the merged array, it can be roughly partition into two parts, either (even total len) two parts are the same length and the median is (part1[-1] + part2[0])/2, or (odd total len) the part1 is one element less and the median is part2[0]. We do binary search to partition nums1, and deduce the partition of nums2 based on `the size of nums1's left partition` and `size of part1`. Given index mid that partition nums1, there are mid+1 elements of num1 that contribute to part1, and then we can deduce that there are floor((len(nums1)+len(nums2)/2)-(mid+1) element in the num2 that contribute to part1. Then we test if the proposed partition really produces the sorted merged array (aka every elements in the part1 is smaller or equal to every elements in the part2), if so, the median is found. if not, we adjust the nums1' partition. Note1: watch for out of bound edgy cases. if either the partition to nums1 or nums2 are out of bounds, we assign dummy value inf or -inf. Note2: strict property of the array to be searched: it must be the longer array; some part of nums1 must be in the part1, otherwise the correct median cannot be found. (see implementation for fixes)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// ? make sure nums1 in which we do binary search is the longer array (why?)
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	// constants
	partOneLen := (len(nums1) + len(nums2)) / 2
	var isEven bool
	if (len(nums1)+len(nums2))&1 == 0 {
		isEven = true
	} else {
		isEven = false
	}

	// edge case1: one of the array is empty
	if len(nums2) == 0 {
		mid := (len(nums1) - 1) / 2
		if isEven {
			return (float64(nums1[mid]) +
				float64(nums1[mid+1])) / 2
		} else {
			return float64(nums1[mid])
		}
	}

	// start binary search
	left, right := 0, len(nums1)-1
	var result float64
	for left <= right {
		mid := (left + right) / 2
		mid2 := partOneLen - (mid + 1) - 1
		// avoid taking too much /too less such that nums2 cannot handle
		if mid2 >= len(nums2) {
			left = mid + 1
			continue
		} else if mid2 < -1 {
			right = mid - 1
			continue
		}
		// manage edge cases (mid/ mid2 or mid+1/mid2+1 out of range)
		var midValue, midValue2 float64
		if mid >= 0 {
			midValue = float64(nums1[mid])
		} else { // no nums1 in part1
			midValue = -math.MaxFloat64
		}
		if mid2 >= 0 {
			midValue2 = float64(nums2[mid2])
		} else { // no nums2 in part1
			midValue2 = -math.MaxFloat64
		}

		var midPlusValue, midPlusValue2 float64
		if mid+1 < len(nums1) { // all nums1 in part1
			midPlusValue = float64(nums1[mid+1])
		} else {
			midPlusValue = math.MaxFloat64
		}
		if mid2+1 >= len(nums2) { // all nums2 in part1
			midPlusValue2 = math.MaxFloat64
		} else {
			midPlusValue2 = float64(nums2[mid2+1])
		}

		if midValue <= midPlusValue2 && midValue2 <= midPlusValue {
			// a vaild partition with the correct size, the median is found
			if isEven {
				result = (max(midValue, midValue2) +
					min(midPlusValue, midPlusValue2)) /
					2
				break
			} else {
				result = min(midPlusValue, midPlusValue2)
				break
			}
		} else if midValue > midPlusValue2 {
			// nums1 boundary is too big, reduce elements from nums1 in the part1
			right = mid - 1
		} else if midValue2 > midPlusValue {
			// nums2 boundary is too big, increase elements from nums1 in the part1
			left = mid + 1
		}
	}

	// edge case2: nums1 binary search terminated cause right < left, right < 0. Because everything in nums1 ought to go to part2
	if right < 0 { // we know that the separator mid (ie last element in part1) must in nums2, so partOneLen <= len(nums1)
		midValue := float64(nums2[partOneLen-1])
		var midPlusValue float64
		if partOneLen == len(nums2) {
			midPlusValue = float64(nums1[0])
		} else if partOneLen < len(nums1) {
			midPlusValue = float64(nums1[partOneLen])
		}
		if isEven {
			return (midPlusValue + midValue) / 2
		} else {
			return midPlusValue
		}
	}

	return result
}

// @lc code=end

