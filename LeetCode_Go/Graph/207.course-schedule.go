/*
 * @lc app=leetcode id=207 lang=golang
 *
 * [207] Course Schedule
 */

// @lc code=start

// There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [a_i, b_i] indicates that you must take course b_i first if you want to take course a_i. Return true if you can finish all courses. Otherwise, return false.

// (DFS + visited) Detect cycle in graph (if there is a cycle, topological sort is not possible). Detect cycle for every nodes (there might be non-connected component), return false immediately if find cycle in any node, mark a node `visited` that has been detected so we can skip it. use hashset or []bool `trace` to record nodes alone the visit trace: add at preorder, remove at postorder. If a current node is already in `trace`, there is a cycle.

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := makeGraph(numCourses, prerequisites)
	trace := make([]bool, numCourses)
	visited := make([]bool, numCourses)
	for i := 0; i < len(graph); i++ {
		if detectCycle(graph, i, trace, visited) {
			return false
		}
	}
	return true
}

func detectCycle(graph [][]int, node int, trace []bool, visited []bool) bool {
	// if a node is seen before, there is a cycle
	if trace[node] {
		return true
	}

	// skip node that has been DFS, and apparently doesn't detected a cycle
	if visited[node] {
		return false
	}

	// preorder
	trace[node] = true
	visited[node] = true
	for _, neighbor := range graph[node] {
		if detectCycle(graph, neighbor, trace, visited) {
			return true
		}
	}
	trace[node] = false
	return false
}

// convert the problem into adjacency list.
func makeGraph(numCourses int, prerequisites [][]int) [][]int {
	graph := make([][]int, numCourses)
	for i, _ := range graph {
		graph[i] = []int{}
	}
	for _, edge := range prerequisites {
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	return graph
}

// @lc code=end

