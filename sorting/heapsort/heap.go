package main

import (
	"fmt"
)

func heapify(arr []int, n, i int) {
	largest := i

	l := i*2 + 1
	r := i*2 + 2

	if l < n && arr[l] > arr[largest] {
		largest = l
	}

	if r < n && arr[r] > arr[largest] {
		largest = r
	}

	if largest != i {
		arr[largest], arr[i] = arr[i], arr[largest]
		heapify(arr, n, largest)
	}
}

func BuildHeap(arr []int, n int) {
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
}

func SortHeap(arr []int, n int) {
	BuildHeap(arr, n)
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func main() {
	var arr = []int{3, 6, 7, 8, 24, 65, 11, 34, 9}

	// heapify(arr, len(arr), 0)
	// BuildHeap(arr, len(arr))
	SortHeap(arr, len(arr))
	fmt.Println(arr)
}
