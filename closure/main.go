package main

import "fmt"

func Outer(f func(s string)) func() {
	s := "Hello World !!!"
	return func() { f(s) }
}

func Inner(s string) {
	fmt.Println(s)
}

func main() {
	x := Outer(Inner)
	x()
	x()
	x()

	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		v := v
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	// wait for all goroutines to complete before exiting
	for range values {
		<-done
	}
}
