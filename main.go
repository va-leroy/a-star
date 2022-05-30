package main

import (
	"a-star/libraries/grid"
	"a-star/libraries/heap"
	"math/rand"
	"time"
)

func main() {
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
