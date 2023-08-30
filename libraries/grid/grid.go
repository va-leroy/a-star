package grid

import (
	"math"
)

type (
	Position struct {
		X int // x coordinate
		Y int // y coordinate
	}
	Grid struct {
		X, Y  int      // Grid size
		Value [][]int  // Grid values
		Mark  [][]int  // Grid marks
		Start Position // Start position
		End   Position // End position
	}
	Node struct {
		Pos   Position // Position
		Cost  float64  // Cost
		Score float64  // Score equals to cost + heuristic
		Par   *Node    // Parent
	}
)

const (
	V_FREE = iota
	V_WALL
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

func IsInGrid(g *Grid, pos Position) bool {
	return pos.X >= 0 && pos.X <= g.X && pos.Y >= 0 && pos.Y <= g.Y
}

func PrintGrid(g *Grid) {
	var s string
	for i := 0; i < g.X+1; i++ {
		for j := 0; j < g.Y+1; j++ {
			// If it's starting position then add s to the string
			if i == g.Start.X && j == g.Start.Y {
				s += "s"
			} else if i == g.End.X && j == g.End.Y {
				s += "e"
			} else {
				if g.Value[i][j] == V_WALL {
					s += "#"
				} else if g.Mark[i][j] == M_PATH {
					s += "."
				} else {
					s += " "
				}
			}
		}
		s += "\n"
	}
	print(s)
}

func HeuristicEuclidean(s, t Position) float64 {
	return math.Sqrt(math.Pow(float64(s.X-t.X), 2) + math.Pow(float64(s.Y-t.Y), 2))
}

func HeuristicManhattan(s, t Position) float64 {
	return math.Abs(float64(s.X-t.X)) + math.Abs(float64(s.Y-t.Y))
}

func CreateNode(par *Node, cost float64, pos Position, g *Grid) *Node {
	var n Node
	if par != nil {
		n.Par = par
	}
	n.Cost = cost
	n.Pos = pos
	n.Score = n.Cost + HeuristicEuclidean(n.Pos, g.End)
	return &n
}
