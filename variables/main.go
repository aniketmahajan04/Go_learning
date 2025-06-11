package main

import "fmt"

func main() {
	var s string
	var n int
	var b bool
	var i, j, k int // multiple variable declaration of same type
	var bo, flo, st = true, 3.14, "hello" // multiple variable declaration of different types
	fmt.Println(s) // " " it will print empty string
	fmt.Println(n) // 0
	fmt.Println(b) // false

	aniket := "Aniket" // short variable declaration
	fmt.Println(aniket)

	freq := rand.float64() * 3.0
}
