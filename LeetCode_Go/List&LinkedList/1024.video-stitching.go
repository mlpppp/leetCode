/*
 * @lc app=leetcode id=1024 lang=golang
 *
 * [1024] Video Stitching
 */

// @lc code=start

// Given an array of intervals 'clips' and an integer 'time'. Return the minimum number of clips needed so that we can rearrange them to cover [0, time]. (If there exist overlaps, intersection will  be eliminated automatically). If the task is impossible, return -1.

// sort clips by start time ascend, and then by end time descend. Start with clips[0]. Maintain the currentEnd time. Each time we choose a newClip whose newClip.Start <= currentEnd and newClip.End is maximized (find argmax(newClip.End)).

func videoStitching(clips [][]int, time int) int {
	sort.Slice(clips, func(i, j int) bool {
		if clips[i][0] == clips[j][0] {
			return clips[i][1] > clips[j][1]
		} else {
			return clips[i][0] < clips[j][0]
		}
	})

	currentEnd := 0
	ans := 0
	bestNewClipEnd := 0 // maximize newClip.End for the newClip, and for the "impossible, return -1" cases

	for i := 0; i < len(clips); i++ {
		// enter disjoint clip
		if clips[i][0] > currentEnd {
			// if after adding the previous newClip already safisfy our need, early return
			if bestNewClipEnd >= time {
				return ans + 1
			}
			// newClip won't connect to the clips[i]
			if bestNewClipEnd < clips[i][0] {
				return -1
			} else { // newClip connects, update and start finding the next newClip
				ans++
				currentEnd = bestNewClipEnd
				bestNewClipEnd = max(bestNewClipEnd, clips[i][1])
			}
		} else { // find argmax(newClip.End) within currentEnd
			bestNewClipEnd = max(bestNewClipEnd, clips[i][1])
		}
	}
	// Deal with the last newClip: iteration may exit while searching for the next newClip. If so, bestNewClipEnd is the farthest we can get.
	if bestNewClipEnd >= time {
		return ans + 1
	}
	// clips are too short to meet the time
	return -1
}

// @lc code=end

