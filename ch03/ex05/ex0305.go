package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

const (
	minX, minY, maxX, maxY = -2, -2, +2, +2
	width, height          = 1024, 1024
	iterations             = 200
	contrast               = 15
)

func main() {

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(maxY-minY) + minY
		for px := 0; px < width; px++ {
			x := float64(px)/width*(maxX-minX) + minX
			z := complex(x, y)
			img.Set(px, py, Mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func Mandelbrot(z complex128) color.Color {
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			switch {
			case n > 50:
				return color.RGBA{0x00, 0x80, 0x00, 0xff} // color.Green
			default:
				logScale := math.Log(float64(n)) / math.Log(float64(iterations))
				return color.RGBA{0, 0, 0xff - uint8(logScale*0xff), 0xff}
			}
		}
	}
	return color.Black
}
