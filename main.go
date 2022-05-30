package main

import (
	"a-star/libraries/grid"
	"a-star/libraries/heap"
	"math/rand"
	"time"
)

func UsingHeap() {
	// Initialize random seed in Go
	rand.Seed(time.Now().UnixNano())

	// Create a heap of 10 elements (that can only fit 9 elements because of the 0 index)
	var h *heap.MinHeap = heap.HeapCreate(5)

	// Add 5 random numbers to the heap
	for i := 0; i < 4; i++ {
		var x grid.Node
		x.Pos.X = rand.Intn(100)
		x.Pos.Y = rand.Intn(100)
		x.Cost = float64(rand.Intn(100))
		x.Score = x.Cost + float64(rand.Intn(100))
		heap.HeapAdd(h, x)
	}

	// Print the heap
	heap.HeapPrint(h)
}

func main() {
	// Create two positions
	var s, e grid.Position
	// Set the start position
	s.X = 0
	s.Y = 0
	// Set the end position
	e.X = 10
	e.Y = 10

	// Create a grid
	var g *grid.Grid = grid.CreateGrid(10, 10, s, e)

	// Create 10 random walls
	for i := 0; i < 10; i++ {
		var x, y int
		x = rand.Intn(10)
		y = rand.Intn(10)
		g.Value[x][y] = grid.V_WALL
	}

	// Save the grid to a file
	grid.PrintGrid(g)
}
