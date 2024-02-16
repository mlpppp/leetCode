/*
 * @lc app=leetcode id=990 lang=golang
 *
 * [990] Satisfiability of Equality Equations
 */

// @lc code=start

// You are given an array of strings equations that represent relationships between variables where each string equations[i] is of length 4 and takes one of two different forms: "x_i==y_i" or "x_i!=y_i". Here, x_i and y_i are one lowercase letters (not necessarily different) that represent a variable names. Return true if it is possible to assign integers to variable names so as to satisfy all the given equations, or false otherwise.

// Input: equations = ["a==b","b!=a"]
// Output: false

// Input: equations = ["b==a","a==b"]
// Output: true

// Use unionFind, for all "==" equations add the connection to unionFind. Then for all "a!=b" equations, if a and b are connected, return false. Return true at the end.

func equationsPossible(equations []string) bool {
	uf := newUnionFind(26) // 26 lower case letters

	// for all "==" equations add the connection to unionFind
	for _, eq := range equations {
		if eq[1] == '=' {
			node1 := eq[0] - 'a'
			node2 := eq[3] - 'a'
			uf.connect(int(node1), int(node2))
		}
	}

	// Then for all "a!=b" equations, if a and b are connected, return false.
	for _, eq := range equations {
		if eq[1] == '!' {
			node1 := eq[0] - 'a'
			node2 := eq[3] - 'a'
			if uf.isConnected(int(node1), int(node2)) {
				return false
			}
		}
	}

	return true
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

// connect two nodes
func (uf *UnionFind) connect(i, j int) {
	iRoot := uf.findRoot(i)
	jRoot := uf.findRoot(j)

	if iRoot != jRoot {
		uf.Parent[iRoot] = jRoot
		uf.Count--
	}
}

func (uf *UnionFind) isConnected(i, j int) bool {
	iRoot := uf.findRoot(i)
	jRoot := uf.findRoot(j)
	return iRoot == jRoot
}

// find root
func (uf *UnionFind) findRoot(n int) int {
	if uf.Parent[n] != n {
		root := uf.findRoot(uf.Parent[n])
		uf.Parent[n] = root
	}
	return uf.Parent[n]
}

// @lc code=end

