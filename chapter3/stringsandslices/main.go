package main

import (
	"bytes"
	"fmt"
	"strings"
)

func intToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(basename("a/b/c.go")) // c
	fmt.Println(basename("c.d.go"))   // c.d
	fmt.Println(basename("abc"))      // abc

	fmt.Println(basename2("a/b/c.go")) // c
	fmt.Println(basename2("c.d.go"))   // c.d
	fmt.Println(basename2("abc"))      // abc

	fmt.Println(comma("12345"))

	fmt.Println(intToString([]int{1, 2, 3}))
}

func basename(s string) string {
	// basename removes directory components and a .suffix

	// Discard last '/' and everything before
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// Preserve everything before last '.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

// simpler version with strings.LastIndex
func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
