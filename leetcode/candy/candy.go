package main

import (
	"fmt"
)

func Candy(rating []int) int {
	len := len(rating)
	result := len

	for i := 1; i < len; i++ {
		if rating[i-1] != rating[i] {
			result++
		}
	}
	return result
}

func main() {
	rating1 := []int{1, 0, 2}       // 5  *** 2-1-2
	rating2 := []int{1, 2, 2}       // 4  *** 1-2-1
	rating3 := []int{1, 3, 2, 2, 1} // 7  *** 1-2-1-2-1

	fmt.Println(Candy(rating1))
	fmt.Println(Candy(rating2))
	fmt.Println(Candy(rating3))
}
