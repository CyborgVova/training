package main

import (
	"fmt"
)

// Intersection finding intersection values into two slices.
func Intersection(a, b []int) []int {
	m := make(map[int]int)
	result := []int{}
	for _, ela := range a {
		m[ela]++
	}

	for _, elb := range b {
		if count, ok := m[elb]; ok && count > 0 {
			result = append(result, elb)
			m[elb]--
		}
	}
	return result
}

func main() {
	a := []int{23, 3, 1, 2}
	b := []int{6, 2, 4, 23} // [2, 23]
	fmt.Printf("%v\n", Intersection(a, b))
	a = []int{1, 1, 1}
	b = []int{1, 1, 1, 1} // [1, 1, 1]
	fmt.Printf("%v\n", Intersection(a, b))
	a = []int{1, 1, 1}
	b = []int{3, 1, 4, 5} // [1]
	fmt.Printf("%v\n", Intersection(a, b))
}
