/*
 * @lc app=leetcode id=875 lang=golang
 *
 * [875] Koko Eating Bananas
 */

// @lc code=start

// Given an array `piles` of size n and an integer `h`. There are n piles of bananas, each pile has piles[i] bananas. Koko need to eat all bananas within `h` hours. Minimize koko's bananas-per-hour eating speed `k`: Each hour, koko chooses one pile of bananas and eats k bananas, if the pile (or ramain of the pile) has less than k bananas, koko will not continue to another pile after finishing the pile.

// Given a speed k and piles, the hour it need to consume all piles is sum(ceil(pile[i]/k)) => h. Notice that this is a monotical decrease function. We can use binary search to find the left boundary of k for a given output h.

// note: the function sum(ceil(pile[i]/k)) => h is discrete, meaning that for some h there might not be any corresponding k. if it is not possible to pinpoint a k that produces h, we still need to return a k (in binary serach, we used -1). We use the k that produce a result smaller than h and closest to h. We implement this by also save a result k when finishPileHours(piles, k) < h

func minEatingSpeed(piles []int, h int) int {
	minSpeed, maxSpeed := 1, slices.Max(piles)
	result := maxSpeed
	for minSpeed <= maxSpeed {
		mid := (minSpeed + maxSpeed) / 2
		finishHour := finishPileHours(piles, mid)
		if finishHour > h {
			minSpeed = mid + 1 // speed up, search in right
		} else if finishHour <= h {
			result = mid       // the speed might be the best koko can do
			maxSpeed = mid - 1 // slow down, search in left
		}
	}
	return result
}

// find out the hour it requires to finish all pile in speed k
func finishPileHours(piles []int, k int) int {
	totalHour := 0
	for _, pile := range piles {
		// ceil(pile / k) integer implementation
		totalHour += pile / k
		if pile%k != 0 {
			totalHour++
		}
	}
	return totalHour
}

// @lc code=end

