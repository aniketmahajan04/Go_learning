package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	var a [3]int         // this is how we define arrays in golang
	b := [4]int{1, 2, 3} // dont know it is slice or and arrays
	// slice := []int{1, 2, 3} // now this is slice coz it has no fixed size

	fmt.Println(a[0])        // 0 coz its empty array
	fmt.Println(b[0])        // 1
	fmt.Println(a[len(a)-1]) // 0

	// Print the indices and elements.
	// for i, v := range a {
	// 	fmt.Printf("%d = %d\n", i, v)
	// 	/*
	// 		Answer of for loop
	// 		0 = 0
	// 		1 = 0
	// 		2 = 0
	// 	*/
	// }

	/*
	   This loop will print only value and that underscore is used for index since we are not
	   doing anything with index we are skipping it (don't try for v := range a) this will only print
	   the elements
	*/
	// for _, v := range a {
	// 	fmt.Printf("%d\n", v)
	// }

	var q [3]int = [3]int{1, 2, 3} // its how golang let us define arrays
	/*
				var q [3]int, q is type of [3]int means it has explicitly given but when you will pass
		it in function as args you have to give same type like below "count function"
		or that will throw error
	*/
	// m := [3]int{1, 2, 3}           // this one is another way
	fmt.Println(count(q)) // 3
	increement(&q)
	for i, v := range q {
		/*
			0 = 1
			1 = 2
			2 = 3
		*/
		fmt.Printf("%d = %d\n", i, v)
	}

	n := [...]int{
		1,
		2,
		3,
	} // [...] is known as "ellipsis" and the length will be considerd from initialization
	fmt.Printf("%T\n", n)

	symbols := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	fmt.Println(RMB, symbols[RMB])         // 3 "¥"
	r := [...]int{99: -1}                  // array of 99 zeros and one -1
	fmt.Printf("%d and %d\n", r[0], r[99]) // 0 and -1
	h := [2]int{1, 2}
	j := [...]int{1, 2}
	k := [2]int{1, 3}
	fmt.Println(h == j, j == k, k == h) // true false false
	// l := [3]int{1, 2}
	// fmt.Println(h == l)
}

func count(q [3]int) int {
	count := 0
	for i := range q {
		count += i
	}
	return count
}

func increement(q *[3]int) {
	for _, v := range q {
		v++
	}
}
