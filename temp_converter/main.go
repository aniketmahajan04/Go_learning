package main

import "fmt"

const boilingF = 212.0 // boilingF is package level declaration (as is main)

func main() {
	// variables f and c are local to the function main
	var f = boilingF
	var c = (f - 32) * 5 / 9
	// In Go (Golang), the %g format verb is used to format floating-point numbers.
	// It automatically chooses between the %e (scientific notation) and %f
	// (decimal notation) formats,
	fmt.Printf("Boiling point = %gF or %gC", f, c)
	// output
	// Boiling point = 212°F or 100°C
}
