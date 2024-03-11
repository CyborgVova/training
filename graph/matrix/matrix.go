package matrix

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	delimiters = ";,/"
)

type Matrix struct {
	Rows   int
	Cols   int
	Values [][]float64
}

func Parser(patch string, haveSize bool) Matrix {
	matrix := Matrix{}
	fd, err := os.Open(patch)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	r := bufio.NewScanner(fd)
	if haveSize {
		r.Scan()
		val, err := strconv.ParseInt(r.Text(), 10, 32)
		if err != nil {
			panic(err)
		}
		matrix.Rows, matrix.Cols = int(val), int(val)
	}
	for r.Scan() {
		records := strings.Fields(r.Text())
		line := make([]float64, len(records))
		for i, v := range records {
			if strings.ContainsAny(v, delimiters) {
				v = v[:len(v)-1]
			}
			line[i], err = strconv.ParseFloat(v, 64)
			if err != nil {
				panic(err)
			}
		}
		matrix.Values = append(matrix.Values, line)
	}
	if err = r.Err(); err != nil {
		panic(err)
	}
	if !haveSize {
		matrix.Rows = len(matrix.Values)
		matrix.Cols = len(matrix.Values[0])
	}
	return matrix
}

func Printer(m Matrix) {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(m.Values[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}
