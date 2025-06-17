package main

import (
	"fmt"
	"strconv"
	"os"
	"flag"
)


func main() {

    conveType := flag.String("type", "C", "Type of conversion: C for Celsius to Fahrenheit, F for Fahrenheit to Celsius, M for Meters to Feet, K for Kilograms to Pounds")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
	    fmt.Println("please provide a number to convert")
	}

	// loop on args and convert each one based on the type of conversion
	for _, arg := range args {
		unit, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		// This will check the type of conversion and call the appropriate function
		switch *conveType {
           case "C":
			f := CToF(unit)
		    fmt.Printf("Celsius to Fahrenheit: %f\n", f)
		   case "F":
		    c := FToC(unit)
		   fmt.Printf("Fahernheit to Celsius: %f\n", c)
		   case "M":
		    ft := MToFt(unit)
		    fmt.Printf("Meters to Feet: %f\n", ft)
		   case "K":
		    k := KgToPound(unit)
		   fmt.Printf("Kilograns to Pounds: %f\n", k)
		}
	}

	// Following code will take arg and convert it to all the uits
	// for _, arg := range os.Args[1:] {
	// 	unit, err := strconv.ParseFloat(arg, 64)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
	// 		os.Exit(1)
	// 	}
	//
	// 	f := CToF(unit)
	// 	c := FToC(unit)
	// 	ft := MToFt(unit)
	// 	p := KgToPound(unit)
	// 	fmt.Println("Celsius to Fahrenheit: ", f)
	// 	fmt.Println("Fahrenheit to Celsius: ", c)
	// 	fmt.Println("Meters to Feet: ", ft)
	// 	fmt.Println("Kilograms to Pounds: ", p)
	// }
}

func CToF(c float64) float64 {
   return (c * 9 / 5 + 32)
}
func FToC(f float64) float64 {
   return ((f - 32) * 5 / 9)
}

func MToFt(m float64) float64 {
   return (m * 3.28084)
}
func KgToPound(kg float64) float64 {
   return (kg * 2.20462)
}
