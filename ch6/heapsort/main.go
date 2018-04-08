package main

import (
	"fmt"
	"log"
)

// Heap : heap
type Heap struct {
	Array []int
	Size  int
}

// LeftChild : get the left child of element i
func LeftChild(a []int, i int) int {
	return (len(a) / 2)
}

// RightChild : get the right child of element i
func RightChild(a []int, i int) int {
	return (len(a) / 2) + 1
}

// MaxHeapify : maximizing heap
func MaxHeapify(heap *Heap, i int) {
	largest := i
	li := LeftChild(heap.Array, i)
	ri := RightChild(heap.Array, i)
	log.Println("i => ", i, " li => ", li, " ri => ", ri)
	if heap.Size >= li && heap.Array[li] > heap.Array[largest] {
		largest = li
	}
	if heap.Size >= ri && heap.Array[ri] > heap.Array[largest] {
		largest = ri
	}
	if largest != i {
		tmp := heap.Array[i]
		heap.Array[i] = heap.Array[largest]
		heap.Array[largest] = tmp
		MaxHeapify(heap, largest)
	}
	log.Println("maximized heap => ", heap.Array)
}

// BuildMaxHeapify : build max heap
func BuildMaxHeapify(heap *Heap) {
	heap.Size = len(heap.Array)
	for i := len(heap.Array)/2 - 1; i >= 0; i-- {
		MaxHeapify(heap, i)
	}
}

// HeapSort : sort the array a by incresing
func HeapSort(heap *Heap) {
	BuildMaxHeapify(heap)
	for i := len(heap.Array) - 1; i > 0; i-- {
		tmp := heap.Array[0]
		heap.Array[0] = heap.Array[i]
		heap.Array[i] = tmp
		heap.Size = heap.Size - 1
		MaxHeapify(heap, i)
	}
}

func main() {
	heap := &Heap{
		Array: []int{12, 44, 4, -3, 4, 8, 3, 23, 21, 9, 11},
	}
	HeapSort(heap)
	fmt.Println("Running...")
	fmt.Println("After sorted => ", heap.Array)
}
