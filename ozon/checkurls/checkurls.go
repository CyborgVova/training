package main

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
)

func CheckUrls(ctx context.Context, counter *atomic.Int32, url string, out chan<- string) {
	select {
	case <-ctx.Done():
		return
	default:
		resp, err := http.Get(url)
		if err != nil {
			out <- fmt.Sprintf("адрес %s - not ok", url)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			out <- fmt.Sprintf("адрес %s - not ok", url)
		} else {
			out <- fmt.Sprintf("адрес %s - ok", url)
			counter.Add(1)
		}
	}
}

func main() {
	var in = make(chan string, 1)
	var out = make(chan string, 1)
	var counter atomic.Int32
	var wg sync.WaitGroup
	var urls = []string{
		"http://ozon.ru",
		"https://ozon.ru",
		"http://google.com",
		"http://somesite.com",
		"http://non-existent.domain.tld",
		"https://ya.ru",
		"http://ya.ru",
		"http://ёёёё",
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	go func() {
		for _, url := range urls {
			in <- url
		}
		close(in)
	}()

	workers := 0
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				if counter.Load() == 2 {
					out <- "End of programm"
					break
				}
				url, ok := <-in
				if !ok {
					break
				}
				CheckUrls(ctx, &counter, url, out)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for c := range out {
		fmt.Println(c)
	}
}
