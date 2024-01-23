/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start

// Given an time series array `prices` where prices[i] is the price of a stock on the ith day. Maximize your profit by choosing a day to buy in and another day in the future to sell out. Return the maximum profit.

// brute force: before buying, each day we need to decide {buy, do nothing}, after buying, each day we need to decide {sell, do nothing}. O(2^n)

// use sliding window to find a pair (left, right) where left is the minimum and right is the maximum

func maxProfit(prices []int) int {
	maxProfit := 0
	left := 0
	for right := 0; right < len(prices); right++ {
		// fmt.Printf("left: %v, right: %v, profit:%v \n", prices[left], prices[right], prices[right]-prices[left])

		if prices[left] < prices[right] {
			maxProfit = max(maxProfit, prices[right]-prices[left])
		}
		if prices[left] >= prices[right] {
			left = right
		}
	}
	return maxProfit
}

// @lc code=end

