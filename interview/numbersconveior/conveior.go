package main

import (
	"fmt"
)

func Conveior(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for ch := range in {
			out <- ch * ch
		}
		close(out)
	}()

	return out
}

func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for c := range Conveior(ch) {
		fmt.Print(c, " ")
	}

	fmt.Println()
}
