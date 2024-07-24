package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Server struct {
	ID             int
	NumConnections int
	mu             sync.RWMutex
}

type Balancer struct {
	servers []*Server
}

func createServers(num int) []*Server {
	servers := make([]*Server, num)

	for i := 0; i < num; i++ {
		servers[i] = &Server{ID: i + 1}
	}

	return servers
}

func NewBalancer(servers []*Server) *Balancer {
	return &Balancer{servers: servers}
}

func (s *Server) handlerServer(w http.ResponseWriter, _ *http.Request) {
	s.mu.Lock()
	s.NumConnections++
	s.mu.Unlock()

	time.Sleep(time.Millisecond * 100)
	fmt.Fprintf(w, "Server #%d response\n", s.ID)

	s.mu.Lock()
	s.NumConnections--
	s.mu.Unlock()
}

func (b *Balancer) GetServerWithLeastConnections() *Server {
	minConn := int(^uint(0) >> 1)
	var server *Server
	for _, s := range b.servers {
		s.mu.RLock()
		if s.NumConnections < minConn {
			minConn = s.NumConnections
			server = s
		}
		s.mu.RUnlock()
	}

	return server
}

func (b *Balancer) handlerBalancer(w http.ResponseWriter, r *http.Request) {
	server := b.GetServerWithLeastConnections()
	server.handlerServer(w, r)
}

func main() {
	numServers := 5
	servers := createServers(numServers)
	balancer := NewBalancer(servers)

	http.HandleFunc("/ping", balancer.handlerBalancer)

	fmt.Println("Server start on :8080 ...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("starting server:", err)
	}

}
