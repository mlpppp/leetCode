/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */

// @lc code=start

// Given an m x n 2D binary grid which represents a map of '1's (land) and '0's (water), return the number of islands. An island is formed by connecting adjacent lands horizontally or vertically (not diagonally). You may assume all four edges of the grid are all surrounded by water.

// FloodFill (visited+DFS): we iterate for every given cell for unvisited land '1', when a land is found, it is an island (island++). From the land we explore the 4 directions (NESW) for land '1' with DFS. we use 'visited' to mark a visited land. DFS guarantee to visit all land in a island. Then we continue the iterate to find the next unvisited land.

func numIslands(grid [][]byte) int {
	islandCount := 0
	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				islandCount++
				DFS(i, j, grid, visited)
			}
		}
	}
	return islandCount
}

// mark every land in current island as visited
func DFS(i, j int, grid [][]byte, visited [][]bool) {
	// out of bound
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}
	// DFS only mark unvisited land
	if grid[i][j] == '0' || visited[i][j] == true {
		return
	}
	visited[i][j] = true
	// DFS 4 directions
	directions := [][2]int{[2]int{-1, 0}, [2]int{1, 0}, [2]int{0, -1}, [2]int{0, 1}}
	for _, dir := range directions {
		DFS(i+dir[0], j+dir[1], grid, visited)
	}
}

// @lc code=end

