/*
 * @lc app=leetcode id=1011 lang=golang
 *
 * [1011] Capacity To Ship Packages Within D Days
 */

// @lc code=start

// we have to ship all packages with weights in array `weights` within `days` days in the order of `weights`. We can decide a fix `capacity` of the ship, the ship can ship at most `capacity` weights. Return the `least capacity` to ship all packages with days.

// given an array of weights, capacity and days are strictly negative correlated and can computed each other with a function. So we can use binary search. Capacity range is [max(weights), sum(weights)], capacity max(weights) ensures that all packages can be shipped, capacity sum(weights) can ship everything in 1 day. We want to find the smallest capacity to make ship days <= days.

func shipWithinDays(weights []int, days int) int {
	minCapacity, maxCapacity := 0, 0
	for _, pak := range weights {
		minCapacity = max(minCapacity, pak)
		maxCapacity += pak
	}
	// start binary search
	result := maxCapacity
	for minCapacity <= maxCapacity {
		mid := (minCapacity + maxCapacity) / 2
		estDays := finishDays(weights, mid)
		if estDays > days { // too much time, increase min capacity
			minCapacity = mid + 1
		} else if estDays <= days { // time left, has room for capacity reduce
			result = mid
			maxCapacity = mid - 1
		}
	}
	return result
}

// given a capacity and a weights, find out how many days it need
func finishDays(weights []int, capacity int) int {
	days := 0
	weightSum := 0
	for i := 0; i < len(weights); i++ {
		weightSum += weights[i]
		if i+1 >= len(weights) || weightSum+weights[i+1] > capacity {
			days++
			weightSum = 0
		}
	}
	return days
}

// @lc code=end

