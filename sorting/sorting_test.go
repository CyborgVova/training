package tests

import (
	"testing"
)

func Bubble(arr []int) {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		for j := 1; j < l-i; j++ {
			if arr[j-1] > arr[j] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

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
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func Insertion(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func merge(left, right []int) []int {
	var out []int
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			out = append(out, left[i])
			i++
		} else {
			out = append(out, right[j])
			j++
		}
	}

	for i < len(left) {
		out = append(out, left[i])
		i++
	}

	for j < len(right) {
		out = append(out, right[j])
		j++
	}

	return out
}

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left := MergeSort(arr[:len(arr)/2])
	right := MergeSort(arr[len(arr)/2:])

	return merge(left, right)
}

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var left, right []int
	for i := 1; i < len(arr); i++ {
		if pivot >= arr[i] {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	result := append(QuickSort(left), pivot)
	result = append(result, QuickSort(right)...)
	return result
}

func BenchmarkBubble(b *testing.B) {
	arr := []int{1, 3, 5, 2, 6, 4}
	for i := 0; i < b.N; i++ {
		Bubble(arr)
	}
}

func BenchmarkHeapSort(b *testing.B) {
	arr := []int{1, 3, 5, 2, 6, 4}
	for i := 0; i < b.N; i++ {
		SortHeap(arr, 0)
	}
}

func BenchmarkInsert(b *testing.B) {
	arr := []int{1, 3, 5, 2, 6, 4}
	for i := 0; i < b.N; i++ {
		Insertion(arr)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	arr := []int{1, 3, 5, 2, 6, 4}
	for i := 0; i < b.N; i++ {
		QuickSort(arr)
	}
}

func BenchmarkMerge(b *testing.B) {
	arr := []int{1, 3, 5, 2, 6, 4}
	for i := 0; i < b.N; i++ {
		MergeSort(arr)
	}
}
