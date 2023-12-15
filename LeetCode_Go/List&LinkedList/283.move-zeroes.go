/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */

// @lc code=start

// Given an integer array nums, move all 0's to the end of it  in-place, while maintaining the relative order of the non-zero elements.

// two pointers: `curIdx`, for populate the `non-zero slice` from the beginning of the array; `searchIdx`, for traverse the array and find non-zero element. Each time a non-zero element is found assign it to the curIdx, and move curIdx to the right. While the bound of non-zero slice is found, fill the rest with 0.

func moveZeroes(nums []int) {
	curIdx, searchIdx := 0, 0
	for searchIdx < len(nums) {

		// fmt.Println(nums)
		// fmt.Printf("searchIdx: %d, curIdx: %d \n", searchIdx, curIdx)

		if nums[searchIdx] != 0 {
			nums[curIdx] = nums[searchIdx]
			curIdx++
		}

		searchIdx++
	}

	for i := curIdx; i < len(nums); i++ {
		nums[i] = 0
	}
}

// @lc code=end

