package main

import (
	"fmt"
	"sort"
)

func main() {
	v := []int{3, 4, 1, 2, 5}
	ap(v)
	sr(v)
	fmt.Println(v) // [1, 2, 3, 4, 5]
}

func ap(arr []int) { // надо передавать через указатель (arr *[]int)
	arr = append(arr, 10)
	// если слайс переданный в функцию при модификации
	// превысит размер капасити изменения не коснутся
	// оригинального слайса
}

func sr(arr []int) {
	sort.Ints(arr)
}
