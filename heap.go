package main

import "fmt"

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
	for i := 1; i <= h.n; i++ {
		fmt.Print(h.array[i], " ")
	}
	fmt.Println()
}

func main() {
	// Create a heap of 10 elements
	var h *MinHeap = heapCreate(10)

	// Add elements to the heap
	heapAdd(h, 12)
	heapAdd(h, 31)
	heapAdd(h, 67)
	heapAdd(h, 24)

	// Print the heap
	heapPrint(h)
}
