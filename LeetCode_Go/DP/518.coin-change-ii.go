/*
 * @lc app=leetcode id=518 lang=golang
 *
 * [518] Coin Change II
 */

// @lc code=start

// Given an integer array coins representing different coins, and an integer amount representing a total amount of money. Return the number of combinations that make up that amount. If that amount of money cannot be made up by any combination of the coins, return 0.

// amount = 5, coins = [1,2,5]
// combinations = [5, 2+2+1, 2+1+1+1, 1+1+1+1+1]

// subproblem: waysOfChange(coins[i:], amount) int: how many ways to exchange a `unlimited source of a subarray of coins coins[i:]` sum up to `amount`. target: waysOfChange(coins[:], amount)

// transition: waysOfChange(coins[i:], amount) = {
// 	if coins[i] < amount, we can add it (unlimited times), or just dont add it:
// 		waysOfChange(coins[i:], amount-coins[i]) + waysOfChange(coins[i+1:], amount)
// 	if coins[i] > amount, we cannot add it:
// 		waysOfChange(coins[i+1:], amount)
// }

// base case: waysOfChange(coins[i:], amount) = {
// 	1, if amount == 0 (found a way to sum up to the amount)
// 	0, if i out of range (no more coins available)
// }

// recursive solution
// func change(amount int, coins []int) int {
// 	//dp table: [len(coins)][amount+1], first column for `amount == 0`. Init ALL with -1
// 	table := make([][]int, len(coins))
// 	for i, _ := range table {
// 		table[i] = make([]int, amount+1)
// 		for j, _ := range table[i] {
// 			table[i][j] = -1
// 		}
// 	}
// 	sol := waysOfChange(coins, 0, amount, table)
// 	return sol
// }

// func waysOfChange(coins []int, i, amount int, table [][]int) int {
// 	// base case
// 	if amount == 0 { // found a way to change
// 		table[i][amount] = 1
// 		return 1
// 	}
// 	if i >= len(coins) { // used all coins but amount is still non-zero
// 		return 0
// 	}
// 	// query DP table
// 	if table[i][amount] != -1 {
// 		return table[i][amount]
// 	}
// 	// recursion
// 	if coins[i] > amount { // skip current coin
// 		table[i][amount] = waysOfChange(coins, i+1, amount, table)
// 	}
// 	if coins[i] <= amount { // either add current coin unlimited times, or do not add it
// 		table[i][amount] = waysOfChange(coins, i, amount-coins[i], table) +
// 			waysOfChange(coins, i+1, amount, table)
// 	}
// 	return table[i][amount]
// }

// iterative solution: [len(coins)+1][amount+1], first column for `amount == 0`, init with 1. Last row for padding. From bottom to top, left to right

// table[i][amount-coins[i]]  <-  table[i][amount]
// 						           |
// 							      table[i+1][amount]

func change(amount int, coins []int) int {
	//dp table: [len(coins)+1][amount+1], first column for `amount == 0`, init with 1. last row for padding
	table := make([][]int, len(coins)+1)
	for i, _ := range table {
		table[i] = make([]int, amount+1)
		table[i][0] = 1
	}
	for i := len(coins) - 1; i >= 0; i-- {
		for amt := 0; amt <= amount; amt++ {
			if coins[i] > amt {
				table[i][amt] = table[i+1][amt]
			}
			if coins[i] <= amt {
				table[i][amt] = table[i][amt-coins[i]] + table[i+1][amt]
			}
		}
	}
	return table[0][amount]
}

// @lc code=end

