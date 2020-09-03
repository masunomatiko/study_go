package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezeingC    Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g℃", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g℉", f) }
func CtoF(c Celsius) Fahrenheit     { return Fahrenheit(c*9/5 + 32) }
func FtoC(f Fahrenheit) Celsius     { return Celsius((f - 32) * 5 / 9) }
func CtoK(c Celsius) Kelvin         { return Kelvin(c - AbsoluteZeroC) }
func FtoK(f Fahrenheit) Kelvin      { return Kelvin(FtoC(f) - AbsoluteZeroC) }
func KtoC(k Kelvin) Celsius         { return Celsius(k) + AbsoluteZeroC }
func KtoF(k Kelvin) Fahrenheit      { return Fahrenheit((k+CtoK(AbsoluteZeroC))*9/5 + 32) }
