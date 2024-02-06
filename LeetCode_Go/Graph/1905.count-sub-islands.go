/*
 * @lc app=leetcode id=1905 lang=golang
 *
 * [1905] Count Sub Islands
 */

// @lc code=start

// You are given two m x n binary matrices `grid1` and `grid2` containing only 0's (representing water) and 1's (representing land), the outside are all waters 0's. An island is a group of 1's connected 4-directionally (horizontal or vertical). An island in grid2 is considered a `sub-island` if there is an island in grid1 that contains all the cells that make up this island in grid2. Return the number of sub-islands in grid2.

// (ref 200.number-of-islands). Rule: for any cells in a `sub-island`, the counterpart in grid1 must also be land. Use grid1 as template. Traverse grid2, and run DFS algorithm DFS(i, j, isSubIsland) to every island in grid2. During the traverse, if there exist a land cell in grid2 that is not land cell in grid1, we know that the island is not subIsland and we update isSubIsland to false.

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	subIslandCount := 0
	for i := 0; i < len(grid2); i++ {
		for j := 0; j < len(grid2[0]); j++ {
			if grid2[i][j] == 1 { // found an island
				isSubIsland := true
				DFS(i, j, grid1, grid2, &isSubIsland)
				if isSubIsland {
					subIslandCount++
				}
			}
		}
	}
	return subIslandCount
}

// determine if the island isSubIsland while flooding the island
func DFS(i, j int, grid1, grid2 [][]int, isSubIsland *bool) {
	// out of bound
	if i < 0 || j < 0 || i >= len(grid2) || j >= len(grid2[0]) {
		return
	}

	// flood the i, j
	if grid2[i][j] == 0 { // already water
		return
	} else if grid2[i][j] == 1 {
		if grid1[i][j] == 0 {
			*isSubIsland = false
		}
		grid2[i][j] = 0
	}
	// DFS
	directions := [][2]int{[2]int{0, -1}, [2]int{0, 1}, [2]int{-1, 0}, [2]int{1, 0}}
	for _, dir := range directions {
		DFS(i+dir[0], j+dir[1], grid1, grid2, isSubIsland)
	}
}

// @lc code=end

