/*
 * @lc app=leetcode id=441 lang=golang
 *
 * [441] Arranging Coins
 */

// @lc code=start

// You have n coins and you want to build a staircase with these coins. A complete staircase consists of k rows where the ith row has exactly i coins, in a non-complete staircase the last row may be incomplete. Given the integer n, return the number of complete rows of the staircase you will build.

// Binary search. given a complete staircase of k rows, we can compute the number of coins needed in O(1). coinRequired(k) -> coin: (k*k-k)/2 + k. the function is monotonic increasing regarding k. Do binary search in the k domain [1, n]. When coinRequired(mid) > n, the target k must be smaller than mid. When coinRequired(mid) < n, the target k must be greater than mid, or equals the mid.

func arrangeCoins(n int) int {
	minRow, maxRow := 1, n
	result := -1
	for minRow <= maxRow {
		mid := (minRow + maxRow) / 2
		fullCoins := (mid*mid-mid)/2 + mid
		if fullCoins <= n { // has room to add more stairs
			result = mid
			minRow = mid + 1
		} else if fullCoins > n { // coin is not sufficient, reduce stairs
			maxRow = mid - 1
		}
	}
	return result
}

// @lc code=end

