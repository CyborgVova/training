package main

import (
	"fmt"
)

func main() {
	var foo []int
	var bar []int

	foo = append(foo, 1) // foo len = 1 cap = 1
	foo = append(foo, 2) // foo len = 2 cap = 2
	foo = append(foo, 3) // foo len = 3 cap = 4
	bar = append(foo, 4) // bar len = 4 cap = 4 // область памяти foo
	foo = append(foo, 5) // foo len = 4 cap = 4

	fmt.Println(foo, bar) // [1, 2, 3, 5] [1, 2, 3, 5] адна область памяти

}
