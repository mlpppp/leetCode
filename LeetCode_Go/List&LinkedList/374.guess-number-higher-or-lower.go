/*
 * @lc app=leetcode id=374 lang=golang
 *
 * [374] Guess Number Higher or Lower
 */

// @lc code=start
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

// play the Guess Game as the guesser. Given a `pick` number in range [1, n], use a guess(int num)int API to find out the pick.

// plain binary search

func guessNumber(n int) int {
	left, right := 1, n
	result := -1
	for left <= right {
		mid := (left + right) / 2
		if guess(mid) == 1 { // guess is lower than the pick
			left = mid + 1
		} else if guess(mid) == -1 { // guess is greater than the pick
			right = mid - 1
		} else {
			result = mid
			break
		}
	}
	return result
}

// @lc code=end

