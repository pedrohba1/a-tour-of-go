package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image is our own image type for this exercise.
type Image struct {
	Width, Height int
}

// ColorModel returns the Image's color model.
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds returns the dimensions of the image.
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Width, img.Height)
}

// At returns the color of the image at a given point.
func (img Image) At(x, y int) color.Color {
	v := uint8((x + y) / 2) // This is a sample function, you can use other formulae to generate interesting patterns
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{256, 256} // Setting the image dimensions here
	pic.ShowImage(m)
}
