package main

import (
	"fmt"
	"strings"
)

// To find the median of the merged array, it can be roughly partition into two parts, either (even total len) two parts are the same length and the median is (part1[-1] + part2[0])/2, or (odd total len) the part1 is one element less and the median is part2[0]. We do binary search to partition nums1, and deduce the partition of nums2 based on `the size of nums1's left partition` and `size of part1`. Given index mid that partition nums1, there are mid+1 elements of num1 that contribute to part1, and then we can deduce that there are floor((len(nums1)+len(nums2)/2)-(mid+1) element in the num2 that contribute to part1. Then we test if the proposed partition really produces the sorted merged array (aka every elements in the part1 is smaller or equal to every elements in the part2), if so, the median is found. if not, we adjust the nums1' partition. Note1: watch for out of bound edgy cases. if either the partition to nums1 or nums2 are out of bounds, we assign dummy value inf or -inf. Note2: strict property of the array to be searched: it must be the longer array; some part of nums1 must be in the part1, otherwise the correct median cannot be found. (see implementation for fixes)

func solveNQueens(n int) [][]string {
	ans := [][]int{}
	backtracking([]int{}, n, &ans)
	// format answers
	ansString := make([][]string, len(ans))
	for i, sol := range ans {
		ansString[i] = make([]string, n)
		for j, solRow := range sol {
			dots := []byte(strings.Repeat(".", n))
			dots[solRow] = 'Q'
			ansString[i][j] = string(dots)
		}
	}
	return ansString
}

func backtracking(curSols []int, n int, ans *[][]int) {
	// exit with a new solution
	if len(curSols) == n {
		temp := make([]int, n)
		copy(temp, curSols)
		*ans = append(*ans, temp)
		return
	}

	// find available columns
	curRow := len(curSols) // the row we are current working with
	for curCol := 0; curCol < n; curCol++ {
		// skip unvaild cols (use col and used diagonals)
		isBreak := false
		for solRow, solCol := range curSols {
			if curCol == solCol || curCol == solCol+(curRow-solRow) || curCol == solCol-(curRow-solRow) {
				isBreak = true
				break
			}
		}
		if isBreak {
			continue
		}

		// do choice
		curSols = append(curSols, curCol)
		backtracking(curSols, n, ans)
		// undo choice
		curSols = curSols[:len(curSols)-1]
	}
}

func main() {

	fmt.Println(solveNQueens(6))

	// fmt.Println(rangeBitwiseAnd("11", "1"))
	// fmt.Println(rangeBitwiseAnd("0", "0"))
	// fmt.Println(rangeBitwiseAnd("01", "1"))

}
