/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 */

// @lc code=start

// Given an `integer array coins` representing coins of different denominations. Return the `fewest number of coins` that you need to make up a integer `amount`. If that amount of money cannot be made up by any combination of the coins, return -1.

// subproblem: coinChange(amount) int. Given an amount, return minimum number of coins needed to make up that amount.

// Transition: coinChange(amount) = min(coinChange(amount-coins[i]) for i in coins[i]) + 1.

// DP table: result of coinChange(amount) => []int. Initalize large value

// unsolvable case: all coins are either -1 or out of range

// pseudocode: start
// coins = [1,2,5]
// amount = 11
// table = [] size of amount+1 (incl 0), initalize with maxInt
// for i := int values in amount:
//     minCoins = large value
//     for denominations in coins:
// 		if i-denominations < 0 || table[i-denominations] == -1:
// 			continue
// 		if table[i-denominations] <= minCoins:
// 			minCoins = table[i-denominations]
// 	if minCoins is not changed:
// 		table[i] = -1
// 	else:
// 		table[i] = minCoins + 1
// return table[amount]

// O(nk) n: amount, k: len(coins)

func coinChange(coins []int, amount int) int {
	table := make([]int, amount+1)
	table[0] = 0
	for i := 1; i < amount+1; i++ { // initialize table
		table[i] = amount + 1
	}

	for i := 1; i < amount+1; i++ {
		// find the best subproblem that can add one coin to satisfy the current problem
		minCoins := amount + 1
		for _, coin := range coins {
			if i-coin < 0 || table[i-coin] == -1 {
				continue
			}
			if table[i-coin] < minCoins {
				minCoins = table[i-coin]
			}
		}
		if minCoins == amount+1 {
			table[i] = -1
		} else {
			table[i] = minCoins + 1
		}
	}
	return table[amount]
}

// @lc code=end

