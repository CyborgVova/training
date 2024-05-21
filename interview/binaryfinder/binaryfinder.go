package main

import (
	"fmt"
)

func BinaryFinder(arr []int, elem int) (int, bool) {
	if len(arr) == 0 {
		return 0, false
	}

	start, end := 0, len(arr)-1

	for start <= end {
		mid := (end-start)/2 + start

		if arr[mid] == elem {
			return mid, true
		} else if arr[mid] < elem {
			start = mid + 1
		} else if arr[mid] > elem {
			end = mid - 1
		}
	}

	return 0, false
}

func main() {
	arr := []int{1, 3, 4, 6, 8, 10, 55, 56, 59, 70, 79, 81, 91, 10001}
	fmt.Println(BinaryFinder(arr, 55))
}
