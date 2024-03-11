// docker pull vkcup2022/golang:lates
// docker run -it --rm -d -p 8080:80 --name web vkcup2022/golang:latest
package main

import (
	"bytes"
	"container/list"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"time"
)

const (
	collageWidth = 32
	imageWidth   = 512
	imageHeight  = 512
)

var (
	reImg  = regexp.MustCompile(`<img src="(/[^"]+.png)"/>`)
	reLink = regexp.MustCompile(`<a href="([^"]+)">`)
)

func main() {
	flag.Parse()

	start := time.Now()

	if len(flag.Args()) != 2 {
		fatalf("usage %s <url> <result.png>", os.Args[0])
	}
	addr := flag.Arg(0)
	resFile := flag.Arg(1)

	fmt.Printf("Start crawling from %s and storing into %s\n", addr, resFile)

	u, err := url.Parse(addr)
	if err != nil {
		fatalf("%v", err)
	}

	visited := map[string]struct{}{}
	queue := list.New()
	queue.PushBack(u.Path)

	cnt := 0

	resImg := image.NewRGBA(image.Rect(0, 0, imageWidth*collageWidth, imageHeight*32))

	for queue.Len() > 0 {
		location := queue.Front().Value.(string)
		queue.Remove(queue.Front())

		if _, exists := visited[location]; exists {
			continue
		}
		visited[location] = struct{}{}

		links, images, err := getHtmlData(u.Scheme + "://" + u.Host + location)
		if err != nil {
			fatalf("%v", err)
		}

		for _, link := range links {
			queue.PushBack(link)
		}

		for _, imgSrc := range images {
			img, err := getImage(u.Scheme + "://" + u.Host + imgSrc)
			if err != nil {
				fatalf("%v", err)
			}

			for x := 0; x < imageWidth; x++ {
				for y := 0; y < imageHeight; y++ {
					_, _, b, a := img.At(x, y).RGBA()
					resImg.Set((cnt%collageWidth)*imageWidth+x, (cnt/collageWidth)*imageHeight+y, color.RGBA{0, 0, uint8(b), uint8(a)})
				}
			}

			cnt++
		}
	}

	f, err := os.Create(resFile)
	if err != nil {
		fatalf("%s", err)
	}
	defer f.Close()
	if err := png.Encode(f, resImg); err != nil {
		fatalf("%s", err)
	}

	fmt.Printf("Processed %d file in %s\n", cnt, time.Since(start))
}

func getHtmlData(url string) ([]string, []string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("http error: %d", resp.StatusCode)
	}

	buf := bytes.Buffer{}
	if _, err := io.Copy(&buf, resp.Body); err != nil {
		fatalf("%v", err)
	}

	var links, images []string

	for _, link := range reLink.FindAllStringSubmatch(buf.String(), -1) {
		links = append(links, link[1])
	}

	for _, img := range reImg.FindAllStringSubmatch(buf.String(), -1) {
		images = append(images, img[1])
	}

	return links, images, nil
}

func getImage(url string) (image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http error: %d", resp.StatusCode)
	}

	return png.Decode(resp.Body)
}

func fatalf(format string, v ...any) {
	_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n", fmt.Sprintf(format, v...))
	os.Exit(255)
}
