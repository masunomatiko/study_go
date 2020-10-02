package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/big"
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

			img.Set(px, py, MandelbrotBigFloat(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
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

func Mandelbrot64(z complex128) color.Color {

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + complex64(z)
		if cmplx.Abs(complex128(v)) > 2 {
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

func MandelbrotBigFloat(z complex128) color.Color {

	zR := (&big.Float{}).SetFloat64(real(z))
	zI := (&big.Float{}).SetFloat64(imag(z))
	var vR, vI = &big.Float{}, &big.Float{}
	for i := uint8(0); i < iterations; i++ {
		// (r+i)^2 = r^2 + 2ri + i^2
		vR2, vI2 := &big.Float{}, &big.Float{}
		vR2.Mul(vR, vR).Sub(vR2, (&big.Float{}).Mul(vI, vI)).Add(vR2, zR)
		vI2.Mul(vR, vI).Mul(vI2, big.NewFloat(2)).Add(vI2, zI)
		vR, vI = vR2, vI2
		squareSum := &big.Float{}
		squareSum.Mul(vR, vR).Add(squareSum, (&big.Float{}).Mul(vI, vI))
		if squareSum.Cmp(big.NewFloat(4)) == 1 {
			switch {
			case i > 50:
				return color.RGBA{0x00, 0x80, 0x00, 0xff} // color.Green
			default:
				logScale := math.Log(float64(i)) / math.Log(float64(iterations))
				return color.RGBA{0, 0, 0xff - uint8(logScale*0xff), 0xff}
			}
		}
	}
	return color.Black
}

// 終わらなかった
// func MandelbrotRat(z complex128) color.Color {
// 	zR := (&big.Rat{}).SetFloat64(real(z))
// 	zI := (&big.Rat{}).SetFloat64(imag(z))
// 	var vR, vI = &big.Rat{}, &big.Rat{}
// 	for i := uint8(0); i < iterations; i++ {
// 		// (r+i)^2 = r^2 + 2ri + i^2
// 		vR2, vI2 := &big.Rat{}, &big.Rat{}
// 		vR2.Mul(vR, vR).Sub(vR2, (&big.Rat{}).Mul(vI, vI)).Add(vR2, zR)
// 		vI2.Mul(vR, vI).Mul(vI2, big.NewRat(2, 1)).Add(vI2, zI)
// 		vR, vI = vR2, vI2
// 		squareSum := &big.Rat{}
// 		squareSum.Mul(vR, vR).Add(squareSum, (&big.Rat{}).Mul(vI, vI))
// 		if squareSum.Cmp(big.NewRat(4, 1)) == 1 {
// 			switch {
// 			case i > 50:
// 				return color.RGBA{0x00, 0x80, 0x00, 0xff} // color.Green
// 			default:
// 				logScale := math.Log(float64(i)) / math.Log(float64(iterations))
// 				return color.RGBA{0, 0, 0xff - uint8(logScale*0xff), 0xff}
// 			}
// 		}
// 	}
// 	return color.Black
// }
