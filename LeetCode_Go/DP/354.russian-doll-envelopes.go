/*
 * @lc app=leetcode id=354 lang=golang
 *
 * [354] Russian Doll Envelopes
 */

// @lc code=start

// You are given a 2D array of integers envelopes where envelopes[i] = [wi, hi] represents the width and the height of an envelope. One envelope can fit into another if its width and height are both smaller than the other. Return the maximum number of envelopes you can Russian doll (i.e., put one inside the other).

// Sort envelopes by width and then by height in decreasing order(because ). Then find the `300.longest-increasing-subsequence` in the height sequence.

// 首先，对宽度 w 从小到大排序，确保了 w 这个维度可以互相嵌套，所以我们只需要专注高度 h 这个维度能够互相嵌套即可。

// 其次，两个 w 相同的信封不能相互包含，所以对于宽度 w 相同的信封，对高度 h 进行降序排序，保证二维 LIS 中不存在多个 w 相同的信封（因为题目说了长宽相同也无法嵌套）。

// sorted: [[1,8],[2,3],[5,4],[5,2],[6,7],[6,4]]
// find longest-increasing-subsequence for [8,3,4,2,7,4]

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] || (envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1])
	})
	table := make([]int, len(envelopes))
	for end := 0; end < len(table); end++ {
		maxLIS := 0
		// find the max(lengthOfLIS(nums, i) for (i in range(0, end-1)) and (nums[i] < nums[end]))
		for i := 0; i < end; i++ {
			if envelopes[i][1] < envelopes[end][1] && table[i] > maxLIS {
				maxLIS = table[i]
			}
		}
		if maxLIS > 0 {
			table[end] = maxLIS + 1
		} else {
			table[end] = 1
		}
	}
	// find the max element in table
	maxLIS := 0
	for _, val := range table {
		if val > maxLIS {
			maxLIS = val
		}
	}
	return maxLIS
}

// @lc code=end

