/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start

// Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to target.  Assume that each input would have exactly one solution. Elements in the array may not be unique.

// brute force: for each element scan the array from it once. O(n^2)

// General solution: sort the array (O(nlogn)), then two pointers from both end, if sum larger then target, shrink the right, else shrink the left. Since the question requests indices, create tuples that stores values and index.

// type Tuple struct {
// 	Num int
// 	Idx int
// }

// func twoSum(nums []int, target int) []int {
// 	numTuples := make([]*Tuple, len(nums))
// 	for idx, n := range nums {
// 		numTuples[idx] = &Tuple{n, idx}
// 	}
// 	quickSort(numTuples, 0, len(numTuples)-1)

// 	result := make([]int, 2)
// 	pLeft, pRight := 0, len(numTuples)-1
// 	for pLeft <= pRight {
// 		if numTuples[pLeft].Num+numTuples[pRight].Num < target {
// 			pLeft++
// 		} else if numTuples[pLeft].Num+numTuples[pRight].Num > target {
// 			pRight--
// 		} else {
// 			result[0] = numTuples[pLeft].Idx
// 			result[1] = numTuples[pRight].Idx
// 			break
// 		}
// 	}
// 	return result
// }

// func quickSort(nums []*Tuple, i, j int) {
// 	// base case: empty or single element array
// 	if i == j || i > j {
// 		return
// 	}

// 	// pre recursion
// 	target := nums[i].Num
// 	pSmall := i + 1
// 	pEdit := i + 1
// 	for pSmall <= j {
// 		// if find one smaller than target, sawp pEdit and pSmall, and move pEdit.
// 		if nums[pSmall].Num < target {
// 			nums[pSmall], nums[pEdit] = nums[pEdit], nums[pSmall]
// 			pEdit++
// 		}
// 		pSmall++
// 	}

// 	// swap pEdit - 1 and target
// 	nums[i], nums[pEdit-1] = nums[pEdit-1], nums[i]

// 	// recursion
// 	quickSort(nums, i, pEdit-2)
// 	quickSort(nums, pEdit, j)
// }

// hashmap solution : make a hashmap {num: idx} for each element in the array. Scan the array once and check if target-nums[i] exists in the hashmap. Add the element to the hashmap only after it is scaned.
func twoSum(nums []int, target int) []int {
	remainder := make(map[int]int, len(nums))
	result := make([]int, 2)
	for i, val := range nums {
		if idx, exists := remainder[target-val]; exists {
			result[0] = i
			result[1] = idx
			break
		}
		remainder[val] = i
	}
	return result
}

// @lc code=end

