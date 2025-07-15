package main

import (
	"fmt"
	"os"
	"strings"
)

func sum(vals ...int) int {
	/*
	 it can take multiple arguments like we are storing it
	 int slice of int
	*/
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
func f(...int) {}
func g([]int)  {}
func main() {
	fmt.Println(sum())           // 0
	fmt.Println(sum(3))          // 3
	fmt.Println(sum(1, 2, 3, 4)) // 10

	value := []int{1, 2, 3, 4}
	fmt.Println(sum(value[2], value[0]))

	fmt.Printf("%T\n", f) // func(...int)
	fmt.Printf("%T\n", g) // func([]int)

	linenum, name := 12, "count"
	errorf(linenum, "undefined: %s", name) // Line 12: undefined: count

	fmt.Println(stringJoinf("Aniket", "Mahajan"))
}

func minf(vals ...int) int {
	if len(vals) == 0 {
		panic("minf requires at least one argument")
	}
	min := vals[0]
	for _, v := range vals[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

func maxf(vals ...int) int {
	if len(vals) == 0 {
		panic("maxf requare at least one argument")
	}
	max := vals[0]

	for _, v := range vals[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func stringJoinf(strs ...string) string {
	// joinedStr := strings[0]
	// for _, str := range strings[1:] {
	// 	joinedStr += str
	// }

	// return joinedStr

	return strings.Join(strs, " ")
}

func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}
