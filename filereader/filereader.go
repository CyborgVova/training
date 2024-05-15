package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		fmt.Println(string(sc.Bytes()))
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
}
