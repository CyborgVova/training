package main

import (
	"fmt"
	"net/http"
	"sync"
)

func CheckUrls(wg *sync.WaitGroup, url string, ch chan string) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("%s - not ok", url)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		ch <- fmt.Sprintf("%s - ok", url)
		wg.Done()
	} else {
		ch <- fmt.Sprintf("%s - not ok", url)
	}
}

func main() {
	urls := []string{
		"http://ozon.ru",
		"https://ozon.ru",
		"http://google.com",
		"http://somesite.com",
		"http://non-existent.domain.tld",
		"https://ya.ru",
		"http://ya.ru",
		"http://ёёёё",
	}

	ch := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)

	for _, url := range urls {
		go CheckUrls(&wg, url, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for c := range ch {
		fmt.Println(c)
	}
	fmt.Println("End of programm")
}
