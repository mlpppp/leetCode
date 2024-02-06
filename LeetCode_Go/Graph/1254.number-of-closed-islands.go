/*
 * @lc app=leetcode id=1254 lang=golang
 *
 * [1254] Number of Closed Islands
 */

// @lc code=start

// Given a 2D grid consists of 0s (land) and 1s (water) (assume lands are outside of the grid). Return the number of closed islands, a closed island is an island totally (all left, top, right, bottom) surrounded by 1s.

// flood all outer ring islands that connects to the outside-grid land with 1, then the remainder island are all closed island (200.number-of-islands)

func closedIsland(grid [][]int) int {
	// flood lands that connect to the outer ring
	for i := 0; i < len(grid); i++ {
		if grid[i][0] == 0 {
			DFSFlood(i, 0, grid)
		}
		if grid[i][len(grid[0])-1] == 0 {
			DFSFlood(i, len(grid[0])-1, grid)
		}
	}
	for i := 0; i < len(grid[0]); i++ {
		if grid[0][i] == 0 {
			DFSFlood(0, i, grid)
		}
		if grid[len(grid)-1][i] == 0 {
			DFSFlood(len(grid)-1, i, grid)
		}
	}

	// 200.number-of-islands DFS algorithm
	islandCount := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				islandCount++
				DFSFlood(i, j, grid)
			}
		}
	}
	return islandCount
}

// flood the current island
func DFSFlood(i, j int, grid [][]int) {
	// out of bound
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}

	// flood the land
	if grid[i][j] == 1 {
		return
	} else {
		grid[i][j] = 1
	}

	// DFS 4 directions
	directions := [][2]int{[2]int{-1, 0}, [2]int{1, 0}, [2]int{0, -1}, [2]int{0, 1}}
	for _, dir := range directions {
		DFSFlood(i+dir[0], j+dir[1], grid)
	}
}

// @lc code=end

