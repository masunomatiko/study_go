package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)
			if isValid([]float64{ax, ay, bx, by, cx, cy, dx, dy}) {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='#%06x'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, getColor(az, bz, cz, dz))
			}
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
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

func getColor(az, bz, cz, dz float64) uint64 {
	z := (az + bz + cz + dz) / 4
	// b := uint32((1.0 - z) / (1.0 - -0.245) * 0xff)
	// c := fmt.Sprintf("%X", 0xff0000-(b<<16)+b)
	// for i := len(c); i < 6; i++ {
	// 	c = "0" + c
	// }
	// return c

	// redFactor -> 0..1, blueFactor 1..0
	redFactor := (z + 1) * 0.5
	blueFactor := (-z + 1) * 0.5

	// 00..ff, ff..00
	redByte := int(255.0 * redFactor)
	blueByte := int(255.0 * blueFactor)

	// left shift of red byte to its position
	var redWord uint64 = uint64(redByte) << 16
	var blueWord uint64 = uint64(blueByte)

	//ORing the words to get final color like 0x7f007f
	return blueWord | redWord
}
