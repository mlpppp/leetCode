/*
 * @lc app=leetcode id=416 lang=golang
 *
 * [416] Partition Equal Subset Sum
 */

// @lc code=start
// https://www.youtube.com/watch?v=IsvocB5BJhw
// Given an integer array nums, return true if you can partition the array into two subsets such that the sum of the elements in both subsets is equal.

// If nums can be partitioned into two subsets of equal sum, the sum of each subset should be exactly sum(nums)/2. So the question becomes finding a subset in nums that sums up to sum(nums)/2.

// subproblem: canFindSum(nums[i:], sum) bool: if I can find a number "sum" in a subarray nums[i:]. target: canFindSum(nums[:], sum(nums)/2)

// transition canFindSum(nums[i:], sum) bool = {
// 	canFindSum(nums[i+1:], sum-nums[i]), add current || canFindSum(nums[i+1:], sum), do not add current
// }

// O(n*sum(nums))

// recursive solution
// func canPartition(nums []int) bool {
// 	if sum(nums)%2 != 0 { // odd sum cannot be partitioned
// 		return false
// 	}
// 	// DP table: [len(nums)][(sum(nums)/2)+1], use 0/1 to indicate bool value, initialize with -1
// 	table := make([][]int, len(nums))
// 	target := sum(nums) >> 1
// 	for i, _ := range table {
// 		table[i] = make([]int, target+1)
// 		for j, _ := range table[i] {
// 			table[i][j] = -1
// 		}
// 	}
// 	return canFindSum(nums, 0, target, table)
// }

// func canFindSum(nums []int, i, sum int, table [][]int) bool {
// 	// base case: i out of range:
// 	if i >= len(table) { // used up all nums
// 		if sum == 0 { // happens to get the sum, good solution
// 			return true
// 		} else {
// 			return false
// 		}
// 	}
// 	// base case: sum out of range
// 	if sum == 0 { // get the sum early, good solution
// 		table[i][sum] = 1
// 		return true
// 	}
// 	if sum < 0 { // added to much to the sum, non-feasible solution
// 		return false
// 	}

// 	// use DP table
// 	if table[i][sum] != -1 {
// 		if table[i][sum] == 0 {
// 			return false
// 		} else {
// 			return true
// 		}
// 	}
// 	// recursive transition
// 	// fmt.Printf("%v is calling canFindSum(%v, %v) and canFindSum(%v, %v)\n", nums[i:], nums[i+1:], sum-nums[i], nums[i+1:], sum)
// 	result := canFindSum(nums, i+1, sum-nums[i], table) || canFindSum(nums, i+1, sum, table)
// 	if result {
// 		table[i][sum] = 1
// 	} else {
// 		table[i][sum] = 0
// 	}
// 	return result
// }

// func sum(nums []int) int {
// 	sum := 0
// 	for i := range nums {
// 		sum += nums[i]
// 	}
// 	return sum
// }

// iterative: bottom to top, left to right
// DP[i][sum-nums[i]] <- DP[i][sum]
//
//	|
//	DP[i+1][sum]
func canPartition(nums []int) bool {
	if sum(nums)%2 != 0 { // odd sum cannot be partitioned
		return false
	}
	// DP table: [len(nums)+1][(sum(nums)/2)+1], column 0 (sum=0) set to true
	table := make([][]bool, len(nums)+1)
	target := sum(nums) >> 1
	for i, _ := range table {
		table[i] = make([]bool, target+1)
		table[i][0] = true
	}
	for i := len(nums) - 1; i >= 0; i-- { // from the last 2nd row to first row
		for sum := 1; sum <= target; sum++ { // from the second column to last column
			if sum-nums[i] < 0 { // the current value cannot be added to the sum
				table[i][sum] = table[i+1][sum]
			} else {
				table[i][sum] = table[i+1][sum] || table[i+1][sum-nums[i]]
			}
		}
	}
	return table[0][target]
}
func sum(nums []int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	return sum
}

// @lc code=end

