package main

import (
	"fmt"
)

const boilingF = 212.0 // boilingF is package level declaration (as is main)

func main() {
	// variables f and c are local to the function main
	var f = boilingF
	var c = (f - 32) * 5 / 9
	// In Go (Golang), the %g format verb is used to format floating-point numbers.
	// It automatically chooses between the %e (scientific notation) and %f
	// (decimal notation) formats,
	fmt.Printf("Boiling point = %g°F or %g°C", f, c)
	// output
	// Boiling point = 212°F or 100°C

	const freezingF, boilingF = 32.0, 212.0                 // it has const decl mean there will be dublicate inside function
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // 32°F = 0°C
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))   // 212°F = 100°C
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
