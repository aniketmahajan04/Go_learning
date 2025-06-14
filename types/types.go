
package main

import "fmt"

type Celsius float64 // this will used as Celsius(value like int or float) it convert it into float64
// but is we give Celsius("100") it will give error
type Fahrenheit float64

const (
	AbsoluteZerC Celsius = -273.15
	FreezingC    Celsius = 0
	BoilingC     Celsius = 100
)

func main() {
	c := Celsius(100)
	f := CToF(c)
	fmt.Printf("Celsius: %.2f, Fahrenheit: %.2f\n", c, f)

	c2 := FToC(f)
	fmt.Printf("Fahrenheit: %.2f, Celsius: %.2f\n", f, c2)

	// we can perform arithmatic operation on Underlying types
	fmt.Printf("%g\n", BoilingC-FreezingC)
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC))
	// fmt.Printf("%g\n", boilingF-FreezingC) // Fehrenheit - Celsius type mimatch compile time error
	var a Celsius
	var b Fahrenheit
	fmt.Println( a == 0) // true
	fmt.Println( b >= 0) // true
	// fmt.Println(a == b) // compile error: type mismatch
	fmt.Println(a == Celsius(b)) // true
	fmt.Println(c2.String())
	fmt.Printf("%v\n", c2) // %v will call String() method
	fmt.Printf("%s\n", c2)
	fmt.Println(c2)
	fmt.Printf("%g\n", c2) // 100; does not call String
	fmt.Println(float64(c2)) // 100; does not call String
}

func CToF(c Celsius) Fahrenheit {
    return Fahrenheit( c * 9 / 5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}
