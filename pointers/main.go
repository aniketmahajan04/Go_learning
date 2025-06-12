package main

import "fmt"
// Pointers
func main() {
	x := 1
	//p := &x // P is a pointer to X holds the address of x
	// fmt.Println("Value of x: ", *p) // *p dereferences the pointer, giving the value stored at that address
	//*p = 2  // equivalent to x = 2, it will change the value of x
	// fmt.Println("Value of x after change: ", x)

	//Two pointers are equal if both points to same variable
	var a, b int
	fmt.Println(&a == &a, &a == &b, &a == nil) // true false false

	var m = f()
	fmt.Println(m) // it will print an address of v
	fmt.Println(f() == f()) // false
	incr(&x) // 2
	fmt.Println(incr(&x)) // 3 will printed

}

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}
