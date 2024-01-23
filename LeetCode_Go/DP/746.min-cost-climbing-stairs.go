/*
 * @lc app=leetcode id=746 lang=golang
 *
 * [746] Min Cost Climbing Stairs
 */

// @lc code=start

// You are given an integer array cost where cost[i] is the cost of ith step on a staircase. Once you pay the cost, you can either climb one or two steps. You can either start from the step 0, or step 1. Return the minimum cost to reach the top of the floor.

// subproblem: minCostClimbingFrom(i) int starting from stair i, the min cost to reach the top of the floor. Target: min(minCostClimbingFrom(0), minCostClimbingFrom(1))

// transition: minCostClimbingFrom(i) = {
// 	cost[i] + min(minCostClimbingFrom(i+1), minCostClimbingFrom(i+2))
// }

// base case: if i out of bounds, already reached the top, cost is 0

// recursive solution
// func minCostClimbingStairs(cost []int) int {
// 	// init with -1
// 	table := make([]int, len(cost))
// 	for i, _ := range table {
// 		table[i] = -1
// 	}
// 	return min(minCostClimbingFrom(0, cost, table), minCostClimbingFrom(1, cost, table))
// }
// func minCostClimbingFrom(i int, cost, table []int) int {
// 	// base case:
// 	if i >= len(cost) {
// 		return 0
// 	}
// 	// query table
// 	if table[i] != -1 {
// 		return table[i]
// 	}
// 	// resursion DFS
// 	table[i] = cost[i] + min(minCostClimbingFrom(i+1, cost, table), minCostClimbingFrom(i+2, cost, table))
// 	return table[i]
// }
// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	} else {
// 		return b
// 	}
// }

// iterative solution: from right to left
func minCostClimbingStairs(cost []int) int {
	// init value is 0.
	table := make([]int, len(cost))
	// base case: last two steps
	table[len(table)-1] = cost[len(table)-1]
	table[len(table)-2] = cost[len(table)-2]

	for i := len(table) - 3; i >= 0; i-- {
		table[i] = cost[i] + min(table[i+1], table[i+2])
	}
	return min(table[0], table[1])
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end

