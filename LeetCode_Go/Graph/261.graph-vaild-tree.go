/**
 * @param n: An integer
 * @param edges: a list of undirected edges
 * @return: true if it's a valid tree, or false
 */

// https://www.lintcode.com/problem/178/
// Example 1:
// Input: n = 5 edges = [[0, 1], [0, 2], [0, 3], [1, 4]]
// Output: true.

// Example 2:
// Input: n = 5 edges = [[0, 1], [1, 2], [2, 3], [1, 3], [1, 4]]
// Output: false.

// Given n nodes labeled from 0 to n - 1 and a list of undirected edges (each edge is a pair of nodes), write a function to check whether these edges make up a valid tree.

// (DFS detect cycle): a graph is a tree <=> is it connected and contain no cycle.
// 1. connected: if a undirected graph is connected without cycle => it can be traversed completely with one DFS call from any node
// 2. detect cycle: see (207.course-schedule). Add to `trace` at preorder, remove from `trace` at postorder

func ValidTree(n int, edges [][]int) bool {
	// create adjacency list
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = []int{}
	}
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}

	// dfs node 0
	trace := make([]bool, n)
	visitedCount := 0
	isCyclic := detectCycle(0, -1, &visitedCount, trace, adj)
	return !isCyclic && visitedCount == n
}

func detectCycle(node, prevNode int, visitedCount *int, trace []bool, adj [][]int) bool {
	if trace[node] {
		return true
	}

	trace[node] = true
	*visitedCount += 1
	for _, neighbor := range adj[node] {
		if neighbor == prevNode { // undirected graph(bi-directional connect), prevent going back
			continue
		}
		if detectCycle(neighbor, node, visitedCount, trace, adj) {
			return true
		}
	}
	trace[node] = false
	return false
}