package main

import (
	"fmt"
)

// 1. Что выведется?
// 2. Зная обо всех таких нюансах, которые могут возникнуть, какие есть рекомендации?

//  Вариант 1

// func mod(a []int) {
// 	for i := range a {
// 		a[i] = 5
// 	}
// 	fmt.Println(a) // [5, 5, 5, 5] не было изменения капасити
// }

// func main() {
// 	sl := []int{1, 2, 3, 5}
// 	mod(sl)
// 	fmt.Println(sl) // [5, 5, 5, 5] изменения применились
// }

//  Вариант 2

// func mod(a []int) {
// 	for i := range a {
// 		a[i] = 5
// 	}
// 	fmt.Println(a) //[5, 5, 5, 5]
// }

// func main() {
// 	sl := make([]int, 4, 8) // len = 4, cap = 8
// 	sl[0] = 1
// 	sl[1] = 2
// 	sl[2] = 3
// 	sl[3] = 5 // [1, 2, 3, 5]
// 	mod(sl)
// 	fmt.Println(sl) // [5, 5, 5, 5]
// }

// Вариант 3

// func mod(a []int) {
// 	a = append(a, 125) // [5, 5, 5, 5, 125] // капасити не поменялся
// 	for i := range a {
// 		a[i] = 5
// 	}
// 	fmt.Println(a) // [5, 5, 5, 5, 5] // та же область памяти
// }

// func main() {
// 	sl := make([]int, 4, 8)
// 	sl[0] = 1
// 	sl[1] = 2
// 	sl[2] = 3
// 	sl[3] = 5 // [1, 2, 3, 5]
// 	mod(sl)
// 	fmt.Println(sl) // [5, 5, 5, 5] тут размер не изменился
// }

// Вариант 4

func mod(a []int) {
	a = append(a, 125) // изменился капасити новая область памяти
	for i := range a {
		a[i] = 5
	}
	fmt.Println(a) // [5, 5, 5, 5, 5, 5]
}

func main() {
	sl := []int{1, 2, 3, 4, 5}
	mod(sl)
	fmt.Println(sl) // [1, 2, 3, 4, 5]
}
