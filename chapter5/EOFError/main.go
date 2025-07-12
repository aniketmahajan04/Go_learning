package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

// WOF is the error returned by Read when no more input is available.
var EOF = errors.New("EOF")

func main() {
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break // finished reading
		}
		if err != nil {
			return fmt.Errorf("read failed: %v", err)
		}
		// use r
	}
}
