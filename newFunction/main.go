
package main

import "fmt"

func main() {
	p := new(int) // new allocates memory for an int and returns a pointer to it
	fmt.Println(*p) // prints 0, the zero value of int
	*p = 2
	fmt.Println(*p) // prints 2, the value we set
}
