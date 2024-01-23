/*
 * @lc app=leetcode id=268 lang=golang
 *
 * [268] Missing Number
 */

// @lc code=start

// Given an array nums containing n distinct numbers in the range [0, n], unsorted, return the only number in the range that is missing from the array. (eg. [9,6,4,2,3,5,7,0,1] => 8)

// Brute force: scan and save every number as index in hashmap/hashset, then for every number in the range check if the key exists.

// difference of sum: for nums of length n, the result is `sum(0, n) - sum(nums)`. sum(0, n) = (n(n+1))/2
// func missingNumber(nums []int) int {
// 	sumSeq := ((len(nums) + 1) * len(nums)) / 2
// 	sum := 0
// 	for _, num := range nums {
// 		sum += num
// 	}
// 	return sumSeq - sum
// }

// XOR: by concatenating nums and the array [0...n] in a row, the problem becomes 136.single-number. XOR eliminate repeated number pairs in an array, no matter how the numbers are ordered!!!
func missingNumber(nums []int) int {
	ans := 0
	// the indices i make up of the array [0...n-1]
	for i, _ := range nums {
		ans ^= nums[i] ^ i
	}
	// XOR the last number n of array [0...n]
	ans ^= len(nums)
	return ans
}

// @lc code=end

