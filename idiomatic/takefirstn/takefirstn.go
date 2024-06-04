package main

import (
	"context"
	"fmt"
)

func TakeFirstN(ctx context.Context, ch chan int, n int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			out <- <-ch
		}
	}()
	return out
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			ch <- i
		}
	}()

	for c := range TakeFirstN(ctx, ch, 7) {
		fmt.Print(c, " ")
	}
	fmt.Println()
}
