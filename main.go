package main

import (
	"a-star/libraries/grid"
	"a-star/libraries/heap"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Initialize random seed in Go
	rand.Seed(time.Now().UnixNano())

	// Create two positions
	var s, e grid.Position

	// Set the start position
	s.X = 0
	s.Y = 0

	// Set the end position
	e.X = 100
	e.Y = 100

	// Declare size of the grid
	var N int = 100

	// Create a grid
	var G *grid.Grid = grid.CreateGrid(N, N, s, e)

	// Create 10 random walls
	for i := 0; i < 10; i++ {
		var x, y int
		x = rand.Intn(10)
		y = rand.Intn(10)
		G.Value[x][y] = grid.V_WALL
	}

	N = N * N                        // Number of nodes
	var size int = (N * (N - 1)) / 2 // Number of edges

	var Q *heap.MinHeap = heap.HeapCreate(size) // Create a heap

	// Create node from start position
	var start = grid.CreateNode(nil, 0, G.Start, G)
	heap.HeapAdd(Q, *start)

	// While heap Q isn't empty
	for !heap.HeapEmpty(Q) {
		// Get the node with the lowest score
		var U grid.Node = heap.HeapTop(Q)
		heap.HeapPop(Q)

		// If u is equal to end position
		if U.Pos.X == G.End.X && U.Pos.Y == G.End.Y {
			fmt.Println("[!] Path found")
			// Return path from u to start
			for U.Par != nil {
				G.Mark[U.Pos.X][U.Pos.Y] = grid.M_PATH
				U = *U.Par
				fmt.Println("Coordinates: ", U.Pos.X, U.Pos.Y)
			}
			G.Mark[U.Pos.X][U.Pos.Y] = grid.M_PATH // Mark start position
			heap.HeapDestroy(Q)                    // Destroy the heap
			grid.PrintGrid(G)                      // Print the grid
			return
		}

		G.Mark[U.Pos.X][U.Pos.Y] = grid.M_USED // Mark U as used

		var dx = []int{-1, 1, 1, -1, -1, 0, 2, 0}
		var dy = []int{0, 1, -1, -1, 0, 2, 0, -2}

		for i := 0; i < 8; i++ {
			var pos_v grid.Position
			pos_v.X = U.Pos.X + dx[i]
			pos_v.Y = U.Pos.Y + dy[i]

			if grid.IsInGrid(G, pos_v) {
				if G.Value[pos_v.X][pos_v.Y] != grid.V_WALL {
					if G.Mark[pos_v.X][pos_v.Y] != grid.M_USED {
						if G.Mark[pos_v.X][pos_v.Y] != grid.M_FRONT {
							var v *grid.Node = grid.CreateNode(&U, U.Cost, pos_v, G)
							heap.HeapAdd(Q, *v)
						}
					}
				}
			}
		}
	}

	heap.HeapDestroy(Q)          // Destroy the heap
	fmt.Println("No path found") // No return instruction because it's redundant
}
