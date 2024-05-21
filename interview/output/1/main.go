package main

import (
	"fmt"
)

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr1 = arr1[:3]
	arr2 := arr1[3:]
	fmt.Println(len(arr2), cap(arr2)) // len = 0, cap = 2
	fmt.Println(arr2)                 // []
	fmt.Println(arr2[:2])             // [4, 5]
}
