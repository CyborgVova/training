package main

import (
	"fmt"
)

type Tree struct {
	elem  int
	left  *Tree
	right *Tree
}

func FindMaxSum(root *Tree) int {
	if root == nil {
		return 0
	}

	left := FindMaxSum(root.left)
	right := FindMaxSum(root.right)
	return max(left, right) + root.elem
}

func main() {
	g := &Tree{elem: 4}
	h := &Tree{elem: 5}
	f := &Tree{elem: 2}
	d := &Tree{elem: 2}
	e := &Tree{elem: 3, left: f}
	b := &Tree{elem: 4, left: d, right: e}
	c := &Tree{elem: 7, left: g, right: h}
	a := &Tree{elem: 1, left: b, right: c}

	fmt.Println(FindMaxSum(a))
}
