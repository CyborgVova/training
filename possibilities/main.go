package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {
	s := "Hello World !!!"
	s += " Hello Friends !!!"
	if strings.Contains(s, "Friends") {
		s = strings.ReplaceAll(s, "Friends", "Volodja")
	}
	fmt.Println(s)
	// fmt.Println(mult(2))

	engToRus := make(map[string]string, 0)
	engToRus["Hi"] = "Привет"
	engToRus["Bye"] = "Пока"
	engToRus["Boo"] = "Бу-у"
	for key, value := range engToRus {
		fmt.Println("Key - [" + key + "]  Value - [" + value + "]")
	}
	fmt.Println(engToRus["Boo"])
	delete(engToRus, "Boo")
	fmt.Println(engToRus["Boo"])
	fmt.Println(engToRus["Bye"])

	const nihongo = "日本語"
	fmt.Println(nihongo[0:6])
	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}

	x, y, z := 1, 2, 3
	fmt.Println(x, y, z)

	func() {
		fmt.Println(777)
	}()

	fmt.Println(mult(2))

	xd := 12
	_ = xd

	str2 := "¿Cómoestás?"
	fmt.Printf("%s\n", str2)

	for _, v := range str2 {
		fmt.Printf("%c", v)
	}
	fmt.Println()

	countdown, err := strconv.Atoi("10")
	if err == nil {
		fmt.Println(countdown)
	}

	diff(10, "string", 129.129, 'Я', "ЯЗЬ")

	xType := reflect.TypeOf('Ы')
	xValue := reflect.ValueOf('Ы')
	fmt.Println(xType, xValue)

	mt := MyType('Ы')
	mt.Printing()

	ms := []MyInt{50, 60, 70, 80, 90}
	for _, m := range ms {
		go m.Show()
	}
	time.Sleep(1000 * time.Millisecond)

	go func() {
		for {
			fmt.Println("I'm running in my own go routine")
		}
	}()
	<-time.After(1 * time.Second)

	ch := make(chan string)
	go func() {
		for {
			fmt.Println("TEST")
			ch <- "Start"
		}
	}()
	<-ch
}

type (
	MyInt  int
	MyType rune
)

func (mi MyInt) Show() {
	fmt.Println(mi)
}

func (mt MyType) Printing() {
	fmt.Printf("'%c'", mt)
}

func diff(x ...interface{}) {
	for _, v := range x {
		switch t := v.(type) {
		case rune:
			fmt.Printf("%c\n", t)
		default:
			fmt.Println(v)
		}
	}
}

func mult(x int) []int {
	table := make([]int, 0)
	for i := 0; i < 11; i++ {
		table = append(table, x*i)
	}
	return table
}
