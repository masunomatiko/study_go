package main

import (
	"image/color"
	"testing"
)

func benchmarkMandelbrot(b *testing.B, f func(complex128) color.Color) {
	for i := 0; i < b.N; i++ {
		f(complex(float64(i), float64(i)))
	}
}

func BenchmarkMandelbrotComplex128(b *testing.B) {
	benchmarkMandelbrot(b, Mandelbrot)
}

func BenchmarkMandelbrotComplex64(b *testing.B) {
	benchmarkMandelbrot(b, Mandelbrot64)
}

func BenchmarkMandelbrotBigFloat(b *testing.B) {
	benchmarkMandelbrot(b, MandelbrotBigFloat)
}

func BenchmarkMandelbrotRat(b *testing.B) {
	benchmarkMandelbrot(b, MandelbrotRat)
}
