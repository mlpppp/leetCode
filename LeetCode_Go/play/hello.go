package main

import "fmt"

// https://labuladong.github.io/algo/data-structure/prefix-sum/

// build an diff array from the input array. When increment(start, end, inc), raise diff[start] by inc, and lower diff[end+1] by inc (provided that end+1 not out of range). To return the array after k operation, build it from the beginning (assume start from 0).

type DiffArray struct {
	Diff []int
}

func constructor(nums []int) *DiffArray {
	diff := make([]int, len(nums))
	// 根据初始数组构造差分数组
	diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return &DiffArray{Diff: diff}
}
func (d *DiffArray) increment(start, end, inc int) {
	d.Diff[start] += inc
	if end+1 < len(d.Diff) {
		d.Diff[end+1] -= inc
	}
}

func (d *DiffArray) getResult() []int {
	ans := make([]int, len(d.Diff))
	for i := range d.Diff {
		if i == 0 {
			ans[0] = d.Diff[0]
		} else {
			ans[i] = ans[i-1] + d.Diff[i]
		}
	}
	return ans
}

func getModifiedArray(init []int, updates [][]int) []int {
	diff := constructor(init)
	for _, update := range updates {
		i, j, inc := update[0], update[1], update[2]
		diff.increment(i, j, inc)
	}

	return diff.getResult()
}

func main() {
	// 输入: length = [0,0,0,0,0], updates = [[1,3,2],[2,4,3],[0,2,-2]]
	// 输出: [-2,0,3,5,3]
	nums := make([]int, 5)
	fmt.Println(getModifiedArray(nums, [][]int{{1, 3, 2}, {2, 4, 3}, {0, 2, -2}}))

	// fmt.Println(n.SumRegion(1, 1, 2, 2))
	// fmt.Println(n.SumRegion(1, 2, 2, 4))

	// board := {][]byte{{'o', 'a', 'b', 'n'}, {'o', 't', 'a', 'e'}, {'a', 'h', 'k', 'r'}, {'a', 'f', 'l', 'v'}}
	// words := []string{"oa", "oaa"}
	// fmt.Println(findWords(board, words))

	// param_3 := obj.StartsWith(prefix);6}, 1))
	// fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 2))
	// fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 3))
	// fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
	// fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 5))
	// fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 6))
	// fmt.maxPathSum(countSubstrings("abcbcddabbdccd"))

	// fmt.Println(isPalindrome("race a car"))

	// fmt.Println(reverseWords("  hello world  "))
	// fmt.Println(reverseWords("a good   example"))
	// fmt.Println(maxProduct([]int{-2, 0, -1}))
	// fmt.Println(maxProduct([]int{2, -2, 3, 6, -8, 2}))
	// fmt.Println(maxProduct([]int{2, -2, 0, 3, 6, -8, 2}))

	// fmt.Println(rangeBitwiseAnd("11", "1"))
	// fmt.Println(rangeBitwiseAnd("0", "0"))
	// fmt.Println(rangeBitwiseAnd("01", "1"))

}
