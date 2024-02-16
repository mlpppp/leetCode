/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

// @lc code=start

// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

// (hashmap + bucket sort) traverse and count each element's frequency in hashmap `count`. then traverse the hashmap. Then bucket sort the frequencies in `count`: the possible frequency for any element is at most len(nums). So `buckets` size if len(nums). traverse the hashmap and move frequencies to `buckets`, and then reversely traverse buckets to retrieve the last k elements from buckets.

func topKFrequent(nums []int, k int) []int {
	// build frequency map: count
	count := make(map[int]int)
	for _, num := range nums {
		count[num] += 1
	}

	//  build buckets from freq to element
	buckets := make([][]int, len(nums)+1)
	for key, val := range count {
		buckets[val] = append(buckets[val], key)
	}

	// get top k elements
	ans := []int{}
	for i := len(buckets) - 1; i > 0 && k > 0; i-- {
		for j := 0; j < len(buckets[i]) && k > 0; j++ {
			ans = append(ans, buckets[i][j])
			k--
		}
	}
	return ans
}

// @lc code=end

