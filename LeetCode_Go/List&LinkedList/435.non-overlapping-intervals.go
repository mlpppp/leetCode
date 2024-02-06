/*
 * @lc app=leetcode id=435 lang=golang
 *
 * [435] Non-overlapping Intervals
 */

// @lc code=start

// Given an array of right-open intervals intervals where intervals[i] = [starti, endi], return the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.

// trick, 死记.  Sort by the end_i (ascend), then for each end position remove all intervals that overlap with it.

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	count := 0
	curEnd := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < curEnd { // overlap with current interval
			count++
		} else { // the next interval that is kept
			curEnd = intervals[i][1]
		}
	}
	return count
}

// @lc code=end

