package main

import (
	"fmt"
	"time"
)

func writeToChannel(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	close(ch)
}

func main() {
	var ch = make(chan int)
	go writeToChannel(ch)

	for v := range ch {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
