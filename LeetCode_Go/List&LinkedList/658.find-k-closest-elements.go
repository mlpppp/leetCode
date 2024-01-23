/*
 * @lc app=leetcode id=658 lang=golang
 *
 * [658] Find K Closest Elements
 */

// @lc code=start

// Given a sorted integer array arr, two integers k and x, return the k closest integers (abs value) to x in the array. The result should also be sorted in ascending order. When two numbers `a, b` where a < x < b and abs(a, x) == abs(b, x), the smaller `a` is preferred.

// sliding windows O(n). Since the array is sorted, the `k closest integers` must be continuous (be a subarray). Use sliding windows, maintain the sum of distance as defined by the integer distance definition. Slide the window as it size to k, and update the minimum sum of distance, save the current window to the result with the minimum sum is updated the first time.

// func findClosestElements(arr []int, k int, x int) []int {
// 	distSum := 0
// 	minDistSum := math.MaxInt
// 	result := [2]int{}
// 	// initial window
// 	for i := 0; i < k-1; i++ {
// 		distSum += int(math.Abs(float64(arr[i] - x)))
// 	}
// 	left, right := 0, k-1
// 	for right < len(arr) {
// 		// expand
// 		distSum += int(math.Abs(float64(arr[right] - x)))
// 		right++
// 		if right-left == k {
// 			if distSum < minDistSum {
// 				minDistSum = distSum
// 				result[0] = left
// 				result[1] = right
// 			}
// 			distSum -= int(math.Abs(float64(arr[left] - x)))
// 			left++
// 		}
// 	}
// 	return arr[result[0]:result[1]]
// }

// binary search: given k, we focus on finding a window (must be subarray) of size k in `arr`. we find the window by searching for the left bound of the window (window[0]'s position). A window's closeness to x can be characterized by the sum(abs(window[i]-x)), and we want to find the window that minimize the sum. Suppose our window[0] is at arr[mid]'s position. !!: We decide how to optimize the window by compare window mid with window mid+1, if window mid+1 yields a smaller sum, then we should search for the window in the right (left = mid+1), if window mid+1 yields a larger sum or equal sum(as the tie breaker rule prefer the left), we should search for the window in the left include the mid (right = mid).
// Implementation: 1. the search range is [0, len(arr)-k], includes all possible windows. 2. design the compare rule of window mid and window mid+1, we only need to consider the difference between window[mid] and window[mid+k]. 3. when window mid is the last window, we cannot compare window mid with the next window (mid+k >= len(arr), window mid+1 will out of bound), we also search in the left include the mid

func findClosestElements(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-k
	result := -1
	for left <= right {
		mid := (left + right) / 2
		if mid+k >= len(arr) || x-arr[mid] <= arr[mid+k]-x { // search in the left
			result = mid
			right = mid - 1
		} else if x-arr[mid] > arr[mid+k]-x { // search in the right
			left = mid + 1
		}
	}
	return arr[result : result+k]
}

// @lc code=end

