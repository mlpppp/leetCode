/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 */

// @lc code=start

// Given an integer n, return the number of structurally unique BST's (binary search trees) which has exactly n nodes of unique values from 1 to n.

// think all nodes as an ordered array. Any number in the array could be the root of the BST, and numbers in the left would be the left subtree, and numbers in the right would be the right subtree. There are definite number of unique BSTs with size n, which means we can build a table(hashmap) by multipling the possibilities.

// O(n^2)

func numTrees(n int) int {
	sizeUniqueNumTable := make(map[int]int)
	// trival base cases
	sizeUniqueNumTable[0] = 1
	sizeUniqueNumTable[1] = 1
	sizeUniqueNumTable[2] = 2

	for i := 3; i <= n; i++ {
		sizeUniqueNumTable[i] = 0
		for j := 1; j <= i; j++ { // current root node
			right := i - j // number of node in the right subtree
			left := i - right - 1
			sizeUniqueNumTable[i] += sizeUniqueNumTable[left] * sizeUniqueNumTable[right]
			// fmt.Printf("size = %d, curNode = %d, leftNum = %d, rightNum = %d, curEnum = %d\n", i, j, left, right, sizeUniqueNumTable[i])
		}
	}

	return sizeUniqueNumTable[n]
}

// @lc code=end

