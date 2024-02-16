/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */

// @lc code=start

// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence in O(n). Eg: [100,4,200,1,3,2] -> [1,2,3,4], [100], [200] -> return 4

// (Hashmap: O(n)) the start of any sequence, says n, has the condition: `n exist in array, but n-1 does not exist`. We identify all start of a sequence and count how many subsequent elements (n+1, n+2...) also exist in the array. Implement with hashmap/hashset: scan once and save all elements to hashmap. Then scan again to complete the algorithm in O(1) each operation.

func longestConsecutive(nums []int) int {
	hashmap := make(map[int]bool)
	for _, num := range nums {
		hashmap[num] = true
	}
	ans := 0
	for _, num := range nums {
		// skip element that is not start of a sequence
		if _, exists := hashmap[num-1]; exists {
			continue
		}
		// count number of consecutive elements from num
		count := 0
		exists := true
		for exists {
			count++
			num++
			_, exists = hashmap[num]
		}
		ans = max(ans, count)
	}
	return ans
}

// @lc code=end

