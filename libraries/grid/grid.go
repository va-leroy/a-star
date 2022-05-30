package grid

import (
	"math"
)

type Position struct {
	X int // x coordinate
	Y int // y coordinate
}

type Grid struct {
	X, Y  int      // Grid size
	Value [][]int  // Grid values
	Mark  [][]int  // Grid marks
	Start Position // Start position
	End   Position // End position
}

type Node struct {
	Pos   Position // Position
	Cost  float64  // Cost
	Score float64  // Score = Cost + Heuristic
	Par   *Node    // Previous node (parent)
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
	g.X = x
	g.Y = y

	g.Value = make([][]int, x+1)
	for i := 0; i < x+1; i++ {
		g.Value[i] = make([]int, y+1)
		for j := 0; j < y+1; j++ {
			g.Value[i][j] = V_FREE
		}
	}

	g.Mark = make([][]int, x+1)
	for i := 0; i < x+1; i++ {
		g.Mark[i] = make([]int, y+1)
		for j := 0; j < y+1; j++ {
			g.Mark[i][j] = M_NULL
		}
	}

	g.Start = s
	g.End = e
	return &g
}

func (g *Grid) DestroyGrid() {
	g.Value = nil
	g.Mark = nil
}

func Heuristic(s, t Position) float64 {
	return math.Max(float64(t.X-s.X), float64(t.Y-s.Y))
}

func CreateNode(par *Node, c float64, pos Position, g *Grid) *Node {
	var n Node
	if par != nil {
		n.Par = par
	}
	n.Cost = c
	n.Pos = pos
	n.Score = Heuristic(g.Start, g.End)
	return &n
}

func CompareNode(a, b *Node) int {
	return int(a.Score) - int(b.Score)
}
