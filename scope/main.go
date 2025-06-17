package main

import (
	"fmt"
	"os"
	"log"
)

var cwd string

func init() {
	// cwd, err := os.Getwd() // := is declaring both cwd and err, but we want to assign cwd to the package-level variable cwd
	var err error
    cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	log.Printf("Working directory: %s", cwd)
}

func f() {}

var g = "g"

func main(){
	f := "f"
	fmt.Println(f)  // "f"; local variable f shadows the package-level func f
	fmt.Println(g) // "g"; package-level variable
	// fmt.Println(h) // compile error: undefined: h
	x := "hello!"
	// range return index and value, but we only need index thats why we use _ to ignore value
	// and we can use only i like for i := range x to get value
	for i := range x {
	    x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x)
		}
	}

	// for _, x := range x {
	// 	x := x + 'A' - 'a'
	// 	fmt.Printf("%c", x)
	// }
}
