package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	ch <- 1 // залочится главная горутина так как никто не вычитывает канал
	go func() {
		fmt.Println(<-ch)
	}()
}
