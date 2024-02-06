/*
 * @lc app=leetcode id=986 lang=golang
 *
 * [986] Interval List Intersections
 */

// @lc code=start

// You are given two lists of closed intervals, firstList and secondList, Each list of intervals is pairwise disjoint (no overlap) and in sorted order. Return the intersection of these two interval lists. (For example, the intersection of [1, 3] and [2, 4] is [2, 3].)

// intersect of two overlapping intervals: [max(a[start], b[start]), min(a[end], b[end])]. Use two pointers to find overlapping intervals in two arrays, and save intersects to result according to the rule. In each iteration we only move one pointer: pointer to the interval that end first.

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	pA, pB := 0, 0
	ans := [][]int{}
	for pA < len(firstList) && pB < len(secondList) {
		// if overlap, cut the intersect and save to answer
		// (overlap is mutual)
		if firstList[pA][0] <= secondList[pB][1] && firstList[pA][1] >= secondList[pB][0] {
			ans = append(ans, []int{
				max(firstList[pA][0],
					secondList[pB][0]),
				min(firstList[pA][1],
					secondList[pB][1])})
		}
		// move the point that end earlier
		if firstList[pA][1] >= secondList[pB][1] {
			pB++
		} else {
			pA++
		}
	}
	return ans
}

// @lc code=end

