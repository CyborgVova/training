package main

import (
	"fmt"
)

func Filter(ch chan int, f func(int) bool) chan int {
	out := make(chan int)
	go func() {
		for c := range ch {
			if f(c) {
				out <- c
			}
		}
		close(out)
	}()

	return out
}

func main() {
	ch := make(chan int)
	f := func(i int) bool {
		return i%2 == 0
	}

	go func() {
		for i := 1; i < 25; i++ {
			ch <- i
		}
		close(ch)
	}()

	for c := range Filter(ch, f) {
		fmt.Print(c, " ")
	}
	fmt.Println()
}
