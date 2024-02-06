/**
 * Definition of Interval:
 * type Interval struct {
 *     Start, End int
 * }
 */

/**
 * @param intervals: an array of meeting time intervals
 * @return: if a person could attend all meetings
 */

// Given an array of meeting time intervals consisting of start and end times [[s1,e1],[s2,e2],...] (si < ei), determine if a person could attend all meetings.

// sort, scan and determine if there is any overlapping interval

import (
	"sort"
)

func CanAttendMeetings(intervals []*Interval) bool {
	if len(intervals) == 0 {
		return true
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	lastInterval := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start < lastInterval.End {
			return false
		}
		lastInterval = intervals[i]
	}
	return true
}