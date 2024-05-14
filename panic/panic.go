package main

import "fmt"

func f() {
	fmt.Println("Start f function")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recover with i = %v\n", r)
		}
	}()

	g(0)
	fmt.Println("End of f function")
}

func g(i int) {
	fmt.Printf("Inside g function i = %v\n", i)
	if i > 3 {
		fmt.Println("Panicing ...")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Printf("defer in g. i = %v\n", i)
	g(i + 1)
}

func main() {
	fmt.Println("Start programm")
	f()
	fmt.Println("End of programm")
}
