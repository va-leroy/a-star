package main

import (
	"a-star/libraries/grid"
	"a-star/libraries/heap"
	"fmt"
	"math/rand"
	"time"
)

func initializeGrid(N, W int) *grid.Grid {
	/* Create start and end positions */
	var s, e grid.Position
	s.X, s.Y = 0, 0
	e.X, e.Y = N-1, N-1

	/* Create grid */
	var G *grid.Grid = grid.CreateGrid(N, N, s, e)

	/* Create walls */
	for i := 0; i < W; i++ {
		var x, y int
		x = rand.Intn(N)
		y = rand.Intn(N)

		if G.Value[x][y] != grid.V_WALL && (x != s.X || y != s.Y) && (x != e.X || y != e.Y) {
			G.Value[x][y] = grid.V_WALL // A wall can't exist on the start or end position
		} else {
			i--
		}
	}

	/* Return grid */
	return G
}

func main() {
	/* Timer and random seed */
	timer := time.Now()
	rand.Seed(time.Now().UnixNano())

	/* Initialize the grid */
	var N int = 20
	var W int = 50
	var G *grid.Grid = initializeGrid(N, W)

	/* Print grid information */
	fmt.Printf("Grid size  : %dx%d\n", N, N)
	fmt.Printf("Wall units : %d\n", W)
	fmt.Println("Start      :", G.Start)
	fmt.Println("End        :", G.End)
	fmt.Println("")

	/* Create heap Q */
	N = N * N                        // Number of nodes
	var size int = (N * (N - 1)) / 2 // Number of edges
	var Q *heap.MinHeap = heap.HeapCreate(size)

	/*                     */
	/* ACTUAL A* ALGORITHM */
	/*                     */

	var start = grid.CreateNode(nil, 0, G.Start, G) // Add the start node to the heap
	heap.HeapAdd(Q, *start)

	// While heap Q isn't empty
	for !heap.HeapEmpty(Q) {
		var U grid.Node = heap.HeapTop(Q) // Get the node with the lowest score
		heap.HeapPop(Q)

		if U.Pos.X == G.End.X && U.Pos.Y == G.End.Y {
			// If U is the end position then we found the path
			for U.Par != nil {
				G.Mark[U.Pos.X][U.Pos.Y] = grid.M_PATH
				U = *U.Par
			}

			G.Mark[U.Pos.X][U.Pos.Y] = grid.M_PATH // Mark start position
			heap.HeapDestroy(Q)                    // Destroy the heap
			grid.PrintGrid(G)                      // Print the grid to the console
			G.DestroyGrid()                        // Destroy the grid
			fmt.Println("Time elapsed:", time.Since(timer))
			return
		}

		G.Mark[U.Pos.X][U.Pos.Y] = grid.M_USED // Mark U as used

		// Loop through the neighbors of U
		for i := 0; i < 4; i++ {
			var n grid.Position
			n.X = U.Pos.X
			n.Y = U.Pos.Y
			switch i {
			case 0:
				n.X = n.X - 1 // Left
			case 1:
				n.X = n.X + 1 // Right
			case 2:
				n.Y = n.Y - 1 // Up
			case 3:
				n.Y = n.Y + 1 // Down
			}
			// If this neighbor is not a wall, not used and not in the heap, add it to the heap
			if grid.IsInGrid(G, n) && G.Value[n.X][n.Y] != grid.V_WALL && G.Mark[n.X][n.Y] != grid.M_USED && G.Mark[n.X][n.Y] != grid.M_FRONT {
				var v *grid.Node = grid.CreateNode(&U, U.Cost, n, G)
				heap.HeapAdd(Q, *v)
			}
		}
	}

	heap.HeapDestroy(Q) // Destroy the heap
	G.DestroyGrid()     // Destroy the grid
	fmt.Println("No path found")
}
