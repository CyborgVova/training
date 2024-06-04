package main

import (
	"fmt"
)

// 1. Расскажи подробно что происходит

// Вариант 1

// func main() {
// 	a := []int{1, 2}  // [1 2]
// 	a = append(a, 3)  // [1 2 3] len = 3 cap = 4
// 	b := append(a, 4) // [1 2 3 4] та же память len = 4 cap = 4
// 	c := append(a, 5) // [1 2 3 5] та же память len = 4 cap = 4

// 	fmt.Println(b) // [1 2 3 5]
// 	fmt.Println(c) // [1 2 3 5]
// }

// Вариант 2

func main() {
	a := []int{1, 2}  // [1 2]
	a = append(a, 3)  // [1 2 3]
	a = append(a, 7)  // [1 2 3 7] len = 4 cap = 4
	b := append(a, 4) // [1 2 3 7 4] новая область памяти
	c := append(a, 5) // [1 2 3 7 5] новая область памяти

	fmt.Println(b) // [1 2 3 7 4]
	fmt.Println(c) // [1 2 3 7 5]
}
