/*
 * @lc app=leetcode id=123 lang=golang
 *
 * [123] Best Time to Buy and Sell Stock III
 */

// @lc code=start

// Given an time series array `prices` where prices[i] is the price of a stock on the ith day. Maximize your profit by buying and selling stock. You can buy and sell 2 times, you can only hold at most one share of the stock at any time. Return the maximum profit.

// states = [isHolding, day, remainTransaction]
// choices = [buy, do nothing, sell]

// subproblems: maxProfitAtDayI(isHolding, day, remainTransaction): starting from `day`, with/without stock, with a remainTransaction times, find the maximum profit I can achieve. Goal: maxProfitAtDayI(isHolding=false, day=0, remainTransaction=2).

// transition: check remainTransaction at buy time. Decrease remainTransaction in sell time to mark a end of a transaction. Dont allow to buy when remainTransaction==0

// maxProfitAtDayI(isHolding, day, remainTransaction) = {
// 	if not isHolding {
// 		doNothing = maxProfitAtDayI(not isHolding, day+1, remainTransaction)
// 		if remainTransaction > 0 {
// 			buy = maxProfitAtDayI(isHolding, day+1, remainTransaction)-prices[i]
// 		}
// 		return max(doNothing, buy)
// 	}
// 	if isHolding {
// 		doNothing = maxProfitAtDayI(isHolding, day+1, remainTransaction)
// 		sell = maxProfitAtDayI(not isHolding, day+1, remainTransaction)+prices[i],
// 		return max(doNothing, sell)
// 	}
// }

// base case: {
// 	day out of range, you cannot buy/sell anything now, return 0,
// 	remainTransaction is used up, because you cannot do anything now, return 0
// }

// recursive solution
// func maxProfit(prices []int) int {
// 	// DP table [remainTransaction=2][isHolding=2][day], initialize with -1
// 	table := make([][][]int, 2)
// 	for i, _ := range table {
// 		table[i] = make([][]int, 2)
// 		for j, _ := range table[i] {
// 			table[i][j] = make([]int, len(prices))
// 			for k, _ := range table[i][j] {
// 				table[i][j][k] = -1
// 			}
// 		}
// 	}

// 	return maxProfitAtDayI(prices, 1, 0, 0, table)
// }

// func maxProfitAtDayI(prices []int, remainTransaction, isHolding, day int, table [][][]int) int {
// 	// base case: day out of range
// 	if day >= len(prices) {
// 		return 0
// 	}
// 	if remainTransaction < 0 {
// 		return 0
// 	}

// 	// query DP table
// 	if table[remainTransaction][isHolding][day] != -1 {
// 		return table[remainTransaction][isHolding][day]
// 	}
// 	// recursion
// 	if isHolding == 0 {
// 		// do nothing
// 		doNothing := maxProfitAtDayI(prices, remainTransaction, 0, day+1, table)
// 		// buy in
// 		buy := maxProfitAtDayI(prices, remainTransaction, 1, day+1, table) - prices[day]
// 		table[remainTransaction][isHolding][day] = max(doNothing, buy)
// 	}
// 	if isHolding == 1 {
// 		// do nothing
// 		doNothing := maxProfitAtDayI(prices, remainTransaction, 1, day+1, table)
// 		// sell out
// 		sell := maxProfitAtDayI(prices, remainTransaction-1, 0, day+1, table) + prices[day]
// 		table[remainTransaction][isHolding][day] = max(doNothing, sell)
// 	}
// 	return table[remainTransaction][isHolding][day]
// }

// [3,3,5,0,0,3,1,4]
// [[[-1 -1 4 4 4 3 3 0] [-1 -1 -1 4 4 4 4 4]] [[6 6 6 6 6 3 3 0] [-1 9 9 6 6 6 4 4]]]

// iterative solution: [remainTransaction=0][isHolding][day], then  [remainTransaction=1][isHolding][day].
func maxProfit(prices []int) int {
	// DP table [remainTransaction=2 + 1 padding][isHolding=2][day]
	table := make([][][]int, 3)
	for i, _ := range table {
		table[i] = make([][]int, 2)
		for j, _ := range table[i] {
			table[i][j] = make([]int, len(prices))
		}
	}

	// base case: last day (last column)
	for remainTransaction := 1; remainTransaction <= 2; remainTransaction++ {
		table[remainTransaction][0][len(prices)-1] = 0
		table[remainTransaction][1][len(prices)-1] = prices[len(prices)-1]
	}

	for remainTransaction := 1; remainTransaction <= 2; remainTransaction++ {
		for day := len(prices) - 2; day >= 0; day-- {
			table[remainTransaction][0][day] = max(
				table[remainTransaction][0][day+1],
				table[remainTransaction][1][day+1]-prices[day],
			)
			table[remainTransaction][1][day] = max(
				table[remainTransaction][1][day+1],
				table[remainTransaction-1][0][day+1]+prices[day],
			)
		}
	}

	return table[2][0][0]
}

// @lc code=end

