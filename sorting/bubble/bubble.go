package main

import (
	"fmt"
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

func main() {
	arr := []int{}
	Bubble(arr)
	fmt.Println(arr)
	arr = []int{12}
	Bubble(arr)
	fmt.Println(arr)
	arr = []int{12, 11}
	Bubble(arr)
	fmt.Println(arr)
	arr = []int{12, 23}
	Bubble(arr)
	fmt.Println(arr)
	arr = []int{12, 23, 3}
	Bubble(arr)
	fmt.Println(arr)
	arr = []int{12, 23, 3, 1, 34, 11}
	Bubble(arr)
	fmt.Println(arr)
}
