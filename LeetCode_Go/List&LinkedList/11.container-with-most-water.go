/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start

// You are given an integer array `height` of size n. height[i] is the height of a containers' rim at location i. Find two locations i to form a container, such that the container contains the most water. Return the maximum amount of water a container can store.

// two pointers: Find two location 'l, r' such that maximize(diff(l, r)*min(height[l], height[r])). Use two pointers from both side, each time we shrink either l or r. In each shrink, diff(l, r) decreases 1. So we want to optimize the second part: shrink in a way that min(height[l], height[r]) is optimized: argmin(height[l], height[r]), since it is the only possible way to optimize min(height[l], height[r])

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxCap := 0
	for l < r {
		maxCap = max(maxCap, (r-l)*min(height[l], height[r]))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxCap
}

// @lc code=end

