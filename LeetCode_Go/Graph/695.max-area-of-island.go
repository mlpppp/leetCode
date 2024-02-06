/*
 * @lc app=leetcode id=695 lang=golang
 *
 * [695] Max Area of Island
 */

// @lc code=start

// You are given an m x n binary matrix grid. An island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water. Return the maximum area of an island in grid. If there is no island, return 0.

// (essentially 200.number-of-islands), count the area of a island when traverse it with DFS.

func maxAreaOfIsland(grid [][]int) int {
	maxIslandArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				count := 0
				DFS(i, j, grid, &count)
				maxIslandArea = max(maxIslandArea, count)
			}
		}
	}
	return maxIslandArea
}

// flood the current island and count its landmass
func DFS(i, j int, grid [][]int, count *int) {
	// out of bounds
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}
	// flood the i, j
	if grid[i][j] == 0 {
		return
	} else {
		*count++
		grid[i][j] = 0
	}
	// DFS
	directions := [][2]int{[2]int{0, -1}, [2]int{0, 1}, [2]int{-1, 0}, [2]int{1, 0}}
	for _, dir := range directions {
		DFS(i+dir[0], j+dir[1], grid, count)
	}
}

// @lc code=end

