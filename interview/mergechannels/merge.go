package main

import (
	"fmt"
	"sync"
)

func Merge(in ...<-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var wg sync.WaitGroup
		for i := range in {
			i := i
			wg.Add(1)
			go func() {
				for v := range in[i] {
					out <- v
				}
				wg.Done()
			}()
		}
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ch1, ch2, ch3 := make(chan int), make(chan int), make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i + 3
		}
		close(ch1)
	}()

	go func() {
		for i := 0; i < 7; i++ {
			ch2 <- i + 2
		}
		close(ch2)
	}()

	go func() {
		for i := 0; i < 5; i++ {
			ch3 <- i * 2
		}
		close(ch3)
	}()

	for ch := range Merge(ch1, ch2, ch3) {
		fmt.Print(ch, " ")
	}

	fmt.Println()
}
