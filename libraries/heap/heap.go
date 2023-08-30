package heap

import (
	"a-star/libraries/grid"
	"fmt"
	"math"
)

type MinHeap struct {
	Array []grid.Node // Array of numbers that starts from 1 not 0
	N     int         // Amount of numbers in the heap
	Nmax  int         // Maximum amount of numbers in the heap
}

func HeapCreate(nmax int) *MinHeap {
	return &MinHeap{make([]grid.Node, nmax), 0, nmax} // Allocate memory
}

func HeapDestroy(h *MinHeap) {
	h.Array = nil // Free memory
}

func HeapEmpty(h *MinHeap) bool {
	return h.N == 0 && h.Array[1].Score == 0
}

func HeapTop(h *MinHeap) grid.Node {
	if HeapEmpty(h) {
		panic("Heap is empty")
	}
	return h.Array[1]
}

func HeapAdd(h *MinHeap, x grid.Node) {
	// If the heap is full, double its size
	if h.N+1 == h.Nmax {
		var newArray []grid.Node = make([]grid.Node, h.Nmax*2)
		for i := 0; i < h.Nmax; i++ {
			newArray[i] = h.Array[i]
		}
		h.Array = newArray
		h.Nmax *= 2
	}

	h.N++
	h.Array[h.N] = x
	var i int = h.N
	for i > 1 && h.Array[i/2].Score > h.Array[i].Score {
		h.Array[i/2], h.Array[i] = h.Array[i], h.Array[i/2] // Swap the parent and the child
		i /= 2
	}
}

func HeapPop(h *MinHeap) float64 {
	if HeapEmpty(h) {
		return -1
	}

	var x grid.Node = h.Array[1] // Save the top element to return it later
	h.Array[1] = h.Array[h.N]    // Move the last element to the top
	h.N--                        // Decrease the amount of elements in the heap

	var i int = 1     // Start from the top
	var l int = 2 * i // Left child
	var r int = l + 1 // Right child

	for l <= h.N && r <= h.N {
		if h.Array[l].Score < h.Array[r].Score {
			l = r
		}
		if h.Array[l].Score < h.Array[i].Score {
			h.Array[l], h.Array[i] = h.Array[i], h.Array[l]
			i = l
		} else {
			break
		}
		l = 2 * i
		r = l + 1
	}

	return x.Score
}

func HeapPrint(h *MinHeap) {
	fmt.Println()
	var height int = int(math.Log2(float64(h.N))) + 1 // Calculate the height of the heap
	for i := 1; i <= h.N; i++ {
		var level int = int(math.Log2(float64(i))) + 1 // Calculate the level of the current element
		var space int = (height - level + 1) * 2       // Calculate the amount of spaces to print
		fmt.Printf("%*s%.2f", space, "", h.Array[i].Score)
		if int(math.Pow(2, float64(level)))-1 == i {
			fmt.Println()
		}
	}
	fmt.Println()
}
