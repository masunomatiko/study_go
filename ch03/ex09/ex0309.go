package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"net/url"
	"strconv"
)

const (
	// minX, minY, maxX, maxY = -2, -2, +2, +2
	width, height = 1024, 1024
	iterations    = 200
	contrast      = 15
)

type Params struct {
	minX, minY, maxX, maxY int
	zoom                   int
}

func (p *Params) set(values url.Values) {
	if values.Get("maxX") == "" {
		p.maxX = 2

	} else {
		v, _ := strconv.Atoi(values.Get("maxX"))
		p.maxX = v
	}
	if values.Get("maxY") == "" {
		p.maxY = 2

	} else {
		v, _ := strconv.Atoi(values.Get("maxY"))
		p.maxY = v
	}
	if values.Get("minX") == "" {
		p.minX = 2

	} else {
		v, _ := strconv.Atoi(values.Get("minX"))
		p.minX = v
	}
	if values.Get("minX") == "" {
		p.minY = 2

	} else {
		v, _ := strconv.Atoi(values.Get("minX"))
		p.minY = v
	}
	if values.Get("zoom") == "" {
		p.zoom = 1

	} else {
		v, _ := strconv.Atoi(values.Get("zoom"))
		p.zoom = v
	}
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := new(Params)
		p.set(r.URL.Query())
		// for name := range Params {
		// 	s := r.FormValue(name)
		// 	if s == "" {
		// 		continue
		// 	}
		// 	f, err := strconv.ParseFloat(s, 64)
		// 	if err != nil {
		// 		http.Error(w, fmt.Sprintf("query param %s: %s", name, err), http.StatusBadRequest)
		// 		return
		// 	}
		// 	params[name] = f
		// }
		// if params["maxX"] <= params["minX"] || params["maxY"] <= params["minY"] {
		// 	http.Error(w, fmt.Sprintf("min coordinate greater than max"), http.StatusBadRequest)
		// 	return
		// }
		minX := p.minX
		maxX := p.maxX
		minY := p.minY
		maxY := p.maxY
		zoom := p.zoom

		lenX := maxX - minX
		midX := minX + lenX/2
		minX = midX - lenX/2/zoom
		maxX = midX + lenX/2/zoom
		lenY := maxY - minY
		midY := minY + lenY/2
		minY = midY - lenY/2/zoom
		maxY = midY + lenY/2/zoom

		img := image.NewRGBA(image.Rect(0, 0, width, height))
		for py := 0; py < height; py++ {
			y := float64(py)/height*float64(maxY-minY) + float64(minY)
			for px := 0; px < width; px++ {
				x := float64(px)/width*float64(maxX-minX) + float64(minX)
				z := complex(x, y)

				img.Set(px, py, Mandelbrot(z))
			}
		}
		err := png.Encode(w, img)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
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

// // Some other interesting functions:

// func acos(z complex128) color.Color {
// 	v := cmplx.Acos(z)
// 	blue := uint8(real(v)*128) + 127
// 	red := uint8(imag(v)*128) + 127
// 	return color.YCbCr{192, blue, red}
// }

// func sqrt(z complex128) color.Color {
// 	v := cmplx.Sqrt(z)
// 	blue := uint8(real(v)*128) + 127
// 	red := uint8(imag(v)*128) + 127
// 	return color.YCbCr{128, blue, red}
// }

// // f(x) = x^4 - 1
// //
// // z' = z - f(z)/f'(z)
// //    = z - (z^4 - 1) / (4 * z^3)
// //    = z - (z - 1/z^3) / 4
// func newton(z complex128) color.Color {
// 	const iterations = 37
// 	const contrast = 7
// 	for i := uint8(0); i < iterations; i++ {
// 		z -= (z - 1/(z*z*z)) / 4
// 		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
// 			return color.Gray{255 - contrast*i}
// 		}
// 	}
// 	return color.Black
// }
