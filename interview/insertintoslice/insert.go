package main

import (
	"errors"
	"fmt"
)

func InsertIntoSlice(arr []int, idx, value int) ([]int, error) {
	if idx < 0 || idx > len(arr) {
		return arr, errors.New("index out of range")
	}
	return append(arr[:idx], append([]int{value}, arr[idx:]...)...), nil
}

func InsertIntoSliceWithoutRelocation(arr []int, value, idx int) ([]int, error) {
	if idx < 0 || idx > len(arr) {
		return arr, errors.New("index out of range")
	}
	result := make([]int, 0, len(arr)+1)
	result = append(result, arr[:idx]...)
	result = append(result, value)
	return append(result, arr[idx:]...), nil
}

func main() {
	fmt.Println(InsertIntoSlice([]int{0, 1, 2, 4}, -1, 3))
	fmt.Println(InsertIntoSlice([]int{0, 1, 2, 4}, 5, 3))
	fmt.Println(InsertIntoSlice([]int{0, 1, 2, 4}, 3, 3))
	fmt.Println(InsertIntoSlice([]int{0, 1, 2, 3}, 0, -1))
	fmt.Println(InsertIntoSlice([]int{0, 1, 2, 3}, 4, 4))
	fmt.Println(InsertIntoSliceWithoutRelocation([]int{0, 1, 2, 3}, 4, 4))
	x, _ := InsertIntoSlice([]int{0, 1, 2, 3}, 4, 4)
	fmt.Println(cap(x))
	x, _ = InsertIntoSliceWithoutRelocation([]int{0, 1, 2, 3}, 4, 4)
	fmt.Println(cap(x))
}
