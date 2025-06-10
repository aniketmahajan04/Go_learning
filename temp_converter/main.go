package main

import "fmt"

const boilingF = 212.0 // boilingF is package level declaration (as is main)

func main() {
	// variables f and c are local to the function main
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("Boiling point = %gF or %gC", f, c)
}
