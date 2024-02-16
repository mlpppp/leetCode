/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */

// @lc code=start

// Given an integer array nums (unsorted, may contain duplicates) and an integer k, return the kth largest element in the array. Can you solve it without sorting? (nums = [3,2,3,1,2,4,5,5,6], k = 4, return 4)

// (Quick Selection), based on quick sort. Select an element `n` and place it in the correct position (O(n)), then compare k with the element's position. If k > idx(n), do quickSelection in arr[idx(n)+1:], if k < idx(n), do quickSelection in arr[:idx(n)], else return n.

// best case: O(n) target always separate the slice into two half: N + N/2 + N/4 + N/8 + ... + 1 = 2N = O(N)
// worst case: O(n^2): target always separate the slice into 1 element and the rest: N + (N - 1) + (N - 2) + ... + 1 = N(N+1)/2 = O(N^2)
// balanced the input: we randomly shuffle the input to avoid the worst case, to get average O(n) performance

// func findKthLargest(nums []int, k int) int {
// 	// shuffle the input
// 	for i := range nums {
// 		j := rand.Intn(i + 1)
// 		nums[i], nums[j] = nums[j], nums[i]
// 	}

// 	// DFS
// 	searchIndex := len(nums) - k
// 	return quickSelection(nums, 0, len(nums)-1, searchIndex)
// }

// // return the value at index `pos` in sorted(arr)
// func quickSelection(arr []int, start, end, pos int) int {
// 	// relocate the first element
// 	target := arr[start]
// 	pSmall, pLargeEq := start+1, start+1
// 	for pSmall <= end {
// 		if arr[pSmall] < target {
// 			// swap(pSmall, pLargeEq)
// 			arr[pSmall], arr[pLargeEq] = arr[pLargeEq], arr[pSmall]
// 			pLargeEq++
// 		}
// 		pSmall++
// 	}
// 	// swap(start, pLargeEq-1)
// 	newPos := pLargeEq - 1
// 	arr[start], arr[newPos] = arr[newPos], arr[start]

// 	// DFS
// 	if pos < newPos {
// 		return quickSelection(arr, start, newPos-1, pos)
// 	} else if pos > newPos {
// 		return quickSelection(arr, newPos+1, end, pos)
// 	} else {
// 		return target
// 	}
// }

// (heap/priority queue): Use a minHeap of size k. Traverse the array and add items to the heap, when heap size is greater k (eq k+1), pop the top of heap. After the iteration, the heap keeps k largest items, and the heap top is the kth largest item.

// ref: https://pkg.go.dev/container/heap

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	numPq := &IntHeap{}
	heap.Init(numPq)

	for _, num := range nums {
		heap.Push(numPq, num)
		if numPq.Len() > k {
			heap.Pop(numPq)
		}
	}

	return heap.Pop(numPq).(int)
}

// @lc code=end

