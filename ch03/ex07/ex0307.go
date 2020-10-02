package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
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

type Func func(complex128) complex128

var colorPalette = []color.RGBA{
	{0x48, 0x3d, 0x8b, 0xff},
	{0xf0, 0xf8, 0xff, 0xff},
	{0x00, 0x00, 0xff, 0xff},
	{0x8a, 0x2b, 0xe2, 0xff},
}

var chosenColors = map[complex128]color.RGBA{}

func main() {

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(maxY-minY) + minY
		for px := 0; px < width; px++ {
			x := float64(px)/width*(maxX-minX) + minX
			z := complex(x, y)
			img.Set(px, py, Z4(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func Z4(z complex128) color.Color {
	f := func(z complex128) complex128 {
		return z*z*z*z - 1
	}
	fPrime := func(z complex128) complex128 {
		return (z - 1/(z*z*z)) / 4
	}
	return Newton(z, f, fPrime)
}

func Newton(z complex128, f Func, fPrime Func) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= fPrime(z)
		if cmplx.Abs(f(z)) < 1e-6 {
			root := complex(round(real(z), 4), round(imag(z), 4))
			c, ok := chosenColors[root]
			if !ok {
				if len(colorPalette) == 0 {
					log.Fatal("no colors left")
				}
				c = colorPalette[0]
				colorPalette = colorPalette[1:]
				chosenColors[root] = c
			}
			// Convert to YCbCr to make producing different shades easier.
			y, cb, cr := color.RGBToYCbCr(c.R, c.G, c.B)
			scale := math.Log(float64(i)) / math.Log(iterations)
			y -= uint8(float64(y) * scale)
			return color.YCbCr{y, cb, cr}
		}
	}
	return color.Black
}

func round(f float64, digits int) float64 {
	if math.Abs(f) < 0.5 {
		return 0
	}
	pow := math.Pow10(digits)
	return math.Trunc(f*pow+math.Copysign(0.5, f)) / pow
}
