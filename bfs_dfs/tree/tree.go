package main

import (
	"errors"
	"fmt"
)

type Tree struct {
	Left  *Tree
	Right *Tree
	Value int
}

var (
	h Tree = Tree{nil, nil, 65}
	i Tree = Tree{nil, nil, 11}

	f Tree = Tree{nil, nil, 34}
	g Tree = Tree{nil, nil, 9}

	d Tree = Tree{&f, &g, 8}
	e Tree = Tree{nil, nil, 24}

	b Tree = Tree{&d, &e, 6}
	c Tree = Tree{&h, &i, 7}

	a Tree = Tree{&b, &c, 3}

	/*
		        3
		      /   \
			 6     7
			/ \   / \
		   8  24 65  11
		  / \
		 34  9
		depth	= 3, 6, 8, 34, 9, 24, 7, 65, 11 +++++++
		breadth = 3, 6, 7, 8, 24, 65, 11, 34, 9 +++++++
	*/
)

func bfsTree(root *Tree) (result []int, err error) {
	if root == nil {
		return result, errors.New("empty tree")
	}

	result = append(result, root.Value)
	visited := map[*Tree]struct{}{}
	queue := []*Tree{root}

	for len(queue) > 0 {
		if root.Left != nil {
			if _, ok := visited[root.Left]; !ok {
				queue = append(queue, root.Left)
				visited[root.Left] = struct{}{}
				result = append(result, root.Left.Value)
			}
		}

		if root.Right != nil {
			if _, ok := visited[root.Right]; !ok {
				queue = append(queue, root.Right)
				visited[root.Right] = struct{}{}
				result = append(result, root.Right.Value)
			}
		}

		if len(queue) == 1 {
			queue = []*Tree{}
		}

		if len(queue) > 1 {
			queue = queue[1:]
			root = queue[0]
		}
	}
	return
}

func dfsTree(root *Tree) (result []int, err error) {
	if root == nil {
		return result, errors.New("empty tree")
	}
	result = append(result, root.Value)
	visited := map[*Tree]struct{}{root: {}}
	stack := []*Tree{root}

	for len(stack) > 0 {
		have_unvisited := false
		if root.Left != nil {
			if _, ok := visited[root.Left]; !ok {
				root = root.Left
				visited[root] = struct{}{}
				stack = append(stack, root)
				result = append(result, root.Value)
				have_unvisited = true
				continue
			}
		}

		if root.Right != nil {
			if _, ok := visited[root.Right]; !ok {
				root = root.Right
				visited[root] = struct{}{}
				stack = append(stack, root)
				result = append(result, root.Value)
				have_unvisited = true
			}
		}

		if !have_unvisited {
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				root = stack[len(stack)-1]
			}
		}
	}
	return
}

func main() {
	fmt.Println(dfsTree(&a))
	fmt.Println(bfsTree(&a))
	fmt.Println(dfsTree(nil))
	fmt.Println(bfsTree(nil))
}
