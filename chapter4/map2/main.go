package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

var graph = make(map[string]map[string]bool)

func main() {
	// seen := make(map[string]bool) // a set of strings
	// input := bufio.NewScanner(os.Stdin)
	// for input.Scan() {
	// 	line := input.Text()
	// 	if !seen[line] {
	// 		seen[line] = true
	// 		fmt.Println(line)
	// 	}
	// }
	// if err := input.Err(); err != nil {
	// 	fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
	// 	os.Exit(1)
	// }

	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of length of UTF-8 encoding
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // return rune, nbytes and error
		if err == io.EOF {         // EOF means End of file
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcound: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
