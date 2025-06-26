package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"io"
	"os"
)

func main() {
	useSHA384 := flag.Bool("sha384", false, "Use sha384 hashing")
	useSHA512 := flag.Bool("sha512", false, "Use sha512 hashing")
	flag.Parse()

	// if *useSHA384 {
	// 	fmt.Println(sha512.New384())
	// 	fmt.Println("sha384")
	// } else if *useSHA512 {
	// 	fmt.Println(sha512.New())
	// 	fmt.Println("sha512")
	// } else {
	// 	fmt.Println(sha256.New())
	// 	fmt.Println("sha256")
	// }
	var h hash.Hash
	switch {
	case *useSHA384:
		h = sha512.New384()
	case *useSHA512:
		h = sha512.New()
	default:
		h = sha256.New()
	}

	_, err := io.Copy(h, os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)

		os.Exit(1)
	}
	fmt.Printf("%x\n", h.Sum(nil))
	// echo "hello" | go run cryptoSha2/main.go -sha512
}
