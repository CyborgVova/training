package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func TikTak() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	for {
		fmt.Print("Tik")
		time.Sleep(time.Second)
		fmt.Println("-Tak, ")
		select {
		case <-ch:
			time.Sleep(time.Second)
			fmt.Println("Stop")
			return
		case <-time.After(time.Second):
		}
	}
}

func main() {
	TikTak()
}
