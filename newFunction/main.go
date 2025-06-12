
package main

import "fmt"

func main() {
	p := new(int) // new allocates memory for an int and returns a pointer to it
	fmt.Println(*p) // prints 0, the zero value of int
	*p = 2
	fmt.Println(*p) // prints 2, the value we set
	a := newInt()
	fmt.Println(a)
	b := newInt2()
	fmt.Println(b)

	c := new(int)
	d := new(int)
	fmt.Println(c == d) // false, c and d are different pointers
}

func newInt() *int {
	return new(int) // returns a pointer to a new int
}

func newInt2() *int {
    var dummy int
	return &dummy
}
