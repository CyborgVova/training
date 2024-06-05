package main

import (
	"fmt"
)

// Что выведет код? Должны выводится все значения

func main() {
	a := 5000
	for i := 0; i < a; i++ {
		go fmt.Println(i) // undefined behavior
	}

	// программа закончится раньше чем отработают все горутины
	// надо использовать sync.WaitGroup{} и обернуть go func()

	// var wg sync.WaitGroup
	// wg.Add(5000)
	// for i := 0; i < a; i++ {
	// 	i := i
	// 	go func() {
	// 		fmt.Println(i) // undefined behavior
	// 		defer wg.Done()
	// 	}()
	// }

	// wg.Wait()
}
