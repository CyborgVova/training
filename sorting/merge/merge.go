package main

import (
	"fmt"
)

// merge getting two parts of slice and merging it simple comparison
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

// MergeSort dividing slice on two part and put it in merge function
func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left := MergeSort(arr[:len(arr)/2])
	right := MergeSort(arr[len(arr)/2:])

	return merge(left, right)
}

func main() {
	arr := []int{}
	fmt.Println(MergeSort(arr))
	arr = []int{5}
	fmt.Println(MergeSort(arr))
	arr = []int{1, 1}
	fmt.Println(MergeSort(arr))
	arr = []int{1, 2}
	fmt.Println(MergeSort(arr))
	arr = []int{5, 2}
	fmt.Println(MergeSort(arr))
	arr = []int{1, 3, 5, 2, 4}
	fmt.Println(MergeSort(arr))
}
