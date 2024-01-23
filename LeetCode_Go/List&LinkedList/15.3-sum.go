/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start

// Given an integer array nums, return all the non-duplicate triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

// Sort the input array. Implement `2Sum(nums, n)` to find unique twoSums that sums to `n` in a sorted array. Iterate the array and for each `first encounter of a unique element` nums[i] call 2Sum(nums[i+1:], 0-nums[i])

func threeSum(nums []int) [][]int {
	quickSort(nums, 0, len(nums)-1)
	var result [][]int
	i := 0
	for i < len(nums) {
		// find twoSums
		twoSumResult := twoSum(nums, i+1, 0-nums[i])
		if len(twoSumResult) > 0 {
			for idx := range twoSumResult {
				twoSumResult[idx] = append(twoSumResult[idx], nums[i])
			}
			result = append(result, twoSumResult...)
		}
		// skip all repeated values
		val := nums[i]
		for i < len(nums) && nums[i] == val {
			i++
		}
	}
	return result
}

// find unique pairs that sums to number n in sorted array nums[i:]
func twoSum(nums []int, i, n int) [][]int {
	var result [][]int
	i, j := i, len(nums)-1
	for i < j {
		sum := nums[i] + nums[j]
		if sum < n {
			i++
		} else if sum > n {
			j--
		}
		if sum == n {
			result = append(result, []int{nums[i], nums[j]})
			// skip all repeated elements
			leftVal, rightVal := nums[i], nums[j]
			for i < j && nums[i] == leftVal {
				i++
			}
			for i < j && nums[j] == rightVal {
				j--
			}
		}
	}
	return result
}

func quickSort(nums []int, i, j int) {
	// base case: empty or single element array
	if i == j || i > j {
		return
	}

	// pre recursion
	target := nums[i]
	pSmall := i + 1
	pEdit := i + 1
	for pSmall <= j {
		// if find one smaller than target, sawp pEdit and pSmall, and move pEdit.
		if nums[pSmall] < target {
			nums[pSmall], nums[pEdit] = nums[pEdit], nums[pSmall]
			pEdit++
		}
		pSmall++
	}

	// swap pEdit - 1 and target
	nums[i], nums[pEdit-1] = nums[pEdit-1], nums[i]

	// recursion
	quickSort(nums, i, pEdit-2)
	quickSort(nums, pEdit, j)
}

// @lc code=end

