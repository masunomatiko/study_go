package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

// https://godoc.org/golang.org/x/image/colornames　を参照
// color.Black, color.Green
var palette = []color.Color{
	color.Black,
	color.RGBA{0x48, 0x3d, 0x8b, 0xff},
	color.RGBA{0xf0, 0xf8, 0xff, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff},
	color.RGBA{0x8a, 0x2b, 0xe2, 0xff},
	color.RGBA{0x5f, 0x9e, 0xa0, 0xff},
	color.RGBA{0x64, 0x95, 0xed, 0xff},
	color.RGBA{0x00, 0x00, 0x8b, 0xff},
	color.RGBA{0x00, 0xbf, 0xff, 0xff},
	color.RGBA{0x1e, 0x90, 0xff, 0xff},
	color.RGBA{0xad, 0xd8, 0xe6, 0xff},
	color.RGBA{0x87, 0xce, 0xfa, 0xff},
	color.RGBA{0xb0, 0xc4, 0xde, 0xff},
	color.RGBA{0x00, 0x00, 0xcd, 0xff},
	color.RGBA{0x7b, 0x68, 0xee, 0xff},
	color.RGBA{0x19, 0x19, 0x70, 0xff},
	color.RGBA{0xb0, 0xe0, 0xe6, 0xff},
	color.RGBA{0x41, 0x69, 0xe1, 0xff},
	color.RGBA{0x87, 0xce, 0xeb, 0xff},
	color.RGBA{0x6a, 0x5a, 0xcd, 0xff},
	color.RGBA{0x46, 0x82, 0xb4, 0xff},
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajeous(os.Stdout)
}

func lissajeous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64()
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(i))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
