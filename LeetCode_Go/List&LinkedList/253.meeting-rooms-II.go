/**
 * Definition of Interval:
 * type Interval struct {
 *     Start, End int
 * }
 */

/**
 * @param intervals: an array of meeting time intervals
 * @return: the minimum number of conference rooms required
 */

// Given an array of meeting time intervals consisting of start and end times [[s1,e1],[s2,e2],...] (si < ei), find the minimum number of conference rooms required.)

// 输入: intervals = [(0,30),(5,10),(15,20)]
// 输出: 2

// (trick, two pointers) find the maximum number of overlapping intervals at any point in time. Use two array to store sorted start and end times. Two pointers iterates two arrays 'by value': next = argmin(startArray[i+1], endArray[j+1]), prefer endArray if startArray[i+1] == endArray[j+1] (Because if end time and start time overlap, we have to end a meeting before enter a new meeting.) Maintain a counter: countMeet, if it is next is start array, countMeet++, else countMeet--.

func MinMeetingRooms(intervals []*Interval) int {
	startArray := make([]int, len(intervals))
	endArray := make([]int, len(intervals))
	for i, interval := range intervals {
		startArray[i] = interval.Start
		endArray[i] = interval.End
	}
	sort.Slice(startArray, func(i, j int) bool {
		return startArray[i] < startArray[j]
	})
	sort.Slice(endArray, func(i, j int) bool {
		return endArray[i] < endArray[j]
	})

	i, j := 0, 0
	countMeet, ans := 0, 0
	// we only care about maximum of countMeet, since only start can increase count countMeet, we can terminate when start array is finished
	for i < len(intervals) {
		if endArray[j] <= startArray[i] { // tie, prefer end array
			countMeet--
			j++
		}
		if startArray[i] < endArray[j] {
			countMeet++
			ans = max(ans, countMeet)
			i++
		}
	}
	return ans
}

func max(a, b int) int {
	if a <= b {
		return b
	} else {
		return a
	}
}




