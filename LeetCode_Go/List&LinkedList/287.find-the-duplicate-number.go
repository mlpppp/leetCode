/*
 * @lc app=leetcode id=287 lang=golang
 *
 * [287] Find the Duplicate Number
 */

// @lc code=start

// Given an unsorted array of integers nums containing n + 1 integers in the range [1, n] inclusive. There are exactly 1 number in nums that repeats two or more times. Return this repeated number without modifying the array nums in O(n) time and O(1) space. (eg. [1, 2, 2, 2, 4, 5])

// https://www.youtube.com/watch?v=wjYnzkAhcNk

// (142.linked-list-cycle-ii) find the start of cycle of a linked list: use the index of array as linked list node value, use the value in array as the `Next` field. Then there must be multiple node pointing to one node, the node also is the start of a cycle. This node is the repeated number. Use nums[0] as entry point, use two pointers to find the start of a cycle  // 不懂，只能死记硬背

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

// @lc code=end

