/*
 * @lc app=leetcode id=130 lang=golang
 *
 * [130] Surrounded Regions
 */

// @lc code=start

// Given an m x n matrix board containing 'X' and 'O', flipping all 'O' regions that are 4-directionally surrounded by 'X', into 'X's. (very similar problem to 1254. Number of Closed Islands)

// Input: board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
// Output: [["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]

// surrounded by "X" :=> not connect (4 directional directly) to any "O" in the 4 border sides.

// DFS solution: DFS flood all outer ring islands (turn them into "#"), then DFS flood all remain islands "O" into "X", finally resume all "#" cell back to "O"

// unionFind solution:
// nodes: [i, j], use hashmap; edges: four directional directly connection
// build unionFind to all "O" cells in the board, make all outer ring islands the same CC, and the rest "surround island" their own CCs.
// then traverse the inner board, for each cell that does not connect to the outerring-island-CC, flip it.

func solve(board [][]byte) {
	uf := newUnionFind()
	uf.addNode(-1, -1, -1, -1) // dummy cell at(-1, -1), point to itself

	// scan the board, create Nodes for all "O", also add outer-ring cells to dummy CC
	for i := 0; i < len(board); i++ {
		// outer ring cells
		for _, j := range []int{0, len(board[0]) - 1} {
			if board[i][j] == 'O' {
				uf.addNode(i, j, -1, -1) // connect to dummy
			}
		}

		for j := 1; j < len(board[0])-1; j++ {
			if board[i][j] == 'O' {
				if i == 0 || i == len(board)-1 { // outer ring cells
					uf.addNode(i, j, -1, -1) // connect to dummy
				} else { // inner part
					uf.addNode(i, j, i, j) // connect to dummy
				}
			}
		}
	}

	// scan a second time in inner board, and connect all 'O' to 'Os' in its 4 directions
	direction := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := 1; i < len(board)-1; i++ {
		for j := 1; j < len(board[0])-1; j++ {
			if board[i][j] == 'O' {
				for _, dir := range direction {
					newI, newJ := i+dir[0], j+dir[1]
					if board[newI][newJ] == 'O' {
						uf.connect([2]int{i, j}, [2]int{newI, newJ})
					}
				}
			}
		}
	}

	// scan the third time for inner board, filp all "O" that does not connect to dummy
	for i := 1; i < len(board)-1; i++ {
		for j := 1; j < len(board[0])-1; j++ {
			if board[i][j] == 'O' {
				if !uf.isConnect([2]int{i, j}, [2]int{-1, -1}) {
					board[i][j] = 'X'
				}
			}
		}
	}

	// for row := range board {
	// 	fmt.Println(string(board[row]))
	// }
}

type UnionFind struct {
	Count   int
	Parents map[[2]int][2]int
}

func newUnionFind() *UnionFind {
	uf := &UnionFind{}
	uf.Count = 0
	uf.Parents = make(map[[2]int][2]int)
	// uf.Parent = make([]int, n)
	return uf
}

func (uf *UnionFind) addNode(i, j int, parentI, parentJ int) {
	uf.Parents[[2]int{i, j}] = [2]int{parentI, parentJ}
	uf.Count++
}

func (uf *UnionFind) count() int {
	return uf.Count
}

func (uf *UnionFind) connect(a, b [2]int) {
	iRoot := uf.findRoot(a)
	jRoot := uf.findRoot(b)

	if iRoot != jRoot {
		uf.Parents[iRoot] = jRoot
		uf.Count--
	}
}

// find root
func (uf *UnionFind) findRoot(n [2]int) [2]int {
	if uf.Parents[n] != n {
		root := uf.findRoot(uf.Parents[n])
		uf.Parents[n] = root
	}
	return uf.Parents[n]
}

func (uf *UnionFind) isConnect(a, b [2]int) bool {
	iRoot := uf.findRoot(a)
	jRoot := uf.findRoot(b)
	return iRoot == jRoot
}

// @lc code=end

