package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func ExecuteTask() {
	time.Sleep(2 * time.Second)
	fmt.Println("ExecuteTask success")
}

func ExecuteTaskWrapper(t time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()

	ctrl := make(chan struct{})
	go func() {
		ExecuteTask()
		close(ctrl)
	}()

	select {
	case <-ctx.Done():
		return errors.New("stoping by timeout")
	case <-ctrl:
		return nil
	}
}

func main() {
	err := ExecuteTaskWrapper(3 * time.Second)
	if err != nil {
		fmt.Println(err)
	}
}
