package main

import (
	"fmt"
)

func div1(x, y int) (out int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	out = x / y
	return out
}

func div2(x, y int) (out int) {
	out = x / y
	return out
}

func main() {
	fmt.Println("Start programm")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	fmt.Println(div1(10, 2))
	fmt.Println(div1(10, 0))
	fmt.Println(div2(10, 5))
	fmt.Println(div2(10, 0))
	fmt.Println("End of programm")
}
