package main

import (
	"context"
	"fmt"
)

func Map(ctx context.Context, ch chan int, f func(int) int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for {
			select {
			case c, ok := <-ch:
				if !ok {
					return
				}

				out <- f(c)
			case <-ctx.Done():
				return
			}
		}
	}()

	return out
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := make(chan int)
	f := func(i int) int {
		return i * i
	}

	go func() {
		defer close(ch)
		for i := 1; i < 15; i++ {
			ch <- i
		}
	}()

	for c := range Map(ctx, ch, f) {
		fmt.Print(c, " ")
	}
	fmt.Println()
}
