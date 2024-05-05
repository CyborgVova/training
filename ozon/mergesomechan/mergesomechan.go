package main

import (
	"fmt"
	"sync"
)

func Merge(chs ...chan int) chan int {
	var wg sync.WaitGroup
	out := make(chan int, 1)
	for _, ch := range chs {
		wg.Add(1)
		go func(ch chan int) {
			defer wg.Done()
			for c := range ch {
				out <- c
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
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	go func() {
		ch1 <- 10
		ch1 <- 20
		close(ch1)
	}()

	go func() {
		ch2 <- 30
		close(ch2)
	}()

	go func() {
		ch3 <- 40
		ch3 <- 50
		ch3 <- 60
		ch3 <- 70
		close(ch3)
	}()

	for ch := range Merge(ch1, ch2, ch3) {
		fmt.Println(ch)
	}

}
