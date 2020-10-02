// package main

// import (
// 	"image"
// 	"image/color"
// 	"image/png"
// 	"math"
// 	"math/cmplx"
// 	"os"
// )

// func main() {
// 	const (
// 		minX, minY, maxX, maxY = -2, -2, +2, +2
// 		width, height          = 1024, 1024
// 		epsX                   = (maxX - minX) / width
// 		epsY                   = (maxY - minY) / height
// 	)

// 	img := image.NewRGBA(image.Rect(0, 0, width, height))
// 	for py := 0; py < height; py++ {
// 		y := float64(py)/height*(maxY-minY) + minY
// 		for px := 0; px < width; px++ {
// 			x := float64(px)/width*(maxX-minX) + minX
// 			// Supersampling:
// 			subPixels := make([]color.Color, 0)
// 			for i := 0; i < 2; i++ {
// 				for j := 0; j < 2; j++ {
// 					z := complex(x+offX[i], y+offY[j])
// 					subPixels = append(subPixels, mandelbrot(z))
// 				}
// 			}
// 			img.Set(px, py, avg(subPixels))
// 		}
// 	}
// 	png.Encode(os.Stdout, img) // NOTE: ignoring errors
// }

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
	epsX                   = (maxX - minX) / width
	epsY                   = (maxY - minY) / height
)

// func avg(colors []color.Color) color.Color {
// 	var r, g, b, a uint32
// 	n := len(colors)
// 	for _, c := range colors {
// 		r, g, b, a := c.RGBA()
// 		r += r / uint32(n)
// 		g += g / uint32(n)
// 		b += b / uint32(n)
// 		a += a / uint32(n)
// 	}
// 	return color.RGBA64{uint16(r), uint16(g), uint16(b), uint16(a)}
// }
func avg(colors []color.Color) color.Color {
	var r, g, b, a uint16
	n := len(colors)
	for _, c := range colors {
		rr, gg, bb, aa := c.RGBA()
		r += uint16(rr / uint32(n))
		g += uint16(gg / uint32(n))
		b += uint16(bb / uint32(n))
		a += uint16(aa / uint32(n))
	}
	return color.RGBA64{r, g, b, a}
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

func main() {
	diffX := []float64{-epsX, epsX}
	diffY := []float64{-epsY, epsY}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(maxY-minY) + minY
		for px := 0; px < width; px++ {
			x := float64(px)/width*(maxX-minX) + minX
			// スーパーサンプリング
			subPixels := make([]color.Color, 0)
			for i := 0; i < 2; i++ {
				for j := 0; j < 2; j++ {
					z := complex(x+diffX[i], y+diffY[j])
					subPixels = append(subPixels, Mandelbrot(z))
				}
			}
			img.Set(px, py, avg(subPixels))
		}
	}
	png.Encode(os.Stdout, img)
}
