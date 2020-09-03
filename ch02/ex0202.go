package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Celsius float64
type Fahrenheit float64
type Feet float64
type Metre float64
type Pound float64
type Kilogram float64

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (f Feet) String() string       { return fmt.Sprintf("%gft", f) }
func (m Metre) String() string      { return fmt.Sprintf("%gm", m) }
func (p Pound) String() string      { return fmt.Sprintf("%glb", p) }
func (k Kilogram) String() string   { return fmt.Sprintf("%gkg", k) }
func CtoF(c Celsius) Fahrenheit     { return Fahrenheit(c*9/5 + 32) }
func FtoC(f Fahrenheit) Celsius     { return Celsius((f - 32) * 5 / 9) }
func FtoM(f Feet) Metre             { return Metre(f * 0.3048) }
func MtoF(m Metre) Feet             { return Feet(m / 0.3048) }
func PtoK(p Pound) Kilogram         { return Kilogram(p / 2.20462) }
func KtoP(k Kilogram) Pound         { return Pound(k * 2.20462) }

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	if len(os.Args) > 1 {
		arg := os.Args[1]
		val, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fail to ParseFloat: %s", arg)
			os.Exit(1)
		}
		printConv(val)

	} else {
		var val float64

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Printf("input value: ")
		var tmp string
		if scanner.Scan() {
			tmp = scanner.Text()
		}
		val, err := strconv.ParseFloat(tmp, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fail to ParseFloat: %s", tmp)
			os.Exit(1)
		}

		printConv(val)
	}
}

func printConv(val float64) {
	fa := Fahrenheit(val)
	ce := Celsius(val)
	fmt.Printf("%s = %s, %s = %s\n", fa, FtoC(fa), ce, CtoF(ce))
	fe := Feet(val)
	me := Metre(val)
	fmt.Printf("%s = %s, %s = %s\n", fe, FtoM(fe), me, MtoF(me))
	po := Pound(val)
	ki := Kilogram(val)
	fmt.Printf("%s = %s, %s = %s\n", po, PtoK(po), ki, KtoP(ki))
}
