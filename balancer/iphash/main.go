package main

import (
	"errors"
	"fmt"
	"hash/fnv"
	"log"
	"net"
	"net/http"
	"time"
)

type Server struct {
	ID int
}

type IPHash struct {
	servers []*Server
}

func createServers(num int) []*Server {
	servers := make([]*Server, num)
	for i := 0; i < num; i++ {
		servers[i] = &Server{ID: i + 1}
	}

	return servers
}

func NewBalancer(servers []*Server) *IPHash {
	return &IPHash{servers: servers}
}

func (s *Server) handlerServer(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Handle server #%d\n", s.ID)
	time.Sleep(time.Millisecond * 100)
}

func (b *IPHash) GetServer(ip net.IP) (*Server, error) {
	hasher := fnv.New32a()
	_, err := hasher.Write([]byte(ip))
	if err != nil {
		return nil, errors.Join(errors.New("index not got: "), err)
	}

	index := hasher.Sum32() % uint32(len(b.servers))

	return b.servers[index], nil
}

func (b *IPHash) handlerBalancer(w http.ResponseWriter, r *http.Request) {
	ip := net.ParseIP(r.RemoteAddr)
	server, err := b.GetServer(ip)
	if err != nil {
		fmt.Fprintf(w, "Internal server error %d\n", http.StatusInternalServerError)
		log.Println(err)
		return
	}
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
