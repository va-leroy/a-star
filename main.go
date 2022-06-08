/*
 * Copyright (c) 2022, Valentin Leroy
 * All rights reserved.
 */
package main

import (
	"a-star/libraries/grid"
	"a-star/libraries/heap"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Start timer for benchmarking
	timer := time.Now()

	// Initialize random seed in Go
	rand.Seed(time.Now().UnixNano())

	// Create two positions
	var s, e grid.Position

	// Set the start position
	s.X = 0
	s.Y = 0

	// Set the end position
	e.X = 20
	e.Y = 20

	// Declare size of the grid
	var N int = 20
	fmt.Println("[*] Grid size:", N)

	// Create a grid
	var G *grid.Grid = grid.CreateGrid(N, N, s, e)

	// Set the number of walls
	var walls int = 50

	// Create 50 random walls
	for i := 0; i < walls; i++ {
		var x, y int
		x = rand.Intn(N)
		y = rand.Intn(N)
		G.Value[x][y] = grid.V_WALL
	}
	fmt.Println("[*] Number of walls:", walls)

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
			}

			G.Mark[U.Pos.X][U.Pos.Y] = grid.M_PATH // Mark start position
			heap.HeapDestroy(Q)                    // Destroy the heap
			grid.PrintGrid(G)                      // Print the grid

			fmt.Println("[!] Time elapsed:", time.Since(timer))
			return
		}

		G.Mark[U.Pos.X][U.Pos.Y] = grid.M_USED // Mark U as used

		// For each four neighbors of U
		for i := 0; i < 4; i++ {
			// Get the neighbor position
			var n grid.Position
			n.X = U.Pos.X
			n.Y = U.Pos.Y
			switch i {
			case 0:
				n.X = n.X - 1
			case 1:
				n.X = n.X + 1
			case 2:
				n.Y = n.Y - 1
			case 3:
				n.Y = n.Y + 1
			}

			// If the neighbor is in the grid
			if grid.IsInGrid(G, n) {
				// If the neighbor is not a wall
				if G.Value[n.X][n.Y] != grid.V_WALL {
					// If the neighbor is not used
					if G.Mark[n.X][n.Y] != grid.M_USED {
						// If the neighbor is not in the heap
						if G.Mark[n.X][n.Y] != grid.M_FRONT {
							var v *grid.Node = grid.CreateNode(&U, U.Cost, n, G)
							heap.HeapAdd(Q, *v)
						}
					}
				}
			}
		}
	}

	heap.HeapDestroy(Q)              // Destroy the heap
	fmt.Println("[!] No path found") // No return instruction because it's redundant
}
