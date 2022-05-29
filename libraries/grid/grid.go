package grid

import (
	"math"
)

type Position struct {
	x int // x coordinate
	y int // y coordinate
}

type Grid struct {
	x, y  int      // Grid size
	value [][]int  // Grid values
	mark  [][]int  // Grid marks
	start Position // Start position
	end   Position // End position
}

type Node struct {
	pos   Position // Position
	cost  float64  // Cost
	score float64  // Score = Cost + Heuristic
	par   *Node    // Previous node (parent)
}

const (
	V_FREE = iota
	V_WALL
)

const (
	M_NULL = iota
	M_USED
	M_FRONT
	M_PATH
)

func CreateGrid(x, y int, s, e Position) *Grid {
	var g Grid
	g.x = x
	g.y = y

	g.value = make([][]int, x+1)
	for i := 0; i < x+1; i++ {
		g.value[i] = make([]int, y+1)
		for j := 0; j < y+1; j++ {
			g.value[i][j] = V_FREE
		}
	}

	g.mark = make([][]int, x+1)
	for i := 0; i < x+1; i++ {
		g.mark[i] = make([]int, y+1)
		for j := 0; j < y+1; j++ {
			g.mark[i][j] = M_NULL
		}
	}

	g.start = s
	g.end = e
	return &g
}

func (g *Grid) DestroyGrid() {
	g.value = nil
	g.mark = nil
}

func Heuristic(s, t Position) float64 {
	return math.Max(float64(t.x-s.x), float64(t.y-s.y))
}

func CreateNode(par *Node, c float64, pos Position, g *Grid) *Node {
	var n Node
	if par != nil {
		n.par = par
	}
	n.cost = c
	n.pos = pos
	n.score = Heuristic(g.start, g.end)
	return &n
}

func CompareNode(a, b *Node) int {
	return int(a.score) - int(b.score)
}
