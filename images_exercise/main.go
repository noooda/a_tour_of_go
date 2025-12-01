package main

import (
	"image"
	"image/color"
	"golang.org/x/tour/pic"
)

type MyImage struct {
	Width, Height int
}

func (m MyImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (m MyImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.Width, m.Height)
}

func (m MyImage) At(x, y int) color.Color {
	r := uint8((x * 255) / m.Width)
	g := uint8((y * 255) / m.Height)
	b := uint8(128)
	return color.RGBA{r, g, b, 255}
}

func main() {
	m := MyImage{Width: 800, Height: 600}	
	pic.ShowImage(m)
}

