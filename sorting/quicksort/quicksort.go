package main

import (
	"fmt"
)

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

func main() {
	arr := []int{}
	arr = QuickSort(arr)
	fmt.Println(arr)
	arr = []int{12}
	arr = QuickSort(arr)
	fmt.Println(arr)
	arr = []int{12, 11}
	arr = QuickSort(arr)
	fmt.Println(arr)
	arr = []int{12, 23}
	arr = QuickSort(arr)
	fmt.Println(arr)
	arr = []int{12, 23, 3}
	arr = QuickSort(arr)
	fmt.Println(arr)
	arr = []int{12, 23, 3, 1, 34, 11}
	arr = QuickSort(arr)
	fmt.Println(arr)
}
