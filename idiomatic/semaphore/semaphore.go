package main

import (
	"fmt"
)

type Semaphore struct {
	C chan struct{}
}

func (s *Semaphore) Acquire() {
	s.C <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.C
}

func main() {
	lim := 5

	sem := &Semaphore{C: make(chan struct{}, lim)}

	out := make([]int, 20)

	for i := 0; i < 20; i++ {
		sem.Acquire()
		go func(x int) {
			defer sem.Release()
			out[x] = x * x
		}(i)
	}

	for _, c := range out {
		fmt.Print(c, " ")
	}
	fmt.Println()
}
