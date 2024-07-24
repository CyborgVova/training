package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Server struct {
	ID int
}

type RoundRobin struct {
	servers []*Server
	current int
	mu      sync.Mutex
}

func createServers(num int) []*Server {
	servers := make([]*Server, num)
	for i := 0; i < num; i++ {
		servers[i] = &Server{ID: i + 1}
	}

	return servers
}

func NewBalancer(servers []*Server) *RoundRobin {
	return &RoundRobin{servers: servers}
}

func (s *Server) handlerServer(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Handle server #%d\n", s.ID)
	time.Sleep(time.Millisecond * 100)
}

func (b *RoundRobin) GetNextServer() *Server {
	b.mu.Lock()
	defer b.mu.Unlock()
	server := b.servers[b.current]
	b.current++
	if b.current == len(b.servers) {
		b.current = 0
	}
	return server
}

func (b *RoundRobin) handlerBalancer(w http.ResponseWriter, r *http.Request) {
	server := b.GetNextServer()
	server.handlerServer(w, r)
}

func main() {
	numServers := 5
	servers := createServers(numServers)
	balancer := NewBalancer(servers)

	http.HandleFunc("/ping", balancer.handlerBalancer)

	fmt.Println("Starting server on :8080 ...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("starting server:", err)
	}

}
