package main

import (
	"bytes"
	"container/list"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"

	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"regexp"
	"time"
)

func fatalf(format string, v ...any) {
	fmt.Printf(format, v...)
}

var (
	weight     = 512
	height     = 512
	newImgSize = weight * 5
	resImg     = image.NewRGBA(image.Rect(0, 0, newImgSize, newImgSize))
)

func Reverse(s string) (result string) {
	for _, ch := range s {
		result = string(ch) + result
	}
	return
}

func main() {
	fmt.Println(Reverse("ZombI"))
	fmt.Println(len("ZombI"))
	fmt.Println(len("ЗомбИ"))
	str := "The best from the west>"
	idx := strings.IndexByte(str, '>')
	fmt.Println(str[idx:])
	start := time.Now()
	defer func() { fmt.Println("Time of executing:", time.Since(start)) }()
	file, err := os.Open("packman.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	f, err := os.Create("test.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < newImgSize/weight*newImgSize/weight; i++ {
		for x := 0; x < newImgSize/height*height; x++ {
			for y := i * weight; y < i*weight+weight; y++ {
				r, g, b, a := img.At(x%height, y%weight).RGBA()
				resImg.Set(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
			}

		}
	}

	if err := png.Encode(f, resImg); err != nil {
		fatalf("%s", err)
	}

	var (
		reImg  = regexp.MustCompile(`<img src="(/[^"]+.png)"/>`)
		reLink = regexp.MustCompile(`<a href="([^"]+)">`)
		html   = regexp.MustCompile(`<a href="([^"]+)">([^"]+)</a>`)
	)
	fmt.Println(reImg)
	fmt.Println(reLink)

	l := list.New()
	l.PushBack("one")
	l.PushBack(2)
	for l1 := l.Front(); l1 != nil; {
		fmt.Println(l1.Value)
		l1 = l1.Next()
	}
	flag.Parse()

	if flag.NArg() != 2 {
		fatalf("Position parametrs is not 2\n")
	}

	fmt.Println(flag.Arg(0))
	fmt.Println(flag.Args())
	var out []interface{}
	out = append(out, "Price", 13.08)
	fatalf("%s: %.2f;\n", out...)
	var m = make([]int, 3, 5)
	m[2] = 4
	fmt.Println(m)

	var links, images, htmls []string
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	buf := bytes.Buffer{}
	if _, err := io.Copy(&buf, resp.Body); err != nil {
		fatalf("%v", err)
	}
	// fmt.Println(buf.String())
	for _, link := range reLink.FindAllStringSubmatch(buf.String(), -1) {
		links = append(links, link[1])
	}

	for _, img := range reImg.FindAllStringSubmatch(buf.String(), -1) {
		images = append(images, img[1])
	}

	for _, h := range html.FindAllStringSubmatch(buf.String(), -1) {
		// fmt.Println("h: ", h)
		htmls = append(htmls, h[2])
	}

	fmt.Println(images)
	fmt.Println(links)
	fmt.Println(htmls)
}
