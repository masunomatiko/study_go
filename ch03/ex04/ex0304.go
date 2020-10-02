package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	// width, height = 600, 320
	cells   = 100
	xyrange = 30.0

	angle = math.Pi / 6
)

type Params struct {
	height float64
	width  float64
	color  string
}

func (p *Params) set(values url.Values) {
	if values.Get("width") == "" {
		p.width = 600.0

	} else {
		w, _ := strconv.Atoi(values.Get("width"))
		p.width = float64(w)
	}
	if values.Get("height") == "" {
		p.height = 320.0

	} else {
		h, _ := strconv.Atoi(values.Get("height"))
		p.height = float64(h)
	}
	if values.Get("color") == "" {
		p.color = "00ff00"
	} else {
		c := values.Get("color")
		p.color = c
	}
}

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func svg(out io.Writer, p *Params) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", p.width, p.height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _ := corner(i+1, j, p)
			bx, by, _ := corner(i, j, p)
			cx, cy, _ := corner(i, j+1, p)
			dx, dy, _ := corner(i+1, j+1, p)
			if isValid([]float64{ax, ay, bx, by, cx, cy, dx, dy}) {
				fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke: #%s; fill: #000000'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, p.color)
			}
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int, p *Params) (float64, float64, float64) {
	xyscale := p.width / 2 / xyrange
	zscale := p.height * 0.4

	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := p.width/2 + (x-y)*cos30*xyscale
	sy := p.height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func isValid(vals []float64) bool {
	for _, v := range vals {
		if math.IsInf(v, 0) {
			return false
		}
	}
	return true
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	handler := func(w http.ResponseWriter, r *http.Request) {
		p := new(Params)
		p.set(r.URL.Query())
		w.Header().Set("Content-Type", "image/svg+xml")

		svg(w, p)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
