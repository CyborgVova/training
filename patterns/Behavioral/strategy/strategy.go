package strategy

type Strategy interface {
	Sort([]int)
}

type HeapSort struct{}

func (s *HeapSort) Sort(arr []int) {
	len := len(arr)

	for i := len/2 - 1; i >= 0; i-- {
		MakeHeap(arr, len, i)
	}

	for i := len - 1; i >= 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		MakeHeap(arr, i, 0)
	}
}

func MakeHeap(arr []int, len, root int) {
	largest := root
	left := root*2 + 1
	right := root*2 + 2

	if left < len && arr[left] > arr[largest] {
		largest = left
	}

	if right < len && arr[right] > arr[largest] {
		largest = right
	}

	if largest != root {
		arr[root], arr[largest] = arr[largest], arr[root]
		MakeHeap(arr, len, largest)
	}
}

type BubbleSort struct{}

func (sort *BubbleSort) Sort(arr []int) {
	len := len(arr)
	if len < 2 {
		return
	}

	for i := len - 1; i > 0; i-- {
		max := arr[i]
		for j := 0; j < i; j++ {
			if arr[j] > max {
				max = arr[j]
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

type InsertSort struct{}

func (sort *InsertSort) Sort(arr []int) {
	for i := 1; i < len(arr); i++ {
		x := arr[i]
		j := i - 1
		for ; j >= 0 && x < arr[j]; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = x
	}
}

type Object struct {
	Strategy Strategy
}

func NewObject(strategy Strategy) *Object {
	return &Object{Strategy: strategy}
}

func (o *Object) Toggler(strategy Strategy) {
	o.Strategy = strategy
}

func (o *Object) Sort(arr []int) {
	o.Strategy.Sort(arr)
}
