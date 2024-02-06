/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
 */

// @lc code=start

// You are given an array of non-overlapping intervals `intervals`. intervals is sorted in ascending order by start_i. Insert a 'newInterval' into intervals such that intervals is still sorted in ascending order by start_i and intervals still does not have any overlapping intervals (merge overlapping intervals if necessary). (ps. same number means overlap, eg. [1,2]and [2,3])

// Scan intervals, 1. determine interval that are affected, draw possible relations to find the rule: start <= end_new && end >= start_new. 2. identify the firstInterval and the lastInterval that were affected. 3. new interval is [min(firstInterval.start, firstInterval.start), max(lastInterval.end, newInterval.end)] 4. replace affected intervals with the new interval
// note: It's not required to do it in place. So it's more convenient to implement using another arrry

// dried out code
func insert(intervals [][]int, newInterval []int) [][]int {
	var ans [][]int

	for i := 0; i < len(intervals); i++ {
		// non-affected, the left part
		if intervals[i][1] < newInterval[0] {
			ans = append(ans, intervals[i])
		}
		// affected, update the newInterval (merge)
		if intervals[i][0] <= newInterval[1] && intervals[i][1] >= newInterval[0] {
			newInterval[0] = min(intervals[i][0], newInterval[0])
			newInterval[1] = max(intervals[i][1], newInterval[1])
		}

		// non-affected, the right part, if exists, append and return
		if intervals[i][0] > newInterval[1] {
			ans = append(ans, newInterval)
			ans = append(ans, intervals[i:]...)
			return ans
		}
	}

	// if there is no non-affected right part, append newInterval and return
	return append(ans, newInterval)
}

// original
// func insert(intervals [][]int, newInterval []int) [][]int {
// 	var ans [][]int

// 	// memorize first and last affected intervals
// 	first, last := -1, -1
// 	// memorize first non affected interval after newInterval
// 	firstNonAffected := -1

// 	for i := 0; i < len(intervals); i++ {
// 		// non-affected, left part
// 		if intervals[i][1] < newInterval[0] {
// 			ans = append(ans, intervals[i])
// 			continue
// 		}

// 		// affected,
// 		if intervals[i][0] <= newInterval[1] && intervals[i][1] >= newInterval[0] {
// 			if first == -1 {
// 				first = i
// 			}
// 			firstNonAffected = i + 1
// 			continue
// 		}

// 		// non-affected, right part
// 		if intervals[i][0] > newInterval[1] {
// 			firstNonAffected = i
// 			break
// 		}
// 	}

// 	// there are one or more affected intervals, merge them
// 	if first != -1 {
// 		last = firstNonAffected - 1
// 		ans = append(ans, []int{
// 			min(intervals[first][0], newInterval[0]),
// 			max(intervals[last][1], newInterval[1]),
// 		})
// 		ans = append(ans, intervals[firstNonAffected:]...)
// 	} else { // there is no affected intervals
// 		ans = append(ans, newInterval)
// 		if firstNonAffected != -1 {
// 			ans = append(ans, intervals[firstNonAffected:]...)
// 		}
// 	}

// 	return ans
// }

// @lc code=end

