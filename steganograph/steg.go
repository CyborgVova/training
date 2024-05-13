package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strings"
)

func Decode(file string) string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	bound := img.Bounds()
	out := []byte{}
	bt := byte(0)
	for y := 0; y < bound.Dy(); y++ {
		for x := 0; x < bound.Dx(); x++ {
			r, g, b, a := img.At(x, y).RGBA()
			if x%2 == 0 {
				bt |= byte(r & 1)
				bt = bt << 1
				bt |= byte(g & 1)
				bt = bt << 1
				bt |= byte(b & 1)
				bt = bt << 1
				bt |= byte(a & 1)
				bt = bt << 1
			} else {
				bt |= byte(r & 1)
				bt = bt << 1
				bt |= byte(g & 1)
				bt = bt << 1
				bt |= byte(b & 1)
				bt = bt << 1
				bt |= byte(a & 1)
				out = append(out, bt)
				bt = 0
			}
		}
	}

	return string(out)
}

func Run(file string, rgba *image.RGBA, payload string) {
	str := strings.Builder{}
	str.Grow(len(payload) + 2)
	str.WriteString("<" + payload + ">")
	bytes := []byte(str.String())
	mask := uint32(254)
	var left, right uint32
	j := 0
	for y := 0; y < rgba.Rect.Dy(); y++ {
		for x := 0; x < rgba.Rect.Dx(); x++ {
			r, g, b, a := rgba.At(x, y).RGBA()
			if j < len(bytes) {
				r &= mask
				g &= mask
				b &= mask
				a &= mask
				left = uint32(bytes[j]) >> 4
				tmp := bytes[j] << 4
				right = uint32(tmp) >> 4
				if x%2 == 0 {
					r |= left >> 3 & 1
					g |= left >> 2 & 1
					b |= left >> 1 & 1
					a |= left & 1
				} else {
					r |= right >> 3 & 1
					g |= right >> 2 & 1
					b |= right >> 1 & 1
					a |= right & 1
					j++
				}
			}
			rgba.Set(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
		}
	}
	ff, _ := os.Create("BlackQuad_st.png")
	png.Encode(ff, rgba)
}

func Encode(file string, payload string) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	rgba := image.NewRGBA(img.Bounds())

	for y := 0; y < rgba.Rect.Dy(); y++ {
		for x := 0; x < rgba.Rect.Dx(); x++ {
			r, g, b, a := img.At(x, y).RGBA()
			rgba.Set(x, y, color.RGBA{byte(r), byte(g), byte(b), byte(a)})
		}
	}

	Run(file, rgba, payload)
}

func main() {
	Encode("BlackQuad.png",
		"Этот текст находится внутри картинки\n"+
			"Этот текст находится внутри картинки\n"+
			"Этот текст находится внутри картинки\n"+
			"Этот текст находится внутри картинки\n"+
			"Этот текст находится внутри картинки\n"+
			"Этот текст находится внутри картинки")
	fmt.Println(Decode("BlackQuad_st.png"))
}
