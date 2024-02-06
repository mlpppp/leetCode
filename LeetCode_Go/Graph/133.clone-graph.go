/*
 * @lc app=leetcode id=133 lang=golang
 *
 * [133] Clone Graph
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

//  Given a reference of a node in a connected undirected graph. Return a deep copy (clone) of the graph.

// DFS + visited(hashmap), when traverse a ref graph node: create clone in preorder, create connections to the clone in postorder, return the clone node. Visited hashmap record clone node of a visited ref node.

func cloneGraph(node *Node) *Node {
	// hashmap: refNode->cloneNode
	visited := make(map[*Node]*Node)
	if node == nil {
		return nil
	}
	return DFS(node, visited)
}

func DFS(node *Node, visited map[*Node]*Node) *Node {
	// clone already exists
	if clone, exists := visited[node]; exists {
		return clone
	}
	// create clone
	clone := Node{Val: node.Val, Neighbors: []*Node{}}
	visited[node] = &clone

	for _, n := range node.Neighbors {
		clone.Neighbors = append(clone.Neighbors, DFS(n, visited))
	}
	return &clone
}

// @lc code=end

