package main

import (
	"fmt"
)

func merge(a, b []int) (res []int) {
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}

	if i < len(a) {
		res = append(res, a[i:]...)
	}

	if j < len(b) {
		res = append(res, b[j:]...)
	}

	return
}

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	a := MergeSort(arr[:len(arr)/ 2])
	b := MergeSort(arr[len(arr)/ 2:])

	return merge(a, b)
}

func heapify(arr []int, len, i int) {
	largest := i

	l := i*2 + 1
	r := i*2 + 2

	if l < len && arr[l] > arr[largest] {
		largest = l
	}

	if r < len && arr[r] > arr[largest] {
		largest = r
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, len, largest)
	}
}

func HeapSort(arr []int, len int) {
	for i := len/2 -1; i >= 0; i-- {
		heapify(arr, len, i)
	}

	for i := len - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func main() {
	arr := []int{6, 11, 13, 17, 1, 3, 7, 2, 4}

	HeapSort(arr, len(arr))
	fmt.Println(arr)

	arr = []int{6, 11, 13, 17, 1, 3, 7, 2, 4}
	fmt.Println(MergeSort(arr))
}
