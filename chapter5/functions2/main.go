package main

import (
	"fmt"
	"log"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating system"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

// func square(n int) int     { return n * n }
// func negative(n int) int   { return -n }
// func product(m, n int) int { return m * n }

// anonymous functions
func squares() func() int {
	var x int
	return func() int {
		x++ // x is not valid to anonymous func but it still
		// updating the value of it
		return x * x
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	// f := square
	// fmt.Println(f(3))

	// f = negative
	// fmt.Println(f(3))
	// fmt.Printf("%T\n", f)

	/*
		 	its like overloadingt
			f = product // compile error: can't assign f(int, int) int to f(int) int
	*/
	// l := squares()
	// fmt.Println(l())
	// fmt.Println(l())
	// fmt.Println(l())

	for i, course := range toposort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func toposort(m map[string][]string) []string {
	var order []string            // final sorted list
	seen := make(map[string]bool) // remembers what nodes you've already visited (to avoid repeats)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
