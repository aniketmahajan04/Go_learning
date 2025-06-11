package main

import "fmt"
// Pointers
func main() {
	x := 1
	p := &x // P is a pointer to X holds the address of x
	fmt.Println("Value of x: ", *p) // *p dereferences the pointer, giving the value stored at that address
	*p = 2  // equivalent to x = 2, it will change the value of x
	fmt.Println("Value of x after change: ", x)
}
