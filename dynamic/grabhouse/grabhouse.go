package main

import (
	"fmt"
)

func GrabHouse(arr []int) int {
	len := len(arr)
	if len == 0 {
		return 0
	}

	if len == 1 {
		return arr[1]
	}

	set := make([]int, len)
	set[0] = arr[0] // 4
	set[1] = arr[1] // 11

	for i := 2; i < len; i++ {
		set[i] = max(set[i-1], set[i-2]+arr[i])
	}

	return set[len-1]
}

// 4, 11, 14, 14, 16, 22, 22

func main() {
	fmt.Println(GrabHouse([]int{4, 11, 10, 1, 2, 8, 5})) // 22
}
