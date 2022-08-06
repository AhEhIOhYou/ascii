package main

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
)

// - 10
var ASCII_ITEMS = [...]string{"-", "+", "*", "!", "?", "#", "&", "%", "$", "#"}
var ITEMS_COUNT = 10

// Pixel struct example
type Pixel struct {
	R int
	G int
	B int
	A int
}

// img.At(x, y).RGBA() returns four uint32 values; we want a Pixel
func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}

// Get the bi-dimensional pixel array
func getPixels(file io.Reader) ([][]Pixel, error) {
	img, _, err := image.Decode(file)

	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]Pixel
	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	return pixels, nil
}

func RGBToASCII(r, g, b int) string {
	var result string
	grayScale := 0.21*float64(r) + 0.72*float64(g) + 0.07*float64(b)
	number := int(float64(ITEMS_COUNT-1) * grayScale / 255)
	result = ASCII_ITEMS[number]
	return result
}

func main() {
	// You can register another format here
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	file, err := os.Open("ху.png")

	if err != nil {
		fmt.Println("Error: File could not be opened")
		os.Exit(1)
	}

	defer file.Close()

	pixels, err := getPixels(file)

	if err != nil {
		fmt.Println("Error: Image could not be decoded")
		os.Exit(1)
	}

	var bitmap [320][320]string

	k := 0
	l := 0

	for i := 0; i < len(pixels); i += 3 {
		for j := 0; j < len(pixels[i]); j += 3 {
			bitmap[k][l] = RGBToASCII(pixels[i][j].R, pixels[i][j].G, pixels[i][j].B)
			l++
		}
		k++
		l = 0
	}

	for i := 0; i < 320; i++ {
		for j := 0; j < 320; j++ {
			print(bitmap[i][j])
		}
		println("")
	}
}
