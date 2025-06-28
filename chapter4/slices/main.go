package main

import "fmt"

func main() {
	months := [...]string{
		1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June",
		7: "Jully", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December",
	}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	fmt.Println(summer)
	// endlessSummer := summer[:5]
	// fmt.Println(endlessSummer)

	// fmt.Println(summer[:20]) // out of range
	//
	// for _, s := range summer {
	// 	for _, q := range Q2 {
	// 		if s == q {
	// 			fmt.Printf("%s appearce in both\n", s)
	// 		}
	// 	}
	// }
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5}
	// reverse(a[:]) // a[:] converting array a into slice of same data
	fmt.Println(equal(a, b))
	// nil slices
	// var w []int     // len(s) == 0, s == nil
	// w = nil 		// len(s) == 0, s == nil
	// w = []int(nil)  // len(s) == 0, s == nil
	// w = []int{}     // len(s) == 0, s != nil
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
