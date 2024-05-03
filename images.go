package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (im Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 255, 255)
}

func (im Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (im Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x * y), uint8(x * y), 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
