/*
 * @lc app=leetcode id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 */

// @lc code=start

// https://www.youtube.com/watch?v=DfljaUwZsOk&list=PLot-Xpze53leOBgcVsJBEGrHPd_7x_koV&index=6

// You are given an array of integers nums, there is a `sliding window of size k` which is moving from the very left of the array to the very right, one position each time. Return the maximum in the sliding window in each positions.

// given a window and its maximum `max`. When the window slide the window, the new max can be known when expand to the right ((preMax, window[right])), but cannot be known when shrink in the left: if the left is not the max, the new max remains, if the left is the max, the new max may or may not change.

// use `monotonic queue` as the window: `monotonic queue` is a queue than ensure that all elements in the queue is monotonically increasing/decreasing. Capable to pop maximum element in queue in O(1), as the queue[0] is always the largest.
// assume we push at the end, pop at the beginning. The queue is monotonically decreasing (queue[0] is the largest)
// push(n): (expand window) delete all elements that is smaller than n in the queue
// pop(n): (shrink window) we only care about whether the number to be deleted is the max, if so, n == queue[0], remove queue[0]. Otherwise do nothing.
// popMax(): get the max queue[0] in O(1)

// O(n), since every element will only be added to the `monotonic queue` once (when call push()), and be deleted from the queue once (either when call push() or pop())

type MonotonicQueue struct {
	maxq []int
}

func (q *MonotonicQueue) push(n int) {
	i := len(q.maxq) - 1
	for i >= 0 && q.maxq[i] < n {
		i--
	}
	q.maxq = q.maxq[:i+1]
	q.maxq = append(q.maxq, n)
}
func (q *MonotonicQueue) pop(n int) {
	if n == q.maxq[0] {
		q.maxq = q.maxq[1:]
	}
}
func (q *MonotonicQueue) popMax() int {
	return q.maxq[0]
}

func maxSlidingWindow(nums []int, k int) []int {
	window := MonotonicQueue{}
	var result []int

	// initial window
	for i := 0; i < k-1; i++ {
		window.push(nums[i])
	}

	// slide window start
	left, right := 0, k-1
	for right < len(nums) {
		// expand right
		window.push(nums[right])
		right++
		// save max of the monotonic queue
		result = append(result, window.popMax())
		// shrink left
		window.pop(nums[left])
		left++
	}
	return result

}

// @lc code=end

