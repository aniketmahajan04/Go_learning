package main

import (
	"fmt"
)

func main() {
	/*
			 Go provides two sizes ofcomplex numbers, complex64 and complex128,whose components
		 are float32 and float64 resp ectively.
	*/
	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i
	fmt.Println(x * y)               // (-5+10i)
	fmt.Println(real(x * y))         // -5
	fmt.Println(imag(x * y))         // 10

	/*
		Complex constant support arithmetic with integers and floats, so we can write value
		like 1 + 2i or 2i + 1 naturally.
	*/
	a := 1 + 2i
	fmt.Println(a) // (1+2i)
	b := 3.2 + 4i
	fmt.Println(b) // (3.2+4i)
}
