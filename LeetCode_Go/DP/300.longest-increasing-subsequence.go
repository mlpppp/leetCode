/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

// @lc code=start

// neetcode: https://www.youtube.com/watch?v=cjWnW0hdF1Y

// Given an integer array nums, return the length of the longest strictly increasing subsequence. ([10,9,2,5,3,7,101,18] => [2,3,7,101], therefore the length is 4.)

// notice that two common subsequence must start with the same character, end with the same character.

// subproblem: lengthOfLIS(nums []int, end int): length of longest LIS for a sub sequence start from 0 and end with  idx `end`. Increment n from 0 to len(nums-1) to aggregate the table, and return the maximun LIS in the table

// transition: adding an element a to the end could increate subsequence end with element b by 1, provided that (a > b). lengthOfLIS(nums, end) = 1 + max(lengthOfLIS(nums, i) for (i in range(0, end-1)) and (nums[i] < nums[end]))

// pseudocode: start
// nums = [10,9,2,5,3,7,101,18]
// table = [] of tuples (num, LIS), for num in nums and LIS initalized with 1
// for end := 0; end < len(table); end++ {
// 	maxLIS := 0
// 	for i := 0; i < end; i++ {
// 		// find the max(lengthOfLIS(nums, i) for (i in range(0, end-1)) and (nums[i] < nums[end]))
// 		if table[i][0] < table[end][0] && table[i][1] > maxLIS {
// 			maxLIS = table[i][1]
// 		}
// 		table[end][1] = maxLIS + 1
// 	}
// }
// return argmax(i, table[i][1])

// func lengthOfLIS(nums []int) int {
// 	// tuple: (num, maxLIS)
// 	table := make([][2]int, len(nums))
// 	for idx, val := range nums {
// 		table[idx] = [2]int{val, 1}
// 	}
// 	for end := 0; end < len(table); end++ {
// 		maxLIS := 0
// 		// find the max(lengthOfLIS(nums, i) for (i in range(0, end-1)) and (nums[i] < nums[end]))
// 		for i := 0; i < end; i++ {
// 			if table[i][0] < table[end][0] && table[i][1] > maxLIS {
// 				maxLIS = table[i][1]
// 			}
// 		}
// 		table[end][1] = maxLIS + 1
// 	}

// 	maxLIS := 1
// 	for _, value := range table {
// 		if value[1] > maxLIS {
// 			maxLIS = value[1]
// 		}
// 	}

// 	return maxLIS
// }

// optimized code
func lengthOfLIS(nums []int) int {
	table := make([]int, len(nums))
	for idx := 0; idx < len(nums); idx++ {
		table[idx] = 1
	}
	for end := 0; end < len(table); end++ {
		// find the max(lengthOfLIS(nums, i) for (i in range(0, end-1)) and (nums[i] < nums[end]))
		maxLIS := 0
		for i := 0; i < end; i++ {
			if nums[i] < nums[end] && table[i] > maxLIS {
				maxLIS = table[i]
			}
		}
		if maxLIS > 0 {
			table[end] = maxLIS + 1
		} else {
			table[end] = 1
		}
	}

	maxLIS := 1
	for _, value := range table {
		if value > maxLIS {
			maxLIS = value
		}
	}

	return maxLIS
}

// @lc code=end

