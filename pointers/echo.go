
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() { // main is redeclare in main package that why throwing error
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Print()
	}
}
