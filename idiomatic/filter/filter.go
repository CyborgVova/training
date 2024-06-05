package main

import (
	"context"
	"fmt"
)

func Filter(ctx context.Context, ch chan int, f func(int) bool) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case c, ok := <-ch:
				if !ok {
					return
				}

				if f(c) {
					out <- c
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return out
}

func main() {
	ch := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	f := func(i int) bool {
		return i%2 == 1
	}

	go func() {
		for i := 1; i < 25; i++ {
			ch <- i
		}
		close(ch)
	}()

	for c := range Filter(ctx, ch, f) {
		fmt.Print(c, " ")
	}
	fmt.Println()
}
