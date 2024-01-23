/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 */

// @lc code=start

// Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that: a, b, c, and d are distinct, and nums[a] + nums[b] + nums[c] + nums[d] == target.

// subproblem: nSum(nums[i:], target, n) [][]int: find an array of unique n numbers in nums such that add to target. Implement 3Sum logic while n > 2, and implement 2Sum logic when n == 2.

func fourSum(nums []int, target int) [][]int {
	quickSort(nums, 0, len(nums)-1)
	return nSum(nums, 0, target, 4)
}

func nSum(nums []int, i, target, n int) [][]int {
	var result [][]int
	// base case: n or i out of range
	if i >= len(nums)-1 || n < 2 {
		return result
	}
	// base case: n == 2
	if n == 2 {
		j := len(nums) - 1
		for i < j {
			sum := nums[i] + nums[j]
			if sum < target {
				i++
			} else if sum > target {
				j--
			}
			if sum == target {
				result = append(result, []int{nums[i], nums[j]})
				// skip repeats of nums[i], nums[j]
				iVal, jVal := nums[i], nums[j]
				for i < j && nums[i] == iVal {
					i++
				}
				for i < j && nums[j] == jVal {
					j--
				}
			}
		}
	} else { // recursion
		for i < len(nums) {
			subResult := nSum(nums, i+1, target-nums[i], n-1)
			for idx := range subResult {
				subResult[idx] = append(subResult[idx], nums[i])
			}
			result = append(result, subResult...)
			// skip repeats of nums[i]
			val := nums[i]
			for i < len(nums) && nums[i] == val {
				i++
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

