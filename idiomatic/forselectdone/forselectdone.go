package main

import (
	"context"
	"fmt"
	"time"
)

// Паттерн позволяет исключить утечку горутин

// ForSelectDone with control channel
func ForSelectDone(ctrl chan struct{}, data chan int) {
	for {
		select {
		case d := <-data:
			fmt.Println(d * d)
		case <-ctrl:
			return
		}
	}
}

// ForSelectDoneCtx with context
func ForSelectDoneCtx(ctx context.Context, data chan int) {
	for {
		select {
		case d := <-data:
			fmt.Println(d * d)
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	data := make(chan int)
	ctrl := make(chan struct{})

	go ForSelectDone(ctrl, data)

	for i := 1; i <= 5; i++ {
		data <- i
	}
	close(ctrl)
	time.Sleep(time.Second)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go ForSelectDoneCtx(ctx, data)

	for i := 1; i <= 5; i++ {
		data <- i
	}
	time.Sleep(time.Second)
}
