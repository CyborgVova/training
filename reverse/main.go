package main

import (
	"container/list"
	"fmt"
	"unicode/utf8"
)

func ReversQuick(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < len(s); {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func main() {
	l := list.New()
	l.PushBack(111)
	l.PushBackList(list.New())
	fmt.Println(l)
	fmt.Printf("%p\n", l)
	a := []int{1, 2, 3, 4, 5}
	m := map[int]string{2: "two", 3: "three"}
	fmt.Println(min(2, 3, 6))
	fmt.Println(max(2, 3, 6))
	clear(a)
	fmt.Println(len(a), cap(a))
	clear(m)
	fmt.Println(m)
	fmt.Println(len(m))
	// fmt.Println(ReversQuick(ReversQuick("Hello, 世界")))
	// fmt.Println(ReversQuick("Hello, 世界"))
	// fmt.Println(ReversQuick(ReversQuick("The quick brown 狐 jumped over the lazy 犬")))
}
