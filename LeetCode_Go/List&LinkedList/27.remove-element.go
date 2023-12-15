/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start

// Remove all occurrences of a specified value (val) from an integer array in-place. The order of elements may change. Return the count of elements remaining k. First k elements in the array contain the remaining elements.

// two pointers: curIdx, for populate the `all-val slice` from the end of the array; searchIdx, for traverse the array for then end and search for the `val`. Each time a `val` is found, swap searchIdx and curIdx, and move both to the left.

func removeElement(nums []int, val int) int {
	// edge case: len(nums) == 0
	if len(nums) == 0 {
		return 0
	}

	curIdx := len(nums) - 1
	searchIdx := len(nums) - 1
	for searchIdx >= 0 {
		// searchIdx search for next `val`
		if nums[searchIdx] != val {
			searchIdx-- //
			continue
		}

		// for i := 0; i <= curIdx; i++ {
		// 	fmt.Printf("%d ", nums[i])
		// }
		// fmt.Printf("\n")
		// fmt.Printf("searchIdx: %d, curIdx: %d \n", searchIdx, curIdx)

		// swap(nums[searchIdx], nums[curIdx])
		temp := nums[searchIdx]
		nums[searchIdx] = nums[curIdx]
		nums[curIdx] = temp

		// start next search
		curIdx--
		searchIdx--
	}
	return curIdx + 1
}

// @lc code=end

