package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	s string
}

var (
	singleton *Singleton
	once      sync.Once
)

func GetInstance() {
	singleton = &Singleton{s: "singleton"}
}

func main() {
	once.Do(GetInstance)
	fmt.Printf("%p ", singleton)
	fmt.Println(singleton.s)

	once.Do(GetInstance)
	singleton.s = "Only Once"
	fmt.Printf("%p ", singleton)
	fmt.Println(singleton.s)

	once.Do(GetInstance)
	fmt.Printf("%p ", singleton)
	fmt.Println(singleton.s)
}
