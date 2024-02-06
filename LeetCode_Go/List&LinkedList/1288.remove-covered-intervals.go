/*
 * @lc app=leetcode id=1288 lang=golang
 *
 * [1288] Remove Covered Intervals
 */

// @lc code=start

// Given an array intervals where intervals[i] = [l_i, r_i] represent the interval [l_i, r_i), remove all intervals that are covered by another interval in the list. The interval [a, b) is covered by the interval [c, d) if and only if c <= a and b <= d. Return the number of remaining intervals.

// Sort then Iterate intervals, counting covers while update the cover boundary. Cover rule: coverBy[end] <= cover[end] && coverBy[start] >= cover[start]. For sorted array there are three possible cases:
// 1. intervals[i][start] == cover[start], where we update cover[end] == max(cover[end], intervals[i][end]),
// else 1. intervals[i][end] <= cover[end]: interval[i] should be merged
// else 2. intervals[i][end] > cover[end]: that should begin another cover, add result by 1

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	count := 1
	coverL, coverR := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] == coverL {
			coverR = max(coverR, intervals[i][1])
		} else {
			if intervals[i][1] > coverR {
				coverL = intervals[i][0]
				coverR = intervals[i][1]
				count++
			}
		}
	}
	return count
}

// @lc code=end

