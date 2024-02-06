/*
 * @lc app=leetcode id=1020 lang=golang
 *
 * [1020] Number of Enclaves
 */

// @lc code=start

// You are given an m x n binary matrix grid, where 0 represents a sea cell and 1 represents a land cell. Return the number of land cells in grid for which we cannot move from, in 4-direction, via land, to the outside/boundary of the grid.

// (essentially 1254.number-of-closed-islands) find closed island that is not connected to the outside land. Flood all lands that connects to the outside-grid land with 0, then the remainder islands are all closed island, aka enclaves (200.number-of-islands)

func numEnclaves(grid [][]int) int {
	// flood lands that connect to the outer ring
	for i := 0; i < len(grid); i++ {
		if grid[i][0] == 1 {
			DFSFlood(i, 0, grid)
		}
		if grid[i][len(grid[0])-1] == 1 {
			DFSFlood(i, len(grid[0])-1, grid)
		}
	}
	for i := 0; i < len(grid[0]); i++ {
		if grid[0][i] == 1 {
			DFSFlood(0, i, grid)
		}
		if grid[len(grid)-1][i] == 1 {
			DFSFlood(len(grid)-1, i, grid)
		}
	}

	landCount := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				landCount++
			}
		}
	}
	return landCount
}

// flood the current island
func DFSFlood(i, j int, grid [][]int) {
	// out of bound
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}

	// flood the land
	if grid[i][j] == 0 {
		return
	} else {
		grid[i][j] = 0
	}

	// DFS 4 directions
	directions := [][2]int{[2]int{-1, 0}, [2]int{1, 0}, [2]int{0, -1}, [2]int{0, 1}}
	for _, dir := range directions {
		DFSFlood(i+dir[0], j+dir[1], grid)
	}
}

// @lc code=end

