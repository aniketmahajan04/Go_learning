package main

import "fmt"

func main() {
	s := "hello world"
	fmt.Println(len(s))     // 11
	fmt.Println(s[0], s[7]) // 104 111
	// c := s[len(s)]
	// fmt.Println(c) // panic: runtime error: index out of range [11] with length 11
	fmt.Println(s[0:5]) // hello
	fmt.Println(s[:5])  // hello
	fmt.Println(s[6:])  // world
	fmt.Println(s[:])   // hello world
	// + operator concates the strings
	fmt.Println("goodbye" + s[5:]) // goodbye world

	// strings are immutable
	a := "left foot"
	t := a
	a += ", right foot"

	fmt.Println(a) // "left foot, right foot"
	fmt.Println(t) // "left foot"
	// a[0] = 'L' //cannot assign to a[0] (neither addressable nor a map index expression)
}
