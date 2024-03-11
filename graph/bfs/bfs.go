package bfs

import (
	"container/list"
	"errors"
	"github.com/cyborgvova/training/graph/matrix"
)

func BreadthFirstSearch(matrix matrix.Matrix, idx int) ([]int, error) {
	length := len(matrix.Values)
	if idx < 1 || idx > length {
		return nil, errors.New("nil matrix or vertex out of range")
	}

	if matrix.Rows != matrix.Cols {
		return nil, errors.New("none-square matrix passed")
	}

	result := []int{idx}
	idx--

	queue := list.New()
	queue.PushBack(idx)
	visited := map[int]struct{}{idx: {}}

	for queue.Front() != nil {
		for i := 0; i < length; i++ {
			if _, ok := visited[i]; ok || matrix.Values[idx][i] == 0 {
				continue
			}
			queue.PushBack(i)
			visited[i] = struct{}{}
			result = append(result, i+1)
		}
		queue.Remove(queue.Front())
		if queue.Front() != nil {
			idx = queue.Front().Value.(int)
		}
	}
	return result, nil
}
