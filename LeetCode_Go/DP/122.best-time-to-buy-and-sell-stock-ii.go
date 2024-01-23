/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 */

// @lc code=start

// Given an time series array `prices` where prices[i] is the price of a stock on the ith day. Maximize your profit by buying and selling stock. You can buy and sell unlimited times, but you can only hold at most one share of the stock at any time. Return the maximum profit.

// brute force: for each day, we need to decide whether to {buy, do nothing, sell}. O(3^n)

// Analysis: StateSet: {isHolding, day(prices)}. ActionsSet: {buy, do nothing, sell}

// subproblems: maxProfitAtDayI(isholding bool, prices[i:]) int. maximum profit I can get, if I hold stock or not the at the ith day. Targets: maxProfitAtDayI(not isholding, prices[0:])

// transition: maxProfitAtDayI(isholding, prices[i:]) = {
// 	if not isHolding:
// 		maxProfitAtDayI(not isholding, prices[i+1:]), do nothing
// 		maxProfitAtDayI(isholding, prices[i+1:]) - prices[i], buy in today
// 	if isHolding:
// 		maxProfitAtDayI(isholding, prices[i+1:]), do nothing
// 		maxProfitAtDayI(not isholding, prices[i+1:]) + prices[i], sell out today
// }

// base case: day out of range, return profit 0

// func maxProfit(prices []int) int {
// 	// DP table [isholding][prices], init with -1
// 	table := make([][]int, 2)
// 	for i, _ := range table {
// 		table[i] = make([]int, len(prices))
// 		for j, _ := range table[i] {
// 			table[i][j] = -1
// 		}
// 	}
// 	return maxProfitAtDayI(prices, 0, 0, table)
// }

// func maxProfitAtDayI(prices []int, isHolding, day int, table [][]int) int {
// 	// base case: day out of range
// 	if day >= len(prices) { //
// 		return 0
// 	}
// 	// query DP table
// 	if table[isHolding][day] != -1 {
// 		return table[isHolding][day]
// 	}
// 	// recursion
// 	if isHolding == 1 { // 1 == true
// 		table[1][day] = max(
// 			maxProfitAtDayI(prices, 1, day+1, table),
// 			maxProfitAtDayI(prices, 0, day+1, table)+prices[day])
// 	} else if isHolding == 0 { //  0 == false
// 		table[0][day] = max(
// 			maxProfitAtDayI(prices, 0, day+1, table),
// 			maxProfitAtDayI(prices, 1, day+1, table)-prices[day])
// 	}
// 	return table[isHolding][day]
// }

// iterative solution: from right to left, update two row at a time
func maxProfit(prices []int) int {
	// DP table [isholding][prices]
	table := make([][]int, 2)
	for i, _ := range table {
		table[i] = make([]int, len(prices))
	}
	// base case: last day returns
	table[0][len(prices)-1] = 0
	table[1][len(prices)-1] = prices[len(prices)-1]

	for day := len(prices) - 2; day >= 0; day-- {
		table[0][day] = max(
			table[0][day+1],
			table[1][day+1]-prices[day],
		)
		table[1][day] = max(
			table[1][day+1],
			table[0][day+1]+prices[day],
		)
	}
	return table[0][0]
}

// @lc code=end

