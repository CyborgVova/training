package main

import (
	"fmt"
)

// 1. Что будет содержать s после инициализации?
// 2. Что произойдет в println для слайса и для мапы?

func a(s []int) {
	s = append(s, 37) // изменение не будет видно снаружы
}

func b(m map[int]int) {
	m[3] = 33
}

func main() {
	s := make([]int, 3, 8)    // len = 3 cap = 8 [0, 0, 0]
	m := make(map[int]int, 8) // map size = 8

	// add to slice
	a(s)
	// fmt.Println(s[3]) // out of range (здесь будет паника)

	// add to map
	b(m)
	fmt.Println(m[3]) // 33 (мап передается по ссылке)
}
