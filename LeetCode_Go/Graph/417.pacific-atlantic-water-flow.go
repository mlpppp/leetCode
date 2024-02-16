/*
 * @lc app=leetcode id=417 lang=golang
 *
 * [417] Pacific Atlantic Water Flow
 */

// @lc code=start

// Given a m x n matrix `heights` where heights[i][j] is the height of an island in position (i, j). Island represented by heights matrix is surrounded by the Pacific Ocean (left and top) and Atlantic Ocean (right and bottom), the height level of Ocean is 0. The island receives a lot of rain, and the rain water can flow to 4 directions provided that the next height is less or equal to the current height. Return a 2D list result where result[i] is the coordinates (i, j) where rain water can flow from the cell to both the Pacific and Atlantic oceans.

// https://www.youtube.com/watch?v=s-VkcjHqkGI

// O(nm) DFS all cell that connects directly to Pacific Ocean, expand reversely to find all cell that flows to Pacific Ocean. Then DFS all cell that connects directly to Atlantic Ocean, expand reversely to find all cell that flows to Atlantic Ocean. Finally find all cell that flows to both Pacific Ocean and Atlantic Ocean (overlap of the previous two results' containers).

// DFS reversely: expand to heights[i+1] >= height[i]

func pacificAtlantic(heights [][]int) [][]int {
	flowToAtlantic := make(map[[2]int]bool) // save cells that can flow to Atlantic, also used as `visited`
	flowToPacific := make(map[[2]int]bool)  // save cells that can flow to Pacific, also used as `visited`

	// expand flowToAtlantic and flowToPacific from all shores
	for i := 0; i < len(heights); i++ {
		markFlowable(i, 0, heights, flowToPacific)                  // first col to Pacific
		markFlowable(i, len(heights[0])-1, heights, flowToAtlantic) // last col to Atlantic
	}

	for j := 0; j < len(heights[0]); j++ {
		markFlowable(0, j, heights, flowToPacific)               // first row to Pacific
		markFlowable(len(heights)-1, j, heights, flowToAtlantic) // last row to Atlantic
	}

	// find the overlap of two containers
	ans := [][]int{}
	for key, _ := range flowToAtlantic {
		if _, exists := flowToPacific[key]; exists {
			ans = append(ans, []int{key[0], key[1]})
		}
	}

	return ans
}

// add a cell to either flowToAtlantic, or flowToPacific
func markFlowable(i, j int, heights [][]int, visited map[[2]int]bool) {
	if _, exists := visited[[2]int{i, j}]; exists { // already visited
		return
	}

	visited[[2]int{i, j}] = true

	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, dir := range directions {
		newI, newJ := i+dir[0], j+dir[1]
		if newI >= 0 && newI < len(heights) && // not out of bounds
			newJ >= 0 && newJ < len(heights[0]) &&
			heights[i][j] <= heights[newI][newJ] { // can be flow from
			markFlowable(newI, newJ, heights, visited)
		}
	}
}

// test case ********************************

// graph := make([][]int, 5)
// graph[0] = []int{1, 2, 2, 3, 5}
// graph[1] = []int{3, 2, 3, 4, 4}
// graph[2] = []int{2, 4, 5, 3, 1}
// graph[3] = []int{6, 7, 1, 4, 5}
// graph[4] = []int{5, 1, 1, 2, 4}
// fmt.Println(pacificAtlantic(graph))

// 废案 **************************************************

// DFS choice: ESWN, provided that the next step heights[i+1] <= height[i].

// goal/states: reach both: 4. reach Atlantic only: return 3, reach Pacific only: return 2, reach neither: return 1, uninitialized(not visited): 0

// base case: out of bound, use padding to table: +2 dimension, Pacific: 2, Altantic: 3

// state transition:
// 4: return 4
// for 4 dir exist 3 and 2: return 4
// for 4 dir exist only 3 or only 2: return 3 or 2
// 1: default case before loop, return 1 if 4 dirs are not feasible or return 1

// func pacificAtlantic(heights [][]int) [][]int {
// 	// center init with 0, Ocean init with -1 or -2
// 	state := make([][]int, len(heights)+2)
// 	for i, _ := range state {
// 		state[i] = make([]int, len(heights[0])+2)
// 	}
// 	for i := 1; i < len(state)-1; i++ {
// 		state[i][0] = 2               // Pacific
// 		state[i][len(state[0])-1] = 3 // Altantic
// 	}
// 	for i := range state[0] {
// 		state[0][i] = 2            // Pacific
// 		state[len(state)-1][i] = 3 // Altantic
// 	}

// 	// padding height, fill ocean with 0
// 	for i, _ := range heights {
// 		heights[i] = append([]int{0}, heights[i]...)
// 		heights[i] = append(heights[i], 0)
// 	}
// 	heights = append([][]int{make([]int, len(heights[0]))}, heights...)
// 	heights = append(heights, make([]int, len(heights[0])))

// 	// 不走回头路
// 	path := make([][]bool, len(heights)+2)
// 	for i, _ := range path {
// 		path[i] = make([]bool, len(heights[2])+2)
// 	}

// 	ans := [][]int{}
// 	for i := 1; i < len(heights)-1; i++ {
// 		for j := 1; j < len(heights[0])-1; j++ {
// 			if DFS(i, j, heights, state, &path) == 4 {
// 				ans = append(ans, []int{i - 1, j - 1})
// 			}
// 		}
// 	}
// 	return ans
// }

// func DFS(i, j int, heights, state [][]int, path *[][]bool) int {
// 	if state[i][j] != 0 {
// 		return state[i][j]
// 	}

// 	(*path)[i][j] = true

// 	directions := [][]int{[]int{-1, 0}, []int{1, 0}, []int{0, -1}, []int{0, 1}}
// 	toPacific, toAtlantic := false, false
// 	for _, dir := range directions {
// 		if heights[i][j] >= heights[i+dir[0]][j+dir[1]] &&
// 			!(*path)[i+dir[0]][j+dir[1]] {
// 			nextState := DFS(i+dir[0], j+dir[1], heights, state, path)
// 			if nextState == 4 {
// 				state[i][j] = 4
// 				return 4
// 			}
// 			if nextState == 3 {
// 				toAtlantic = true
// 			}
// 			if nextState == 2 {
// 				toPacific = true
// 			}
// 		}
// 	}
// 	(*path)[i][j] = false
// 	if toPacific && toAtlantic {
// 		state[i][j] = 4
// 		return 4
// 	} else if toPacific {
// 		state[i][j] = 2
// 	} else if toAtlantic {
// 		state[i][j] = 3
// 	} else {
// 		state[i][j] = 1
// 	}
// 	return state[i][j]
// }
// @lc code=end

