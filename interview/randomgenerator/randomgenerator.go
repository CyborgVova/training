package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

func RandomGenerator(num int) <-chan int {
	out := make(chan int)
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	go func() {
		for i := 0; i < num; i++ {
			out <- rnd.Int()
		}
	}()
	return out
}

func RandomGenerator2() (<-chan int, func()) {
	out, exit := make(chan int), make(chan struct{})
	num := uint32(0)

	rnd := rand.New(rand.NewSource(time.Now().Unix()))

	go func() {
		defer close(out)
		for {
			select {
			case <-exit:
				return
			default:
				out <- rnd.Int()
			}
		}
	}()

	return out, func() {
		if atomic.CompareAndSwapUint32(&num, 0, 1) {
			close(exit)
		}
	}
}

func main() {
	num := 10
	ch := RandomGenerator(num)
	for i := 0; i < num; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println()

	ch, stop := RandomGenerator2()
	defer stop()

	for i := 0; i < 7; i++ {
		fmt.Println(<-ch)
	}

	stop()
}
