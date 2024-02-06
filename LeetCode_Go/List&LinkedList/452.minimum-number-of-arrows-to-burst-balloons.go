/*
 * @lc app=leetcode id=452 lang=golang
 *
 * [452] Minimum Number of Arrows to Burst Balloons
 */

// @lc code=start

// Given an array of intervals 'points' as balloon, you want to burst all balloon with minimum number of arrows.  [find the maximum number of non-overlapping intervals (such that you have to use 1 arrow to burst each), (you may eliminate any overlapping intervals)]

// (435.non-overlapping-intervals) sort by end, then for each end position eliminate all overlapping intervals. When enter a new non-overlapping end position, add count by one.

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	count := 1
	end := points[0][1]
	for i := 1; i < len(points); i++ {
		// eliminate overlapping intervals
		if points[i][0] <= end {
			continue
		} else {
			count++
			end = points[i][1]
		}
	}
	return count
}

// @lc code=end

