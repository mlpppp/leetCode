/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start

// Given an array nums of size n, return the majority element. The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array. Follow-up: Could you solve the problem in linear time and in O(1) space?

// brute force: O(n) + O(n): hashmap to count the number of appearance

// func majorityElement(nums []int) int {
// 	count := make(map[int]int)
// 	for _, num := range nums {
// 		if _, exists := count[num]; exists {
// 			count[num] += 1
// 		} else {
// 			count[num] = 1
// 		}
// 		if count[num] > len(nums)/2 {
// 			return num
// 		}
// 	}
// 	return -1
// }

// bit manipulation O(32n) or O(64n): compare the bit representation all numbers in nums: for each bit position, find the majority bit (0 or 1), which is the bit in the answer in that position. (if a number is majority in an array, at any position, the number's bit is also majority)
// func majorityElement(nums []int) int {
// 	ans := 0
// 	majorityCap := len(nums) / 2
// 	mask := 1

// 	for i := 0; i < 64; i++ { // for each bit position
// 		// count '1' appears
// 		count1 := 0
// 		for _, num := range nums { // iterate all numbers and count '1'
// 			if num&mask == mask {
// 				count1++
// 			}
// 		}
// 		// set position '1' if '1' is the majority
// 		if count1 > majorityCap {
// 			ans |= (1 << i)
// 		}
// 		// move to next position
// 		mask <<= 1
// 	}
// 	return ans
// }

// Boyer-Moore Algorithm: thinking each unique number is a team. Teams fight each other: when two ppl from two different teams meet, they eliminate each other. There must be one majority team, and the majority team must stand to the end of game. Design the algorithm: Iterate through the array, we keep track of the current majority team's count we have, when a different number is meet, we decrese the count by one (they eliminate each other). Then the count becomes zero, we start to track the next different number. When the iteration finish, the number we are still tracking is the answer.

func majorityElement(nums []int) int {
	curTrack := -1
	count := 0
	for _, num := range nums {
		if count == 0 {
			curTrack = num
			count = 1
		} else {
			if num == curTrack {
				count++
			} else {
				count--
			}
		}
	}
	return curTrack
}

// @lc code=end

