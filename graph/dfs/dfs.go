package dfs

import (
	"container/list"
	"errors"

	"github.com/cyborgvova/training/graph/matrix"
)

func DepthFirstSearch(matrix matrix.Matrix, idx int) ([]int, error) {
	length := len(matrix.Values)
	if idx < 1 || idx > length {
		return nil, errors.New("nil matrix or vertex out of range")
	}

	if matrix.Rows != matrix.Cols {
		return nil, errors.New("none-square matrix passed")
	}

	result := []int{idx}
	idx--

	stack := list.New()
	stack.PushBack(idx)
	visited := map[int]struct{}{idx: {}}

	for stack.Back() != nil {
		have_unvisited := false
		for i := 0; i < length; i++ {
			if _, ok := visited[i]; !ok && matrix.Values[idx][i] != 0 {
				stack.PushBack(i)
				visited[i] = struct{}{}
				result = append(result, i+1)
				have_unvisited = true
				idx = i
				break
			}
		}

		if !have_unvisited {
			stack.Remove(stack.Back())
			if stack.Back() != nil {
				idx = stack.Back().Value.(int)
			}
		}
	}
	return result, nil
}
