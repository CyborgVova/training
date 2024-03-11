package main

import (
	"fmt"
	"log"

	"github.com/cyborgvova/training/graph/bfs"
	"github.com/cyborgvova/training/graph/dfs"
	"github.com/cyborgvova/training/graph/matrix"
)

func main() {
	fmt.Println("Start Programm")
	m := matrix.Parser("files/matrix.txt", false)
	matrix.Printer(m)
	for i := 1; i <= m.Cols; i++ {
		res, err := bfs.BreadthFirstSearch(m, i)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res)
	}
	m = matrix.Parser("files/test.txt", true)
	matrix.Printer(m)
	for i := 1; i <= m.Cols; i++ {
		res, err := bfs.BreadthFirstSearch(m, i)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res)
	}

	m = matrix.Parser("files/test.txt", true)
	matrix.Printer(m)
	res, err := bfs.BreadthFirstSearch(matrix.Matrix{}, 2)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(res)

	m = matrix.Parser("files/noquad.txt", false)
	matrix.Printer(m)
	res, err = bfs.BreadthFirstSearch(m, 2)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(res)

	m = matrix.Parser("files/new.txt", false)
	matrix.Printer(m)
	for i := 1; i <= m.Cols; i++ {
		res, err := bfs.BreadthFirstSearch(m, i)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res)
	}

	m = matrix.Parser("files/new.txt", false)
	matrix.Printer(m)
	for i := 1; i <= m.Cols; i++ {
		res, err := dfs.DepthFirstSearch(m, i)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res)
	}
	fmt.Println("End Programm")
}
