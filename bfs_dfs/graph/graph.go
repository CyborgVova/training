package main

import (
	"errors"
	"fmt"
)

func bfsGraph(matr [][]int, idx int) (result []int, err error) {
	if matr == nil {
		return result, errors.New("empty matrix")
	} else if len(matr[0]) != len(matr) {
		return result, errors.New("matrix not square")
	} else if idx > len(matr[0])-1 {
		return result, errors.New("index out of range")
	}

	result = append(result, idx)
	queue := []int{idx}
	visited := map[int]struct{}{idx: {}}

	for len(queue) > 0 {
		for i := 0; i < len(matr[0]); i++ {
			if _, ok := visited[i]; !ok && matr[queue[0]][i] != 0 {
				visited[i] = struct{}{}
				queue = append(queue, i)
				result = append(result, i)
			}
		}

		if len(queue) == 1 {
			return
		}
		if len(queue) > 1 {
			queue = queue[1:]
		}
	}

	return
}

func dfsGraph(matr [][]int, idx int) (result []int, err error) {
	if matr == nil {
		return result, errors.New("empty matrix")
	} else if len(matr[0]) != len(matr) {
		return result, errors.New("matrix not square")
	} else if idx > len(matr[0])-1 {
		return result, errors.New("index out of range")
	}

	visited := map[int]struct{}{idx: {}}
	result = append(result, idx)
	stack := []int{idx}

	for len(stack) > 0 {
		have_unvisited := false
		for i := 0; i < len(matr); i++ {
			if _, ok := visited[i]; !ok && matr[stack[len(stack)-1]][i] != 0 {
				result = append(result, i)
				visited[i] = struct{}{}
				stack = append(stack, i)
				have_unvisited = true
				break
			}
		}

		if len(stack) > 0 && !have_unvisited {
			stack = stack[:len(stack)-1]
		}
	}
	return
}

func main() {
	idx := 0
	matr := [][]int{
		{0, 1, 0, 1, 1},
		{1, 0, 1, 0, 1},
		{1, 1, 0, 1, 1},
		{0, 1, 1, 0, 1},
		{1, 1, 1, 1, 0},
	}

	fmt.Println(dfsGraph(matr, idx))
	fmt.Println(bfsGraph(matr, idx))
	fmt.Println(dfsGraph(nil, idx))
	fmt.Println(bfsGraph(nil, idx))
	fmt.Println(dfsGraph(matr, 6))
	fmt.Println(bfsGraph(matr, 6))
	fmt.Println(bfsGraph([][]int{{0, 1, 1}, {1, 0, 1}}, idx))
}
