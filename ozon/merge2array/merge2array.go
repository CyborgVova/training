package main

import "fmt"

func Merge(a, b []int) []int {
	l1 := len(a)
	l2 := len(b)
	out := make([]int, 0, l1+l2)
	i, j := 0, 0
	for i < l1 && j < l2 {
		if a[i] < b[j] {
			out = append(out, a[i])
			i++
		} else {
			out = append(out, b[j])
			j++
		}
	}
	out = append(out, a[i:]...)
	out = append(out, a[j:]...)
	return out
}

func main() {
	a := []int{1, 4, 6, 7, 8, 54}
	b := []int{1, 2, 3, 5, 9, 33, 100}
	fmt.Println(Merge(a, b))
}
