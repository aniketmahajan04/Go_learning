package tempconv

import "fmt"

type Celsius float64
type Farhenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC   Celsius = 0
	BoilingC    Celsius = 100
	AbsoluteZeroK Kelvin = 0
)

func (c Celsius) String() string {
    return fmt.Sprintf("%g°C", c)
}
func (f Farhenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
 func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}

func PrintValue() {
	converted := CToF(100)
	fmt.Println(converted)
}

