package main

import "fmt"

func defer1() (i int) {
	i = 1
	defer func() { i++ }() // defer может изменятьи устанавливать
	// именованные возвращаемые значения
	return i
}

func defer2() {
	i := 1
	defer fmt.Println(i) // defer захватывает переменную в момент планирования
	i++
}

func defer3() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i) // порядок вызова нескольких defer LIFO (stack)
	}
}

func main() {
	fmt.Println(defer1()) // i = 2
	defer2()              // i = 1
	defer3()              // [3 2 1 0]
}
