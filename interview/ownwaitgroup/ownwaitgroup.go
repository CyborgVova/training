package main

import (
	"fmt"
)

type Semaphore chan struct{}

func (s Semaphore) Add(i int) {
	for i > 0 {
		s <- struct{}{}
		i--
	}
}

func (s Semaphore) Done() {
	<-s
}

func (s Semaphore) Wait() {
	for {
		if len(s) == 0 {
			return
		}
	}
}

func main() {
	size := 5
	ch := make(chan int)
	sem := make(Semaphore, size)

	for i := 0; i < size; i++ {
		sem.Add(1)
		i := i
		go func() {
			ch <- i * i
			sem.Done()
		}()

	}

	go func() {
		sem.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
