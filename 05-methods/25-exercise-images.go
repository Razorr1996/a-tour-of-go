package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

func f(x, y int) int {
	//return (x + y) / 2
	return x * y
	//return x ^ y
}

type Image struct {
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	const size = 256
	return image.Rect(0, 0, size, size)
}

func (i Image) At(x, y int) color.Color {
	v := uint8(f(x, y))
	return color.RGBA{R: v, G: v, B: 255, A: 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
