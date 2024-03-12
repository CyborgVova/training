package main

import (
	"fmt"
)

func pointerSlice(s *[]int32) {
	// Изменение сохранится так как работа происходит с указателем на слайс
	*s = append(*s, 123)

}

func valueSlice(s []int32) {
	// Изменение сохранится так как не происходит изменения размера
	s[0] = 777
	// Изменение не сохранится так как изменится размер который передается как значение
	s = append(s, 123)
}

func main() {
	// var k = []int{} // k != nil
	// var к = make([]int, 0) // k != nil

	k := make([]int, 0) // k != nil
	if k != nil {
		fmt.Println(k) // []
	}

	// Объявление слайса
	var s []int // s == nil
	if s == nil {
		fmt.Println(s) // []
	}

	// Это не мешает заполнять слайс
	s = append(s, 200)
	fmt.Println(s, len(s), cap(s))

	var sl = []int32{1, 2, 3}
	fmt.Println("sl valueSlice", sl, cap(sl))
	valueSlice(sl)
	fmt.Println(sl, cap(sl))

	fmt.Println("sl pointerSlice", sl, cap(sl))
	pointerSlice(&sl)
	fmt.Println(sl, cap(sl))
}
