/*
 * @lc app=leetcode id=295 lang=golang
 *
 * [295] Find Median from Data Stream
 */

// @lc code=start

// Implement the MedianFinder class to find the median of a data stream.
// MedianFinder() initializes the MedianFinder object.
// void addNum(int num) adds the integer num from the data stream to the data structure.
// double findMedian() returns the median of all elements so far. Answers within 10-5 of the actual answer will be accepted.

// (two heap, O(logN) addNum, O(1)findMedian) use two heaps, each store half of the data in the data structure: the small half in a max heap, and the large half in a min heap.

// AddNum(): rules:
// 1. small half (maxHeap) allows 1 more element when total size is odd
// 2. when input <= maxHeap.top, by default goes to maxHeap, otherwise goes to larger half (minHeap)
// 3. when even or rule 1 are no longer satisfied, adjust by pop the top and insert to the other heap

// findMedian(): if size is even, get both heap top and average. If size is odd, get maxHeap.top

// implement:
// ref: https://pkg.go.dev/container/heap
/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

type MedianFinder struct {
	MaxHeap *maxHeap
	MinHeap *minHeap
}

func Constructor() MedianFinder {
	mf := MedianFinder{
		MaxHeap: &maxHeap{},
		MinHeap: &minHeap{},
	}
	heap.Init(mf.MaxHeap)
	heap.Init(mf.MinHeap)
	return mf
}

func (this *MedianFinder) AddNum(num int) {
	// when both heap are empty, insert to the MaxHeap
	if this.MaxHeap.Len() == 0 {
		heap.Push(this.MaxHeap, num)
		return
	}

	// push to MaxHeap if num <= peek(MaxHeap)
	if num <= (*this.MaxHeap)[0] {
		heap.Push(this.MaxHeap, num)
	} else {
		heap.Push(this.MinHeap, num)
	}

	// if not even, rearrange the two heap's size
	if this.MinHeap.Len() > this.MaxHeap.Len() {
		minTop := heap.Pop(this.MinHeap)
		heap.Push(this.MaxHeap, minTop)
	}

	if this.MaxHeap.Len() > this.MinHeap.Len()+1 {
		maxTop := heap.Pop(this.MaxHeap)
		heap.Push(this.MinHeap, maxTop)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.MaxHeap.Len() == this.MinHeap.Len() {
		v1 := float64((*this.MaxHeap)[0])
		v2 := float64((*this.MinHeap)[0])
		return (v1 + v2) / 2
	} else if this.MaxHeap.Len() == this.MinHeap.Len()+1 {
		return float64((*this.MaxHeap)[0])
	} else {
		fmt.Println("non desirable state")
		return -1
	}
}

// maxHeap & minHeap
type maxHeap []int
type minHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Follow up: If all integer numbers from the stream are in the range [0, 100], how would you optimize your solution? If 99% of all integer numbers from the stream are in the range [0, 100], how would you optimize your solution?
// Bucket sort: each number hold a position, and store the frequency of the integer number. O(1) add and O(n) findMedian

// @lc code=end

