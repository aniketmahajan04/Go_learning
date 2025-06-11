package main

import "fmt"

func main() {
	var s string
	var n int
	var b bool
	// var i, j, k int // multiple variable declaration of same type
	// var bo, flo, st = true, 3.14, "hello" // multiple variable declaration of different types
	fmt.Println(s) // " " it will print empty string
	fmt.Println(n) // 0
	fmt.Println(b) // false

	// aniket := "Aniket" // short variable declaration
	// fmt.Println(aniket)

	// freq := rand.float64() * 3.0
	// if i write i, f : = 10 and forgot to give value to f, it will throw error assignment mismatch: 2 variables but 1 value
	i, f := 10, 3.14 // if we forgot to initialize f, it will be throw error
	bul, f := true, 3.13 // this f will override the previous f
	fmt.Println(i)
	fmt.Println(f)
	fmt.Println(bul)
	/*
	note: If you're using := (short variable declaration),
	you must provide a value for every variable on the left side. Go doesn’t auto-fill with default values in :=.
		output:
		10
		3.13
		true
	*/
}
