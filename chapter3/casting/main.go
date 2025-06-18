package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		% known as verb
		[1] known as adverb
	*/
	f := 3.141 // a float64
	i := int(f)
	fmt.Println(f, i) // "3.141 3"
	f = 1.99
	fmt.Println(int(i)) // "1"
	f = 1e100           // a float64
	i = int(f)          // result is impelementation-dependent
	fmt.Println(f, i)   // 1e+100 -9223372036854775808
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // 438 666 0666
	x := int64(0xdeadbeef)
	fmt.Printf("%d %#[1]x %#[1]X\n", x) // 3735928559 0xdeadbeef 0XDEADBEEF
	ascii := 'a'
	unicode := '+'
	newLine := '\n'

	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // 3735928559 0xdeadbeef 0XDEADBEEF
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // 3735928559 0xdeadbeef 0XDEADBEEF
	fmt.Printf("%d %[1]c %[1]q\n", newLine) // 3735928559 0xdeadbeef 0XDEADBEEF
	var fl float32 = 16777216
	fmt.Println(fl == fl+1) // true
	for x := range 8 {      // similar as for x := 0; x < 8; x++

		fmt.Printf("x = %d  eA = %8.3f\n", x, math.Exp(float64(x)))
		/*
			result of for loop
						x = 0  eA =    1.000
						x = 1  eA =    2.718
						x = 2  eA =    7.389
						x = 3  eA =   20.086
						x = 4  eA =   54.598
						x = 5  eA =  148.413
						x = 6  eA =  403.429
						x = 7  eA = 1096.633
		*/
	}
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) //  "0 -0 +Inf -Inf NaN"
	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan) // false, false, false
}
