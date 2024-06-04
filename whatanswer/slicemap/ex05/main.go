package main

import (
	"fmt"
)

func main() {
	mutate := func(a []int) {
		a[0] = 0         // пока не менялось капасити меняем элемент по индексу
		a = append(a, 1) // создается новый слайс
		fmt.Println(a)   // [0, 2, 3, 4, 1]
	}
	a := []int{1, 2, 3, 4}
	mutate(a)
	fmt.Println(a) // [0, 2, 3, 4]

}
