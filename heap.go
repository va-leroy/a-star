package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type MinHeap struct {
	array []int // Array of numbers that starts from 1 not 0
	n     int   // Amount of numbers in the heap
	nmax  int   // Maximum amount of numbers in the heap
}

func heapCreate(nmax int) *MinHeap {
	return &MinHeap{make([]int, nmax), 0, nmax} // Allocate memory
}

func heapDestroy(h *MinHeap) {
	h.array = nil // Free memory
}

func heapEmpty(h *MinHeap) bool {
	return h.n == 0 && h.array[1] == 0
}

func heapTop(h *MinHeap) int {
	if heapEmpty(h) {
		return -1
	}
	return h.array[1]
}

func heapAdd(h *MinHeap, x int) {
	if h.n == h.nmax {
		fmt.Println("Heap is full")
		return
	}
	fmt.Println("Adding", x)
	h.n++
	h.array[h.n] = x
	var i int = h.n
	var par int = i / 2
	for par > 1 && h.array[par] > h.array[i] {
		h.array[par], h.array[i] = h.array[i], h.array[par]
		i = par
		par = i / 2
	}
}

func heapPop(h *MinHeap) int {
	if heapEmpty(h) {
		fmt.Println("Heap is empty")
		return -1
	}
	fmt.Println("Popping", h.array[1])

	var x int = h.array[1]    // Save the top element
	h.array[1] = h.array[h.n] // Move the last element to the top
	h.n--                     // Decrease the amount of elements in the heap

	var i int = 1     // Start from the top
	var l int = 2 * i // Left child
	var r int = l + 1 // Right child

	for l <= h.n && r <= h.n {
		if h.array[l] < h.array[r] {
			l = r
		}
		if h.array[l] < h.array[i] {
			h.array[l], h.array[i] = h.array[i], h.array[l]
			i = l
		} else {
			break
		}
		l = 2 * i
		r = l + 1
	}
	return x
}

func heapVerify(h *MinHeap) {
	for i := 1; i <= h.n; i++ {
		var l int = 2 * i
		var r int = l + 1
		if l <= h.n && h.array[l] < h.array[i] {
			fmt.Println("Heap is not valid")
			return
		}
		if r <= h.n && h.array[r] < h.array[i] {
			fmt.Println("Heap is not valid")
			return
		}
	}
	fmt.Println("Heap is valid")
}

func heapPrint(h *MinHeap) {
	var height int = int(math.Log2(float64(h.n))) + 1 // Calculate the height of the heap
	for i := 1; i <= h.n; i++ {
		var level int = int(math.Log2(float64(i))) + 1 // Calculate the level of the current element
		var space int = (height - level + 1) * 2       // Calculate the amount of spaces to print
		fmt.Printf("%*s%d", space, "", h.array[i])
		if int(math.Pow(2, float64(level)))-1 == i {
			fmt.Println()
		}
	}
	fmt.Println()
}

func main() {
	// Initialize random seed in Go
	rand.Seed(time.Now().UnixNano())

	// Create a heap of 10 elements (that can only fit 9 elements because of the 0 index)
	var h *MinHeap = heapCreate(10)

	// Add some random numbers to the heap
	for i := 1; i < 10; i++ {
		heapAdd(h, rand.Intn(100))
	}

	// Print the heap
	heapPrint(h)
}
