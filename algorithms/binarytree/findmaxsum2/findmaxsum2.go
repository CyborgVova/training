package main

import (
	"fmt"
)

var (
	answer = 0
)

type Tree struct {
	elem  int
	left  *Tree
	right *Tree
}

func FindMaxSum(root *Tree) {
	helper(root)
}

func helper(root *Tree) int {
	if root == nil {
		return 0
	}

	left := max(helper(root.left), 0)
	right := max(helper(root.right), 0)
	answer = max(answer, left+right+root.elem)
	return max(left, right) + root.elem
}

func main() {
	i := &Tree{elem: -2}
	f := &Tree{elem: 5, left: i}

	h := &Tree{elem: -2}
	e := &Tree{elem: 8, left: h}

	g := &Tree{elem: -4}
	d := &Tree{elem: -3, left: f, right: g}

	b := &Tree{elem: 9}
	c := &Tree{elem: 20, left: d, right: e}

	a := &Tree{elem: -10, left: b, right: c}
	FindMaxSum(a)
	fmt.Println(answer)
}
