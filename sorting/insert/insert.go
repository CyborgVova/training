package main

import (
	"fmt"
)

func Insertion(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]                // сохраняем очередной элемент вне слайса
		j := i - 1                   // устанавливаем начало поиска с предыдущего элемента
		for j >= 0 && arr[j] > key { // двигаемся к началу слайса и проверяем
			arr[j+1] = arr[j] // если очередной элемент больше текущего
			//	записываем в текущий индекс предыдущий элемент
			// на этом этапе два соседних элемента одинаковы
			j--
		}
		arr[j+1] = key // после того как нашлось место для сохраненного элемента
		// записываем его в это место [j + 1] так как в цикле было j--
	}
}

// 12, 23, 3, 1, 34, 11

func main() {
	arr := []int{}
	Insertion(arr)
	fmt.Println(arr)
	arr = []int{12}
	Insertion(arr)
	fmt.Println(arr)
	arr = []int{12, 11}
	Insertion(arr)
	fmt.Println(arr)
	arr = []int{12, 23}
	Insertion(arr)
	fmt.Println(arr)
	arr = []int{12, 23, 3}
	Insertion(arr)
	fmt.Println(arr)
	arr = []int{12, 23, 3, 1, 34, 11}
	Insertion(arr)
	fmt.Println(arr)
}
