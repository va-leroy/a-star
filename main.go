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
	e.X = 10
	e.Y = 10

	// Declare size of the grid
	var N int = 10

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
			fmt.Println("Path found")
			// Return path from u to start
			for U.Par != nil {
				G.Mark[U.Pos.X][U.Pos.Y] = grid.M_PATH
				U = *U.Par
			}
			G.Mark[U.Pos.X][U.Pos.Y] = grid.M_PATH // Mark start position
			heap.HeapDestroy(Q)                    // Destroy the heap
			grid.PrintGrid(G)                      // Print the grid
		}

		G.Mark[U.Pos.X][U.Pos.Y] = grid.M_USED // Mark U as used

		// Get the neighbors of U
		// TODO
		grid.PrintGrid(G)
	}

	heap.HeapDestroy(Q) // Destroy the heap
	fmt.Println("No path found")
}
