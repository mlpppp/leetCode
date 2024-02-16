// 323.Number-of-connected-components-in-an-undirected-graph
// https://www.lintcode.com/problem/3651/

// Given a graph of n nodes, and an array edges containing all edges [i, j] from node i to node j. Return the number of connected components in the graph.

// Example:
// Input: n = 5, edges = [[0, 1],[1, 2],[3, 4]]
// Output: 2

// use unionFind data structure

func CountComponents(n int, edges [][]int) int {
	uf := newUnionFind(n)
	for _, edge := range edges {
		uf.connect(edge[0], edge[1])
	}
	return uf.count()
}

type UnionFind struct {
	Count  int
	Parent []int
}

// constructor
func newUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for node, _ := range parent {
		parent[node] = node
	}
	uf := &UnionFind{
		Count:  n,
		Parent: parent,
	}
	return uf
}

func (uf *UnionFind) count() int {
	return uf.Count
}

// connect two nodes
func (uf *UnionFind) connect(i, j int) {
	iRoot := uf.findRoot(i)
	jRoot := uf.findRoot(j)

	if iRoot != jRoot {
		uf.Parent[iRoot] = jRoot
		uf.Count--
	}
}

// find parent
func (uf *UnionFind) findRoot(n int) int {
	if uf.Parent[n] != n {
		root := uf.findRoot(uf.Parent[n])
		uf.Parent[n] = root
	}
	return uf.Parent[n]
}
