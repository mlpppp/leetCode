/*
 * @lc app=leetcode id=367 lang=golang
 *
 * [367] Valid Perfect Square
 */

// @lc code=start

// Given a positive integer num, return true if num is x^2 in which x is an integer.

// binary search. sqrt(x) -> x^2, given x in search domain [1, num], finds the x that outputs num, if cannot be found, returns false.

func isPerfectSquare(num int) bool {
	left, right := 1, num
	for left <= right {
		mid := left + (right-left)/2
		sqr := mid * mid
		if sqr < num {
			left = mid + 1
		} else if sqr > num {
			right = mid - 1
		} else if sqr == num {
			return true
		}
	}
	return false
}

// @lc code=end

