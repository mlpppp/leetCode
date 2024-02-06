/*
 * @lc app=leetcode id=797 lang=golang
 *
 * [797] All Paths From Source to Target
 */

// @lc code=start

// Given a directed acyclic graph (DAG) of n nodes labeled from 0 to n - 1 (adjacent list), find all possible paths from node 0 to node n - 1 and return them in any order (return array of array(path)).

// DFS traverse logic, don't use visited array, append to path when enter an node, and undo appending path when finished the node. Record path when reach target.

func allPathsSourceTarget(graph [][]int) [][]int {
	ans := [][]int{}
	path := []int{}
	DFS(graph, 0, &path, &ans)
	return ans
}

func DFS(graph [][]int, v int, path *[]int, ans *[][]int) {
	*path = append(*path, v)
	// reach last node, record path
	if v == len(graph)-1 {
		temp := make([]int, len((*path)))
		copy(temp, *path)
		*ans = append(*ans, temp)
	}

	for _, i := range graph[v] {
		// fmt.Printf("%v is calling DFS(%v)\n", *path, i)
		DFS(graph, i, path, ans)
	}

	*path = (*path)[:len(*path)-1]
}

// @lc code=end

