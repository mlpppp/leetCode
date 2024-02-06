/*
 * @lc app=leetcode id=210 lang=golang
 *
 * [210] Course Schedule II
 */

// @lc code=start

// There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [a_i, b_i] indicates that you must take course b_i first if you want to take course a_i. Return the ordering of courses you should take to finish all courses. If there are many valid answers, return any of them. If it is impossible to finish all courses, return an empty array.

// topological sort: O(n): 1. detect cycle, if exists cycle, return empty array. (207.course-schedule) 2. postorder traverse the graph, return reversed(postorder)

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := makeGraph(numCourses, prerequisites)
	trace := make([]bool, numCourses)
	visited := make([]bool, numCourses)
	postorder := []int{}
	for i := 0; i < len(graph); i++ {
		if detectCycleAndOrder(graph, i, trace, visited, &postorder) {
			return []int{}
		}
	}
	slices.Reverse(postorder)
	return postorder

}

func detectCycleAndOrder(graph [][]int, node int, trace []bool, visited []bool, postorder *[]int) bool {
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
		if detectCycleAndOrder(graph, neighbor, trace, visited, postorder) {
			return true
		}
	}
	// postorder
	*postorder = append(*postorder, node)
	trace[node] = false
	return false
}

// convert the problem into adjacency list: course <- prerequisites
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

