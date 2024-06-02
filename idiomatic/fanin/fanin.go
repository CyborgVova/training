package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
)

func FanIn(ctx context.Context, in ...<-chan interface{}) chan interface{} {
	out := make(chan interface{})
	var wg sync.WaitGroup
	for _, ch := range in {
		wg.Add(1)
		go func(ch <-chan interface{}) {
			defer wg.Done()
			for {
				select {
				case c, ok := <-ch:
					if !ok {
						return
					}
					out <- c
				case <-ctx.Done():
					return
				}
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	ch3 := make(chan interface{})

	go func() {
		defer close(ch1)
		for i := 0; i < 400; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := 4; i < 1000; i++ {
			ch2 <- i
		}
	}()

	go func() {
		defer close(ch3)
		for i := 10; i < 12; i++ {
			ch3 <- i
		}
	}()

	for ch := range FanIn(ctx, ch1, ch2, ch3) {
		fmt.Println(ch)
	}
}
