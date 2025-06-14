package main

import (
	"fmt"
	"os"
	"strconv"

	"mylearning/tempconv" // package importing from tempconv
)

func main() {

	// fmt.Println("Hello, world!")
	// tempconv.PrintValue()
	// fmt.Println(tempconv.BoilingC)
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
		   fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Farhenheit(t)
		c := tempconv.Celsius(f)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
