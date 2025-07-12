package main

import (
	"fmt"
	"math"
)

/*
   Syntax of function
   func name(parameter-list) (result-list) {
      body
   }
*/

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func f(i, j, k int, s, t string) {
	// instead of writing i int, j int, k int we can write i, j, k int fuction will
	// understand that all three are of type int
	fmt.Println(i, j, k, s, t)
}

// type of function declaration
func add(x int, y int) int { return x + y }

func sub(x, y int) (z int) { z = x - y; return }

func first(x int, _ int) int { return x }

func zero(int, int) int { return 0 }

func main() {
	fmt.Println(hypot(3, 4)) // output: 5
	f(0, 1, 2, "world", "hello")
}
