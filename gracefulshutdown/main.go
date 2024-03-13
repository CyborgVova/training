package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
)

func head(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s !!!", r.URL.Query().Get("name"))
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	http.HandleFunc("/hello/", head)
	srv1 := &http.Server{
		Addr: ":8080",
	}

	srv2 := &http.Server{
		Addr: ":8181",
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Starting server 1 on port [::]:8080")
		if err := srv1.ListenAndServe(); err != nil {
			log.Printf("Server 1 exit reason: %v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Starting server 2 on port [::]:8181")
		if err := srv2.ListenAndServe(); err != nil {
			log.Printf("Server 2 exit reason: %v", err)
		}
	}()

	go func() {
		<-ctx.Done()
		srv1.Shutdown(context.Background())
		fmt.Println("Server 1 is shutdown")
	}()

	go func() {
		<-ctx.Done()
		srv2.Shutdown(context.Background())
		fmt.Println("Server 2 is shutdown")
	}()

	wg.Wait()
	fmt.Println("End of programm")
}
