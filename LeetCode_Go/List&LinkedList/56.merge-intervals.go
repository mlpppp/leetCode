/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start

// Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

// sort by start, then iterate over intervals while merge. rule to merge: if interval[i].start < mergeInterval.end  ==> merge: mergeInterval = [mergeInterval.start, max(mergeInterval.end, interval[i])]

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := [][]int{}
	merge := intervals[0]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= merge[1] {
			merge[1] = max(merge[1], intervals[i][1])
		} else { // enter another merge block, save last and update 'merge'
			ans = append(ans, merge)
			merge = intervals[i]
		}
	}
	// append the last merge block
	ans = append(ans, merge)
	return ans
}

// @lc code=end

