/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start

// Remove duplicates in a sorted integer array in-place. Return the number of unique elements (k) and modify the array to include only the first k unique elements in sorted order.

// two pointers: curIdx, for populate the unique part; searchIdx, for traverse the array and search for next unique element. One state: curValue, store value of the last unique element, as a criterion for finding a next unique element.

// O(n)
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 { // len 1 case
		return len(nums)
	}
	// len > 1 case
	curIdx, searchIdx := 1, 1
	curValue := nums[0]
	for searchIdx < len(nums) {
		if nums[searchIdx] <= curValue {
			searchIdx++
			continue
		}
		curValue = nums[searchIdx]
		nums[curIdx] = nums[searchIdx]
		curIdx++
	}
	return curIdx
}

// @lc code=end

